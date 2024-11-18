package store

import (
	"context"
	"errors"
	"fmt"
)

type HashMap struct {
	keyformat string
}

func (h *HashMap) GetKeyFormat() string {
	return h.keyformat
}

func (h *HashMap) SetKeyFormat(format string) {
	h.keyformat = format
}

func (h *HashMap) Delete(id int64) error {
	if RedisIns() == nil {
		return errors.New("redis is nil")
	}
	return RedisIns().RedisCli.HDel(context.Background(), h.keyformat, fmt.Sprintf("%d", id)).Err()
}

func (h *HashMap) HSet(id int64, data interface{}) error {
	if RedisIns() == nil {
		return errors.New("redis is nil")
	}
	return RedisIns().RedisCli.HSet(context.Background(), h.keyformat, fmt.Sprintf("%d", id), data).Err()
}

func (h *HashMap) HGet(id int64) (string, error) {
	if RedisIns() == nil {
		return "", errors.New("redis is nil")
	}
	return RedisIns().RedisCli.HGet(context.Background(), h.keyformat, fmt.Sprintf("%d", id)).Result()
}
