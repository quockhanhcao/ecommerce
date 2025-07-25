package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// configuration struct
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("unmarshal config error: %v", err)
	}
    fmt.Println("Config port: ", config.Server.Port)
    for _, db := range config.Databases {
        fmt.Printf("Database User: %s, Password: %s, Host: %s\n", db.User, db.Password, db.Host)
    }
}
