package config

import (
  "github.com/spf13/viper"
  logger "CourseService/logger"
)

type Config struct {
	Port  string  `mapstructure:"PORT"`
}

func LoadConfig(path string)(config Config, err error){

  var standardLogger = logger.NewLogger()

  standardLogger.Info("Loading config")
  
  viper.AddConfigPath(path)
  viper.SetConfigName("app")
  viper.SetConfigType("env")

  viper.AutomaticEnv()

  err=viper.ReadInConfig()

  if err!=nil{
    standardLogger.Error("Config read error:", err)
	  return
  }

  viper.Unmarshal(&config)
  return
}