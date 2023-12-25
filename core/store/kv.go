// @Title
// @Description
// @Author  Wangwengang  2023/12/25 05:11
// @Update  Wangwengang  2023/12/25 05:11
package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type KV struct {
	*RedisBase
	keyformat string
	expire    time.Duration
	keepTTL   bool
}

func NewKV(redisBase *RedisBase, keyformat string, expire time.Duration, keepTTL bool) *KV {
	if keepTTL {
		expire = redis.KeepTTL
	}
	return &KV{
		RedisBase: redisBase,
		keyformat: keyformat,
		expire:    expire,
		keepTTL:   false,
	}
}

func (kv *KV) Set(id int64, data interface{}) error {
	if strings.Contains(kv.keyformat, "%d") {
		key := fmt.Sprintf(kv.keyformat, id)
		return kv.RedisCli.Set(context.Background(), key, data, kv.expire).Err()
	} else {
		return errors.New("keyformat err")
	}
}

func (kv *KV) Get(id int64) (string, error) {
	if strings.Contains(kv.keyformat, "%d") {
		key := fmt.Sprintf(kv.keyformat, id)
		return kv.RedisCli.Get(context.Background(), key).Result()
	} else {
		return "", errors.New("keyformat err")
	}
}

// 重新计算过期时间，并获取值
func (kv *KV) Expire(id int64) error {
	if !strings.Contains(kv.keyformat, "%d") {
		return errors.New("keyformat err")
	}
	key := fmt.Sprintf(kv.keyformat, id)
	if kv.keepTTL {
		// TODO log
		return nil
	} else {
		return kv.RedisCli.Expire(context.Background(), key, kv.expire).Err()
	}

}
