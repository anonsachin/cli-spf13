package config

import (
	"fmt"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Config to read in for the emojis
type Config struct{
	Emojis map[string]string
}

// NewEmojiConfig to get the emojis for use in cli
func NewEmojiConfig() (*Config,error){
	home, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("Unable to get Home dir: %v",err)
	}
	v := viper.New()
	v.SetConfigName("essentials")
	v.AddConfigPath("emojis")
	homePath := path.Join(home,"emojis")
	v.AddConfigPath(homePath)
	err = v.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("Unable to load config for emojis: %v",err)
	}
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract emojis: %v",err)
	}
	return &config, nil
}