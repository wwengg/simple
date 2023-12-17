// @Title
// @Description
// @Author  Wangwengang  2023/12/16 15:11
// @Update  Wangwengang  2023/12/16 15:11
package global

import (
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/srpc"
	"gorm.io/gorm"
)

var (
	Log    slog.Slog
	RPC    srpc.SRPC
	DBList map[string]*gorm.DB
	DBUpms *gorm.DB

	CONFIG Config
)
