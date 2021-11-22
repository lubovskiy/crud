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
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	return  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
