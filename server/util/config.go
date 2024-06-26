package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBString               string        `mapstructure:"dbString"`
	PORT                   string        `mapstructure:"port"`
	MAILGUN_DOMAIN         string        `mapstructure:"mailgunDomain"`
	MailgunAPIKey          string        `mapstructure:"mailgunApiKey"`
	Access_Token_Duration  time.Duration `mapstructure:"accessTokenDuration"`
	Refresh_Token_Duration time.Duration `mapstructure:"refreshTokenDuration"`
	SecretKeyHex           string        `mapstructure:"secretKeyHex"`
	PublicKeyHex           string        `mapstructure:"publicKeyHex"`
	DOSecretKey            string        `mapstructure:"doSecret"`
	DOSpacesKey            string        `mapstructure:"doSpaces"`
	MailJetApiKey          string        `mapstructure:"mailJetApiKey"`
	MailJetSecretKey       string        `mapstructure:"mailJetSecretKey"`
	VultrAccessKey         string        `mapstructure:"vultrAccessKey"`
	VultrSecretKey         string        `mapstructure:"vultrSecretKey"`
	VultrHostname          string        `mapstructure:"vultrHostname"`
	BlackBlazeSecretKey    string        `mapstructure:"blackBlazeSecretKey"`
	BlackBlazeKeyId        string        `mapstructure:"blackBlazeKeyId"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	viper.SetConfigType("env")
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
