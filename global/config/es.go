package config

import (
	"fmt"
	"github.com/karldoenitz/Tigo/logger"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"github.com/spf13/viper"
	"os"
)

var esClient *elastic.Client

func init() {
	initES()
}

func initES() {
	configPath := GetConfigPath()
	esConfig := viper.New()
	esConfig.AddConfigPath(configPath)
	esConfig.SetConfigName(viper.GetString("dbConf"))
	if err := esConfig.ReadInConfig(); err != nil {
		logger.Error.Println("Can't find the DB config file! \n")
		os.Exit(2)
	}

	runMode := GetRunMode()
	var esConf map[string]string
	switch runMode {
	case DevMode:
		esConf = esConfig.GetStringMapString("dev.es")
	case TestMode:
		esConf = esConfig.GetStringMapString("test.es")
	case ProductionMode:
		esConf = esConfig.GetStringMapString("production.es")
	default:
		logger.Error.Println("Undefined run mode !\n")
		os.Exit(2)
	}
	var ip, port string
	ip = esConf["ip"]
	port = esConf["port"]
	addr := fmt.Sprintf("http://%s:%s", ip, port)
	client, err := elastic.NewClientFromConfig(&config.Config{
		URL: addr,
	})
	if err != nil {
		logger.Error.Printf("init es[address => %s] client failed => (%s)", addr, err.Error())
		panic(err)
	}
	esClient = client
}

func GetESClient() (client *elastic.Client) {
	return esClient
}
