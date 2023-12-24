// @Title
// @Description
// @Author  Wangwengang  2023/12/23 12:16
// @Update  Wangwengang  2023/12/23 12:16
package store

type Cache interface {
	GetHash(k int64) int64
}
