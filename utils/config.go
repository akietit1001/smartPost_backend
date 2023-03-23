package utils

import "github.com/spf13/viper"

type Config struct {
	DbHost              string `mapstructure:"DB_HOST"`
	DbPort              string `mapstructure:"DB_PORT"`
	DbName              string `mapstructure:"DB_NAME"`
	DbUser              string `mapstructure:"DB_USER"`
	DbPassword          string `mapstructure:"DB_PASSWORD"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
