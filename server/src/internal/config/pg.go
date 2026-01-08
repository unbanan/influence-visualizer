package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type PgConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DbName          string        `yaml:"db_name"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func (c PgConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		strconv.Itoa(c.Port),
		c.User,
		os.Getenv(c.Password),
		c.DbName,
		c.SSLMode,
	)
}
