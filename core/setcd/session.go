// @Title
// @Description
// @Author  Wangwengang  2024/1/8 13:00
// @Update  Wangwengang  2024/1/8 13:00
package setcd

import (
	"context"
	v3 "go.etcd.io/etcd/client/v3"
	"time"
)

const defaultSessionTTL = 10

// Session represents a lease kept alive for the lifetime of a client.
// Fault-tolerant applications may use sessions to reason about liveness.
type Session struct {
	client *v3.Client
	opts   *sessionOptions
	id     v3.LeaseID

	cancel context.CancelFunc
	donec  <-chan struct{}
}

// NewSession gets the leased session for a client.
func NewSession(client *v3.Client, opts ...SessionOption) (*Session, error) {
	ops := &sessionOptions{ttl: defaultSessionTTL, ctx: client.Ctx()}
	for _, opt := range opts {
		opt(ops)
	}

	id := ops.leaseID
	if id == v3.NoLease {
		resp, err := client.Grant(ops.ctx, int64(ops.ttl))
		if err != nil {
			return nil, err
		}
		id = v3.LeaseID(resp.ID)
	}

	ctx, cancel := context.WithCancel(ops.ctx)
	// 保证租约一直被续租, 直到通过ctx关闭此租约
	keepAlive, err := client.KeepAlive(ctx, id)
	if err != nil || keepAlive == nil {
		cancel()
		return nil, err
	}

	donec := make(chan struct{})
	s := &Session{client: client, opts: ops, id: id, cancel: cancel, donec: donec}

	// keep the lease alive until client error or cancelled context
	go func() {
		defer close(donec)
		for range keepAlive {
			// eat messages until keep alive channel closes
		}
	}()

	return s, nil
}

// Client is the etcd client that is attached to the session.
func (s *Session) Client() *v3.Client {
	return s.client
}

// Lease is the lease ID for keys bound to the session.
func (s *Session) Lease() v3.LeaseID { return s.id }

// Done returns a channel that closes when the lease is orphaned, expires, or
// is otherwise no longer being refreshed.
func (s *Session) Done() <-chan struct{} { return s.donec }

// Orphan ends the refresh for the session lease. This is useful
// in case the state of the client connection is indeterminate (revoke
// would fail) or when transferring lease ownership.
// 终止刷新租约
func (s *Session) Orphan() {
	s.cancel()
	<-s.donec
}

// Close orphans the session and revokes the session lease.
func (s *Session) Close() error {
	s.Orphan()
	// if revoke takes longer than the ttl, lease is expired anyway
	ctx, cancel := context.WithTimeout(s.opts.ctx, time.Duration(s.opts.ttl)*time.Second)

	// 释放租约
	_, err := s.client.Revoke(ctx, s.id)
	cancel()
	return err
}

type sessionOptions struct {
	ttl     int
	leaseID v3.LeaseID
	ctx     context.Context
}

// SessionOption configures Session.
type SessionOption func(*sessionOptions)

// WithTTL configures the session's TTL in seconds.
// If TTL is <= 0, the default 60 seconds TTL will be used.
func WithTTL(ttl int) SessionOption {
	return func(so *sessionOptions) {
		if ttl > 0 {
			so.ttl = ttl
		}
	}
}

// WithLease specifies the existing leaseID to be used for the session.
// This is useful in process restart scenario, for example, to reclaim
// leadership from an election prior to restart.
func WithLease(leaseID v3.LeaseID) SessionOption {
	return func(so *sessionOptions) {
		so.leaseID = leaseID
	}
}

// WithContext assigns a context to the session instead of defaulting to
// using the client context. This is useful for canceling NewSession and
// Close operations immediately without having to close the client. If the
// context is canceled before Close() completes, the session's lease will be
// abandoned and left to expire instead of being revoked.
func WithContext(ctx context.Context) SessionOption {
	return func(so *sessionOptions) {
		so.ctx = ctx
	}
}
