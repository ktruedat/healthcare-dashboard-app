package config

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
