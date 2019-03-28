package config

import "github.com/yuwenhui/gopkg/dbutil"

type BaseConfig struct {
	DBAddress *dbutil.MySQLAuth `yaml:"db_address"`
}

func (a *BaseConfig) Verify() error {
	return nil
}
