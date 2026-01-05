package config

type ServiceConfig struct {
	Server struct {
		Port uint16 `yaml:"port"`
	}

	InfluenceDB PgConfig `yaml:"influencedb"`
}
