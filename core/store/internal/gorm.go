// @Title
// @Description
// @Author  Wangwengang  2023/12/13 00:04
// @Update  Wangwengang  2023/12/13 00:04
package internal

import (
	"github.com/wwengg/simple/core/slog"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool, log slog.Slog) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	// _default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	// 	SlowThreshold: 200 * time.Millisecond,
	// 	LogLevel:      logger.Warn,
	// 	Colorful:      true,
	// })
	// var logMode DBBASE
	// logMode = new(sconfig.Mysql)

	// switch logMode.GetLogMode() {
	// case "silent", "Silent":
	// 	config.Logger = _default.LogMode(logger.Silent)
	// case "error", "Error":
	// 	config.Logger = _default.LogMode(logger.Error)
	// case "warn", "Warn":
	// 	config.Logger = _default.LogMode(logger.Warn)
	// case "info", "Info":
	// 	config.Logger = _default.LogMode(logger.Info)
	// default:
	// 	config.Logger = _default.LogMode(logger.Info)
	// }
	config.Logger = slog.NewGormZapLogger(log)
	return config
}
