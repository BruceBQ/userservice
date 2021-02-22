package model

const (
	MONGO_DEFAULT_URI      = "mongodb://localhost:27017/"
	MONGO_DEFAULT_DATABASE = "users"
)

type ServiceSettings struct {
	HTTPAddress               *string
	GRPCAddress               *string
	SessionLengthWebInDays    *int
	SessionLengthMobileInDays *int
}

type GRPCSettings struct {
	CameraServiceAddress       *string
	NotificationServiceAddress *string
}

type MongoSettings struct {
	Address  *string
	Username *string
	Password *string
	Database *string
}

type RedisSettings struct {
	Address  *string
	Password *string
	DB       *int
}

func (m *MongoSettings) SetDefaults() {

}

type LogSettings struct {
	KafkaBroker *string
}

type Config struct {
	ServiceSettings ServiceSettings
	MongoSettings   MongoSettings
	LogSettings     LogSettings
	GRPCSettings    GRPCSettings
	RedisSettings   RedisSettings
}

func (c *Config) SetDefaults() {

}
