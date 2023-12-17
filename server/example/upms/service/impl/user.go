// @Title
// @Description
// @Author  Wangwengang  2023/12/17 13:09
// @Update  Wangwengang  2023/12/17 13:09
package impl

import (
	"github.com/wwengg/simple/server/example/upms/global"
	"github.com/wwengg/simple/server/example/upms/model"
)

func CreateUser(user *model.SimpleUser) error {
	return global.DBUpms.Create(&user).Error
}

func UpdateUserNick(id int64, nick string) error {
	var user model.SimpleUser
	if err := global.DBUpms.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	user.Nick = nick
	//global.DBUpms.Save(&user)
	err := global.DBUpms.Save(&user).Error

	return err
}

func DeleteUser(user model.SimpleUser) (err error) {
	err = global.DBUpms.Delete(&user).Error
	return err
}
