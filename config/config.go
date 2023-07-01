package config

import "github.com/spf13/viper"

type configRabbitMq struct {
	HOST string `mapstructure:"HOST"`
	USER string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
}

func LoadConfigs(path string) (*configRabbitMq){
	var config *configRabbitMq
	viper.SetConfigName("rabbit_confg")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
}


// var cfg *conf
// viper.SetConfigName("app_config")
// viper.SetConfigType("env")
// viper.AddConfigPath(path)
// viper.SetConfigFile(".env")
// viper.AutomaticEnv()
// err := viper.ReadInConfig()
// if err != nil {
// 	panic(err)
// }
// err = viper.Unmarshal(&cfg)
// if err != nil {
// 	panic(err)
// }
// cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)
// return cfg, nil