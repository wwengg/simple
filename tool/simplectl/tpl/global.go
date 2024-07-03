// @Title
// @Description
// @Author  Wangwengang  2023/12/24 23:16
// @Update  Wangwengang  2023/12/24 23:16
package tpl

func GlobalTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
package global

import (
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/srpc"
	"github.com/wwengg/simple/core/store"
	"gorm.io/gorm"
)

var (
	CONFIG *Config
	LOG    slog.Slog
	SRPC   srpc.SRPC
	DBList map[string]*gorm.DB
)


func InitSlog() {
	// 初始化日志
	LOG = slog.NewZapLog(&CONFIG.Slog)
}

func InitSRPC() {
	// 初始化SRPC
	SRPC = srpc.NewSRPCClients(&CONFIG.RPC)
}

func InitDB() {
	// 初始化DBList
	DBList = store.DBList(&CONFIG.DBList)
}`)
}

func GlobalConfigTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
package global

import (
	"fmt"
	"github.com/wwengg/simple/core/sconfig"
	"reflect"
)

type Config struct {
	sconfig.Config ` + "`" + `yaml:",inline" mapstructure:",squash"` + "`" + `
    // 下方加入配置

}

func (c *Config) Show() {
	fmt.Printf("%+v\n", *c)

	fmt.Println("===== Global Config =====")
	show(reflect.ValueOf(c).Elem(), reflect.TypeOf(*c))
	fmt.Println("==============================")
}

func show(objVal reflect.Value, objType reflect.Type) {
	if objVal.Kind() == reflect.Slice {
		objVal.Len()
		for i := 0; i < objVal.Len(); i++ {
			b := objVal.Index(i).Kind()
			if b == reflect.Struct || b == reflect.Slice {
				show(objVal.Index(i), objVal.Index(i).Type())
			}
		}
		return
	}
	for i := 0; i < objVal.NumField(); i++ {
		field := objVal.Field(i)
		typeField := objType.Field(i)

		fmt.Printf("%s: %v\n", typeField.Name, field.Interface())
		b := field.Kind()
		if b == reflect.Struct || b == reflect.Slice {
			fmt.Printf("===== %s Desc =====\n", typeField.Name)
			show(field, field.Type())
			fmt.Println("==============================")
			fmt.Println("\n")
		}
	}
}`)
}
