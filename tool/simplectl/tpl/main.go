package tpl

func MainTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
package main

import "{{ .PkgName }}/cmd"

func main() {
	cmd.Execute()
}
`)
}

func RootTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
package cmd

import (
{{- if .Viper }}
	"fmt"{{ end }}
	"github.com/smallnest/rpcx/server"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/srpc"
	"os"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"{{ .PkgName }}/global"
{{- if .Viper }}
	"github.com/spf13/viper"{{ end }}
)

{{ if .Viper -}}
var cfgFile string
{{- end }}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{ .AppName }}",
	Short: "A brief description of your application",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.` + "`" + `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		global.InitSlog()
		global.InitSRPC()
		global.InitDB()
		global.LOG.Error({{ .AppName }}Serve(global.CONFIG.RPC, global.CONFIG.RpcService).Error())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
{{- if .Viper }}
	cobra.OnInitialize(initConfig)
{{ end }}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
{{ if .Viper }}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.{{ .AppName }}.yaml)")
{{ else }}
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.{{ .AppName }}.yaml)")
{{ end }}
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// {{ .AppName }}Serve starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func {{ .AppName }}Serve(rpc sconfig.RPC, rpcService sconfig.RpcService) error {
	s := server.NewServer()
	// 开启rpcx监控
	s.EnableProfile = true
	// 服务注册中心
	srpc.AddRegistryPlugin(s, rpc, rpcService)

	s.RegisterName("Example", new(service.Example), "")

	return s.Serve("tcp", fmt.Sprintf("%s:%s", rpcService.ServiceAddr, rpcService.Port))
}

{{ if .Viper -}}
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".{{ .AppName }}" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("{{ .AppName }}")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		viper.WatchConfig()

		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err := viper.Unmarshal(&global.CONFIG); err != nil {
				fmt.Println(err)
			}
		})

		if err := viper.Unmarshal(&global.CONFIG); err == nil {
			global.CONFIG.Show()
		} else {
			os.Exit(1)
		}
	}
}
{{- end }}
`)
}

func ConfigYamlTemplate() []byte {
	return []byte(`rpc-service:
  service-addr: 127.0.0.1
  port: 9001

rpc:
  register-type: etcdv3
  register-addr:
    - 127.0.0.1:23791
    - 127.0.0.1:23792
    - 127.0.0.1:23793
  base-path: local

redis:
  addr: 127.0.0.1:6379
  password:
  db: 0

db-list:
  - disabled: true # 是否启用
    type: mysql # 数据库的类型,目前支持mysql、pgsql
    alias-name: upms # 数据库的名称,注意: alias-name 需要在db-list中唯一
    path: 127.0.0.1
    port: 3306
    config: charset=utf8mb4&parseTime=True&loc=Local
    db-name: upms
    username: root
    password: root
    max-idle-conns: 10
    max-open-conns: 100
    log-mode: error
    log-zap: true`)
}
