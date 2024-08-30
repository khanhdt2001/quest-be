package util

import "github.com/spf13/viper"

type Config struct {
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
	SERVER_ADDRESS   string `mapstructure:"SERVER_ADDRESS"`
	MAIL_PASSWORD    string `mapstructure:"MAIL_PASSWORD"`
}

var Default Config

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("../.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&Default)
	return
}
