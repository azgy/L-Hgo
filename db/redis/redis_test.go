package redis

import (
	"testing"
)

func TestConnRedis(t *testing.T) {
	/*c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis失败")
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "name", "hui")
	if err != nil {
		fmt.Println("SET error")
	}

	name, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("GET error")
		return
	}
	fmt.Println("name:" + name)*/

	//Set("sex", "boy")
}
