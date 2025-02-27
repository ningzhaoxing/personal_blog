package initialize

import (
	"gin-blog/pkg/globals"
	"github.com/spf13/viper"
	"log"
)

var configPath = "configs/dev.yaml"

func InitConfig() (*globals.Config, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("config Load error %s \n", err.Error())
		return nil, err
	}

	var config globals.Config
	err = viper.Unmarshal(&config)
	log.Println("config结构体:", config)
	if err != nil {
		log.Printf("config bind error %s \n", err.Error())
		return nil, err
	}

	return &config, nil
}
