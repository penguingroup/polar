package config

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type RunMode int

const (
	DevMode RunMode = iota
	TestMode
	ProductionMode
)

func init() {
	pflag.StringP("configPath", "p", "./config/", "配置文件夹路径")
	pflag.StringP("configFile", "c", "course.yaml", "主配置文件名称(course.yaml)")
	pflag.StringP("mode", "m", "dev", "程序运行模式[dev, test, production]")
	pflag.StringP("isFlushAll", "f", "no", "是否清空redis缓存")
	pflag.String("dbConf", "db", "自定义程序使用的数据库相关配置文件")
	pflag.String("serverConf", "server.yaml", "程序启动的配置文件")
	pflag.Usage = usage
	pflag.ErrHelp = errors.New("")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	initConfig()
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`Paper Pen version: 0.0.1
	Usage: Planet [Options]

    Options:
          -c, --configFile string   主配置文件名称(charon.yaml) (default "charon.yaml")
          -p, --configPath string   配置文件夹路径 (default "./config/")
          -d, --debug               是否开启调试模式
          -m, --mode string         程序运行模式[dev, test, production] (default "dev")
              --IP string           程序侦听的IP地址 (default "127.0.0.1")"")
              --port int            程序侦听的端口 (default 8080)
              --http_proxy string   http代理
              --dbconf string       自定义程序使用的数据库配置文件名称,注意不用加后缀名称!(default "db")
              --log_name string     日志文件名称!(default "charon")
    `)
}

func initConfig() {
	cfgFile := viper.GetString("configPath") + viper.GetString("configFile")
	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't find the config file in %s ! \n", cfgFile)
		os.Exit(2)
	}
}

func GetConfigPath() string {
	return viper.GetString("configPath")
}

func GetRunMode() RunMode {
	mode := viper.GetString("mode")
	switch mode {
	case "dev":
		return DevMode
	case "test":
		return TestMode
	case "production":
		return ProductionMode
	default:
		return DevMode
	}
}

func GetModeName() string {
	return viper.GetString("mode")
}

func GetServerConfig() string {
	return GetConfigPath() + viper.GetString("serverConf")
}
