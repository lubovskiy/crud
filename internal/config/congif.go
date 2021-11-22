package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const(
	configPath = "config.json"
)

func init() {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetConfigConnection() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetInt(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	return  fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
}
