package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"runtime"
)

type ConfigDataBaseNews struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5433"`
	DBname   string `yaml:"dbname" env:"DB_NAME" env-default:"admin"`
	Username string `yaml:"username" env:"USERNAME" env-default:"admin"`
	Password string `yaml:"password" env:"PASSWORD"`
	SSLmode  string `yaml:"sslmode" env:"SSL_MODE" env-default:"disable"`
}

func Config() *ConfigDataBaseNews {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
	//Вызов внешних переменных, отткуда вызов и с какой строки
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Fatal("Program didn't read config file")
	} else {
		logrus.Info("Program read config file successful!")
		cfg := ConfigDataBaseNews{
			Port:     viper.GetString("db.port"),
			Host:     viper.GetString("db.host"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			DBname:   viper.GetString("db.dbname"),
			SSLmode:  viper.GetString("db.sslmode"),
		}
		return &cfg
	}
	return nil
}
