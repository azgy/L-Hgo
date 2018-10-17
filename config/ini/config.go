package ini

import (
	cf "github.com/larspensjo/config"
	"flag"
	"fmt"
	"errors"
)

var CFG *cf.Config

func init() {
	var configFile = flag.String("configfile", "conf.ini", "General configuration file")
	flag.Parse()
	cfg, err := cf.ReadDefault(*configFile)
	if err != nil {
		fmt.Println("读取配置文件出错,", err)
	}

	CFG = cfg
}

func GetConfig(group, key string) (string, error) {
	value, err := CFG.String(group, key)
	if err != nil {
		fmt.Printf("读取配置项%s-%s出错\n", group, key)
		return "", errors.New("读取配置项" + group + "-" + key + "出错")
	}

	return value, nil
}

func GetConfigToInt(group, key string) (int, error) {
	value, err := CFG.Int(group, key)
	if err != nil {
		fmt.Printf("读取配置项%s-%s出错\n", group, key)
		return -1, errors.New("读取配置项" + group + "-" + key + "出错")
	}

	return value, nil
}

func GetConfigToBool(group, key string) (bool, error) {
	value, err := CFG.Bool(group, key)
	if err != nil {
		fmt.Printf("读取配置项%s-%s出错\n", group, key)
		return false, errors.New("读取配置项"+group+"-"+key+"出错")
	}

	return value, nil
}
