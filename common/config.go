package common

import (
	"os"

	"github.com/spf13/viper"
)

//Config ..
type Config struct {
	conf *viper.Viper
}

//NewConfig Returns a value of type Viper with config file set
func NewConfig() *Config {
	c := Config{conf: viper.New()}
	c.conf.SetConfigName("config-default")
	c.conf.AddConfigPath(os.Getenv("GOPATH") + "/src/github.com/sjsu-achilis/achlibgo/common")
	c.conf.SetConfigType("json")
	err := c.conf.ReadInConfig()
	if err != nil {
		Log().Error("error Reading config files:  ", err)
	}

	return &c
}

//GetFromConfig ...
func GetFromConfig(key string) interface{} { return c.GetFromConfig(key) }

//GetFromConfig ...
func (c *Config) GetFromConfig(key string) interface{} {
	return c.conf.Get(key)
}

//SetConfigFile ...
func SetConfigFile(file string, path string, ext string) { c.SetConfigFile(file, path, ext) }

//SetConfigFile ...
func (c *Config) SetConfigFile(file string, path string, ext string) {
	c.conf.SetConfigName(file)
	c.conf.SetConfigType(ext)
	c.conf.AddConfigPath(path)
	c.conf.MergeInConfig()
}
