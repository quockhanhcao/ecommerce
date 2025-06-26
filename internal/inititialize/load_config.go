package inititialize

import (
	"fmt"

	"github.com/quockhanhcao/ecommerce/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// configuration struct
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unmarshal config error: %v", err)
	}
}
