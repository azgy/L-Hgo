package mysql

import (
	"database/sql"
	"config/ini"
	_ "github.com/go-sql-driver/mysql" //加载驱动
	"github.com/go-xorm/xorm"
	"logger"
	"errors"
)

var SqlDB *sql.DB

func Client() (*xorm.Engine, error) {
	url, err := ini.GetConfig("Mysql", "url")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	maxIdleConns, err := ini.GetConfigToInt("Mysql", "maxIdleConns")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	maxOpenConns, err := ini.GetConfigToInt("Mysql", "maxOpenConns")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	showSql, err := ini.GetConfigToBool("Mysql", "showSql")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		logger.Error("连接数据库失败,%v", err.Error())
		return nil, errors.New("failed to connect to database")
	}
	engine.SetMaxOpenConns(maxIdleConns)
	engine.SetMaxIdleConns(maxOpenConns)
	engine.ShowSQL(showSql)

	err = engine.Ping()
	if err != nil {
		logger.Error("连接数据库失败,%v", err.Error())
		return nil, errors.New("failed to connect to database")
	}

	return engine, nil
}
