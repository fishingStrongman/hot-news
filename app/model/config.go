package model

type AppConfig struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}
