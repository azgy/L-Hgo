package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"config/ini"
	"logger"
)

var clientPool *redis.Pool

func Init() {
	addr, err := ini.GetConfig("Redis", "url")
	if err != nil {
		logger.Error(err.Error())
	}
	connectTimeout, err := ini.GetConfigToInt("Redis", "connectTimeout")
	if err != nil {
		logger.Error(err.Error())
	}
	readTimeout, err := ini.GetConfigToInt("Redis", "readTimeout")
	if err != nil {
		logger.Error(err.Error())
	}
	writeTimeout, err := ini.GetConfigToInt("Redis", "writeTimeout")
	if err != nil {
		logger.Error(err.Error())
	}
	maxIdle, err := ini.GetConfigToInt("Redis", "writeTimeout")
	if err != nil {
		logger.Error(err.Error())
	}
	maxActive, err := ini.GetConfigToInt("Redis", "writeTimeout")
	if err != nil {
		logger.Error(err.Error())
	}
	idleTimeout, err := ini.GetConfigToInt("Redis", "writeTimeout")
	if err != nil {
		logger.Error(err.Error())
	}

	clientPool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr,
				redis.DialConnectTimeout(time.Duration(connectTimeout) * time.Millisecond),
				redis.DialReadTimeout(time.Duration(readTimeout) * time.Millisecond),
				redis.DialWriteTimeout(time.Duration(writeTimeout) * time.Millisecond),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetConn() (redis.Conn) {
	return clientPool.Get()
}