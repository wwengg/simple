// @Title
// @Description
// @Author  Wangwengang  2023/12/25 05:16
// @Update  Wangwengang  2023/12/25 05:16
package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type RedisHash struct {
	*RedisBase
	keyformat string
}

func NewRedisHash(redisBase *RedisBase, keyformat string) *RedisHash {
	return &RedisHash{
		RedisBase: redisBase,
		keyformat: keyformat,
	}
}

func (rh *RedisHash) HSet(id int64, data interface{}) error {
	hashId := rh.GetHash(id)
	if strings.Contains(rh.keyformat, "%d") {
		key := fmt.Sprintf(rh.keyformat, hashId)
		return rh.RedisCli.HSet(context.Background(), key, fmt.Sprintf("%d", id), data).Err()
	} else {
		return errors.New("keyformat err")
	}
}

func (rh *RedisHash) HGet(id int64) (string, error) {
	hashId := rh.GetHash(id)
	if strings.Contains(rh.keyformat, "%d") {
		key := fmt.Sprintf(rh.keyformat, hashId)
		return rh.RedisCli.HGet(context.Background(), key, fmt.Sprintf("%d", id)).Result()
	} else {
		return "", errors.New("keyformat err")
	}
}

func (rh *RedisHash) HDel(id int64) error {
	hashId := rh.GetHash(id)
	if strings.Contains(rh.keyformat, "%d") {
		key := fmt.Sprintf(rh.keyformat, hashId)
		return rh.RedisCli.HDel(context.Background(), key, fmt.Sprintf("%d", id)).Err()
	} else {
		return errors.New("keyformat err")
	}
}
