// @Title
// @Description
// @Author  Wangwengang  2023/12/12 23:55
// @Update  Wangwengang  2023/12/12 23:55
package store

import (
	"time"

	"gorm.io/gorm"
)

type BASE_MODEL struct {
	ID        int64          `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
