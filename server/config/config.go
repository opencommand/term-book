package config

const DataPath = "data"

type Configure struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}
