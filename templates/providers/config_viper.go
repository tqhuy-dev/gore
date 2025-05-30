package providers

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ConfigProviderWireDI = wire.NewSet(NewConfigProvider)

type IConfigProvider interface {
	GetConfig() *ConfigDto
}

type configProvider struct {
	config ConfigDto
}

func (c *configProvider) GetConfig() *ConfigDto {
	return &c.config
}

type OptionDB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database "`
}

type ConfigDto struct {
	Mongo OptionDB `mapstructure:"mongo"`
}

func NewConfigProvider() IConfigProvider {
	viper.SetConfigName("config") // tên file không có phần mở rộng
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config ConfigDto
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return &configProvider{
		config: config,
	}
}
