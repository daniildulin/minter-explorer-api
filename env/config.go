package env

import (
	"explorer-api/helpers"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.AutomaticEnv()

	viper.AddConfigPath("/etc/minter/") // path to look for the config file in
	viper.AddConfigPath(".")

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	//panic
	helpers.CheckErr(err)
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
