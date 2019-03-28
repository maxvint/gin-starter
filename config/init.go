package config

import (
	"github.com/jinzhu/gorm"
	"github.com/yuwenhui/gopkg/goutil"
	"github.com/yuwenhui/gopkg/logs"
)

var (
	AppConfig BaseConfig
	DBConfig  *gorm.DB
)

func init() {
	err := goutil.LoadConfigFor(&AppConfig, "app.yaml")
	if err != nil {
		logs.Panicf("Load App Config error: %v", err)
	}

	DBConfig, err = AppConfig.DBAddress.DB()
	if err != nil {
		logs.Panicf("DBAddress error: %v", err)
	}

	// Redis = AppConfig.RedisAPI.RedisClient()
}
