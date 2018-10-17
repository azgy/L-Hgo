package redis

import (
	"github.com/garyburd/redigo/redis"
	"errors"
)

var client redis.Conn

func conn() {
	client = GetConn()
}

func Set(key, value string) (error) {
	conn()
	defer client.Close()
	_, err := client.Do("SET", key, value)
	if err != nil {
		return errors.New("写入key=" + key + ", value=" + value + "的缓存失败")
	}

	return nil
}

func Get(key string) (string, error) {
	conn()
	defer client.Close()
	value, err := redis.String(client.Do("GET", key))
	if err != nil {
		return "", errors.New("未获取到key=" + key + "的缓存")
	}

	return value, nil
}

func Del(key string) (error) {
	conn()
	defer client.Close()
	_, err := client.Do("DEL", key)
	if err != nil {
		return errors.New("删除key=" + key + "的缓存失败")
	}

	return nil
}
