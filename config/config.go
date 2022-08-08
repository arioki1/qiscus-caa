package config

import "github.com/kelseyhightower/envconfig"

type config struct {
	AppName          string `envconfig:"APP_NAME" default:"qiscus-caa"`
	Version          string `envconfig:"VERSION" default:"1.0.0"`
	Debug            bool   `envconfig:"DEBUG" default:"false"`
	Port             int    `envconfig:"PORT" default:"8000"`
	QismoBaseURL     string `envconfig:"QISMO_BASE_URL" default:"https://multichannel.qiscus.com"`
	QiscusAdminEmail string `envconfig:"QISCUS_ADMIN_EMAIL" required:"true"`
	QiscusAppId      string `envconfig:"QISCUS_APP_ID" required:"true"`
	QiscusSecretKey  string `envconfig:"QISCUS_SECRET_KEY" required:"true"`
	RedisURL         string `envconfig:"REDIS_URL" required:"true"`
}

type Config interface {
	GetAppName() string
	GetVersion() string
	GetDebug() bool
	GetPort() int
	GetQismoBaseURL() string
	GetQiscusAdminEmail() string
	GetQiscusAppId() string
	GetQiscusSecretKey() string
	GetRedisURL() string
}

func LoadConfig() (Config, error) {
	cfg := new(config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c config) GetAppName() string {
	return c.AppName
}

func (c config) GetVersion() string {
	return c.Version
}

func (c config) GetDebug() bool {
	return c.Debug
}

func (c config) GetPort() int {
	return c.Port
}

func (c config) GetQismoBaseURL() string {
	return c.QismoBaseURL
}
func (c config) GetQiscusAdminEmail() string {
	return c.QiscusAdminEmail
}
func (c config) GetRedisURL() string {
	return c.RedisURL
}
func (c config) GetQiscusAppId() string {
	return c.QiscusAppId
}

func (c config) GetQiscusSecretKey() string {
	return c.QiscusSecretKey
}
