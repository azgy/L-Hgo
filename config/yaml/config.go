package yaml

import (
	"github.com/go-gypsy/yaml"
	"fmt"
)

var config map[string]string

func init() {
	conf, err := yaml.ReadFile("src/doc/conf.yaml")
	if err != nil {
		fmt.Println("读取配置文件conf.yaml出错," , err)
	}

	config = make(map[string]string)
	value, err := conf.Get("mysql.url")
	if err == nil {
		config["mysql.url"] = value
	}
}

func GetConfig(key string) string {
	return config[key]
}
