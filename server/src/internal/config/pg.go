package config

import (
	"time"
)

type PgConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DbName          string        `yaml:"db_name"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func (c PgConfig) DSN() string {
	return "host=" + c.Host + " port=" + string(c.Port) + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DbName + " sslmode=disable"
}
