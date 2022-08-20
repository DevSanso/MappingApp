package main


import (
	"fmt"

	"common/kafka"
	"common/config"
)

type Config struct {
	DB struct {
		Host string `json:"host"`
		Port int `json:"port"`
		User string `json:"userName"`
		Password string `json:"password"`
		DbName string `json:"dbName"`
	} `json:"dbConfig"`

	Kafka *kafka.ResovlerConfig `json:"kafkaConfig"`
}

func (c *Config)DbSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.DbName)
}

func parsingArgs() (*Config,error) {
	var cfg = new(Config)
	cfg.Kafka = kafka.NewResolverConfig()
	err := config.Parsing(cfg)
	return cfg,err
}