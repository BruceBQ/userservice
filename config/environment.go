package config

import "os"

type EnvConfig struct {
	MongoDBAddress       string
	MongoDBUsername      string
	MongoDBPassword      string
	CameraServiceAddress string
	RedisServiceAddress  string
}

func GetEnvironment() map[string]string {
	env := make(map[string]string)

	env["MONGODB_ADDRESS"] = os.Getenv("MONGODB_ADDRESS")
	env["MONGODB_USERNAME"] = os.Getenv("MONGODB_USERNAME")
	env["MONGODB_PASSWORD"] = os.Getenv("MONGODB_PASSWORD")
	env["CAMERA_SERVICE_HOST"] = os.Getenv("CAMERA_SERVICE_HOST")
	env["CAMERA_SERVICE_GRPC_PORT"] = os.Getenv("CAMERA_SERVICE_GRPC_PORT")
	env["REDIS_SERVICE_ADDRESS"] = os.Getenv("REDIS_SERVICE_HOST")
	return env
}
