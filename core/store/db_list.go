// @Title
// @Description
// @Author  Wangwengang  2023/12/12 23:57
// @Update  Wangwengang  2023/12/12 23:57
package store

import (
	"github.com/wwengg/simple/core/sconfig"
	"gorm.io/gorm"
)

func DBList() map[string]*gorm.DB {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range sconfig.S_CONF.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(sconfig.Mysql{GeneralDB: info.GeneralDB})
		//case "mssql":
		//	dbMap[info.AliasName] = GormMssqlByConfig(config.Mssql{GeneralDB: info.GeneralDB})
		//case "pgsql":
		//	dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		//case "oracle":
		//	dbMap[info.AliasName] = GormOracleByConfig(config.Oracle{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	return dbMap
}
