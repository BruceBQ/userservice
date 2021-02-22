package config

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"userservice/model"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func newViper(allowEnvironmentOverrides bool) *viper.Viper {
	v := viper.New()

	v.SetConfigType("json")

	if allowEnvironmentOverrides {
		v.SetEnvPrefix("")
		v.AutomaticEnv()
	}

	return v
}

func unmarsharlConfig(r io.Reader, allowEnvironmentOverrides bool) (*model.Config, map[string]string, error) {
	configData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read")
	}

	var rawConfig model.Config
	if err = json.Unmarshal(configData, &rawConfig); err != nil {
		return nil, nil, err
	}
	rawConfig.SetDefaults()

	v := newViper(allowEnvironmentOverrides)
	v.ReadConfig(bytes.NewBuffer(configData))

	var config model.Config
	unmarshalErr := v.Unmarshal(&config)

	envConfig := GetEnvironment()
	mongoDbAddress := os.Getenv("MONGODB_ADDRESS")
	mongoDbUsername := os.Getenv("MONGODB_USERNAME")
	mongoDbPassword := os.Getenv("MONGODB_PASSWORD")
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	cameraServiceAddress := os.Getenv("CAMERA_SERVICE_ADDRESS")
	notificationServiceAddress := os.Getenv("NOTIFICATION_SERVICE_ADDRESS")

	if mongoDbAddress != "" {
		config.MongoSettings.Address = &mongoDbAddress
		config.MongoSettings.Username = &mongoDbUsername
		config.MongoSettings.Password = &mongoDbPassword
	}

	if redisAddress != "" {
		config.RedisSettings.Address = &redisAddress
		config.RedisSettings.Password = &redisPassword
	}

	if cameraServiceAddress != "" {
		config.GRPCSettings.CameraServiceAddress = &cameraServiceAddress
	}

	if notificationServiceAddress != "" {
		config.GRPCSettings.NotificationServiceAddress = &notificationServiceAddress
	}

	return &config, envConfig, unmarshalErr
}
