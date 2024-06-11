package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	DB     DBConfig
}

type Server struct {
	HostName        string `yaml:"hostname"`
	Port            int    `yaml:"port"`
	CorsAllowOrigin string `yaml:"corsAllowOrigin"`
}

type DBConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
}

var config Config

func init() {
	viper.SetDefault("server.hostname", "127.0.0.1")
	viper.SetDefault("server.port", 9000)
	viper.SetDefault("server.corsAllowOrigin", "http://localhost:3000")

	viper.SetDefault("db.name", "training")
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.host", "db")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.port", 3306)
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// UnmarshalしてCにマッピング
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("config file Unmarshal error", err)
		panic(err)
	}
	return &config
}
