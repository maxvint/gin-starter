package config

import "github.com/yuwenhui/gopkg/dbutil"

type AppConfig struct {
	DBAddress *dbutil.MySQLAuth `yaml:"db_address"`
}

func (a *AppConfig) Verify() error {
	return nil
}
