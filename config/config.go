package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL  string
	APIToken string
}

func LoadConfig() (*Config, error) {
	var conf = &Config{}

	v := viper.New()

	viper.AutomaticEnv()

	v.SetConfigName("conf")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	v.SetEnvPrefix("VTT")
	v.AutomaticEnv()

	v.SetDefault("BaseURL", "<baseurlhere>")
	v.SetDefault("APIToken", "")

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	conf.BaseURL = strings.TrimSpace(v.GetString("BaseURL"))
	conf.APIToken = strings.TrimSpace(v.GetString("APIToken"))

	return conf, nil
}
