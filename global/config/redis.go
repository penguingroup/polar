package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/karldoenitz/Tigo/logger"
	"github.com/spf13/viper"
)

var redisPool *redis.Pool

func init() {
	initRedisPool()
}

func initRedisPool() {
	configPath := GetConfigPath()
	redisConfig := viper.New()
	redisConfig.AddConfigPath(configPath)
	redisConfig.SetConfigName(viper.GetString("dbConf"))
	if err := redisConfig.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't find the DB config file! \n")
		os.Exit(2)
	}

	runMode := GetRunMode()
	var redisConf map[string]string
	switch runMode {
	case DevMode:
		redisConf = redisConfig.GetStringMapString("dev.redis")
	case TestMode:
		redisConf = redisConfig.GetStringMapString("test.redis")
	case ProductionMode:
		redisConf = redisConfig.GetStringMapString("production.redis")
	default:
		fmt.Fprintf(os.Stderr, "Undefined run mode !\n")
		os.Exit(2)
	}
	var ip, port string
	if serviceName, err := redisConf["name_service"]; err {
		//从名字服务获取服务ip和port
		fmt.Println("[nameService:" + serviceName + "] => " + ip + ":" + port)
		logger.Error.Printf("[nameService:" + serviceName + "] => " + ip + ":" + port)
	} else {
		ip = redisConf["ip"]
		port = redisConf["port"]
	}
	maxIdle, _ := strconv.Atoi(redisConf["maxidle"])
	timeout, _ := strconv.Atoi(redisConf["timeout"])
	pwd, _ := redisConf["pwd"]
	auth, _ := redisConf["auth"]
	dbNo, _ := strconv.Atoi(redisConf["select"])
	addr := fmt.Sprintf("%s:%s", ip, port)
	redisPool = produceRedisPool(addr, maxIdle, timeout, auth, dbNo, pwd)
	if viper.GetString("isFlushAll") == "yes" {
		conn := redisPool.Get()
		defer conn.Close()
		conn.Do("FLUSHALL")
	}
}

func GetRedisPool() *redis.Pool {
	if redisPool != nil {
		return redisPool
	}
	initRedisPool()
	return redisPool
}

func produceRedisPool(addr string, maxIdle, timeout int, auth string, dbNo interface{}, pwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr, redis.DialPassword(pwd), redis.DialDatabase(dbNo.(int)))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
