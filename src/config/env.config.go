package config

import (
	"errors"
	"fmt"
	"go-gin-template/src/constant"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func loadEnvVariables(key string) error {
	var err error
	switch key {
	case constant.PRODUCTION:
		err = godotenv.Load(".production.env")
	case constant.DEVELOPMENT:
		err = godotenv.Load(".development.env")
	default:
		return errors.New("unable to find configuration for env=" + key)
	}

	if err != nil {
		return err
	}
	InitLogger(os.Getenv("LOGGER_FILE"))

	log.Println(fmt.Sprintf("Application is running on %s environment", key))
	return nil
}
func GetEnvironmentConfig() {
	env := strings.ToLower(os.Getenv("APP_ENV"))
	log.Info("Selected env = ", env)
	if err := loadEnvVariables(env); err != nil {
		log.Println(fmt.Sprintf("Failed to load env variables, err=%v", err))
		panic(err)
	}
	log.Info("Application Environment Running profile " + env)
}

