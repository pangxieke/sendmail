package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pangxieke/sendmail/log"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	//Port    uint
	LogFile string
}

type RedisConfig struct {
	Address string
}

type MailConfig struct {
	Host   string
	Port   uint
	Sender string
	PWD    string
}

var (
	Env    string
	Server ServerConfig
	Redis  RedisConfig
	Mail   MailConfig
)

func Init(configPaths ...string) (err error) {
	if err := setup(configPaths...); err != nil {
		return err
	}
	if err := initRedis(); err != nil {
		return err
	}

	if err := initServer(); err != nil {
		return err
	}
	if err := initMailServer(); err != nil {
		return err
	}
	return
}

func setup(paths ...string) (err error) {
	Env = os.Getenv("GO_ENV")
	if "" == Env {
		Env = "dev"
	}
	godotenv.Load(".env." + Env)
	godotenv.Load()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err = viper.ReadInConfig()
	if err != nil {
		log.Info("Failed to read config file (but environment config still affected), err:", err)
		err = nil
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return
}

func initRedis() (err error) {
	Redis.Address = viper.GetString("redis.address")
	if Redis.Address == "" {
		return errors.New("redis.address should not be empty")
	}
	return
}

func initServer() (err error) {
	Server.LogFile = viper.GetString("server.log")
	//Server.Port = viper.GetUint("server.port")

	if Server.LogFile == "" {
		return errors.New("server.log should not be empty")
	}
	//if Server.Port == 0 {
	//	return errors.New("server.port should not be empty")
	//}

	return
}

func initMailServer() (err error) {
	Mail.Host = viper.GetString("mail.host")
	if Redis.Address == "" {
		return errors.New("mail.host should not be empty")
	}
	Mail.Port = viper.GetUint("mail.port")
	Mail.Sender = viper.GetString("mail.sender")
	Mail.PWD = viper.GetString("mail.pwd")
	return
}
