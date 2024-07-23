package utils

import "strconv"

type int64String string

// 如果不是int64类型的string会返回0，务必确保是int64
func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
