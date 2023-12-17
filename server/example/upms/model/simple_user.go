// @Title
// @Description
// @Author  Wangwengang  2023/12/16 21:44
// @Update  Wangwengang  2023/12/16 21:44
package model

import "github.com/wwengg/simple/core/store"

type SimpleUser struct {
	store.BASE_MODEL
	Nick  string `json:"nick" form:"nick" gorm:"column:nick;comment:用户昵称;type:varchar(20);size:20;"`
	Phone int64  `json:"phone" form:"phone" gorm:"column:phone;comment:用户昵称;type:bigint(20);size:20;"`
}
