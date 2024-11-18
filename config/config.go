package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/diki-haryadi/ztools/constant"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	App              AppConfig
	Grpc             GrpcConfig
	Http             HttpConfig
	Postgres         PostgresConfig
	SampleExtService GrpcConfig
	Kafka            KafkaConfig
	Sentry           SentryConfig
}

var BaseConfig *Config

type AppConfig struct {
	AppEnv      string `json:"app_env" envconfig:"APP_ENV"`
	AppName     string `json:"app_name" envconfig:"APP_NAME"`
	ConfigOauth ConfigOauth
}

// Config stores all configuration options
type ConfigOauth struct {
	Oauth         OauthConfig
	Session       SessionConfig
	IsDevelopment bool
}

// OauthConfig stores oauth service configuration options
type OauthConfig struct {
	AccessTokenLifetime  int `json:"access_token_lifetime" envconfig:"OAUTH_ACCESS_TOKEN_LIFETIME"`
	RefreshTokenLifetime int `json:"refresh_token_lifetime" envconfig:"OAUTH_REFRESH_TOKEN_LIFETIME"`
	AuthCodeLifetime     int `json:"auth_code_lifetime" envconfig:"OAUTH_AUTH_CODE_LIFETIME"`
}

// SessionConfig stores session configuration for the web app
type SessionConfig struct {
	Secret string `json:"secret" envconfig:"SESSION_SECRET"`
	Path   string `json:"path" envconfig:"SESSION_PATH"`
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge int `json:"max_age" envconfig:"SESSION_MAX_AGE"`
	// When you tag a cookie with the HttpOnly flag, it tells the browser that
	// this particular cookie should only be accessed by the server.
	// Any attempt to access the cookie from client script is strictly forbidden.
	HTTPOnly bool `json:"http_only" envconfig:"SESSION_HTTP_ONLY"`
}

type PostgresConfig struct {
	Host            string `json:"host" envconfig:"PG_HOST"`
	Port            string `json:"port" envconfig:"PG_PORT"`
	User            string `json:"user" envconfig:"PG_USER"`
	Pass            string `json:"pass" envconfig:"PG_PASS"`
	DBName          string `json:"db_name" envconfig:"PG_DB"`
	MaxConn         int    `json:"max_conn" envconfig:"PG_MAX_CONNECTIONS"`
	MaxIdleConn     int    `json:"max_idle_conn" envconfig:"PG_MAX_IDLE_CONNECTIONS"`
	MaxLifeTimeConn int    `json:"max_life_time_conn" envconfig:"PG_MAX_LIFETIME_CONNECTIONS"`
	SslMode         string `json:"ssl_mode" envconfig:"PG_SSL_MODE"`
}
type GrpcConfig struct {
	Port int    `json:"port" envconfig:"GRPC_PORT"`
	Host string `json:"host" envconfig:"GRPC_HOST" `
}

type HttpConfig struct {
	Port int    `json:"port" envconfig:"HTTP_PORT"`
	Host string `json:"host" envconfig:"HTTP_HOST"`
}

type KafkaConfig struct {
	Enabled       bool     `json:"enabled" envconfig:"KAFKA_ENABLED"`
	LogEvents     bool     `json:"log_events" envconfig:"KAFKA_LOG_EVENTS"`
	ClientId      string   `json:"client_id" envconfig:"KAFKA_CLIENT_ID"`
	ClientGroupId string   `json:"client_group_id" envconfig:"KAFKA_CLIENT_GROUP_ID"`
	ClientBrokers []string `json:"client_brokers" envconfig:"KAFKA_CLIENT_BROKERS"`
	Topic         string   `json:"topic" envconfig:"KAFKA_TOPIC"`
}

type SentryConfig struct {
	Dsn string `json:"dsn" envconfig:"SENTRY_DSN"`
}

func init() {
	//BaseConfig = newConfig()
	//BaseConfig = LoadConfig()
}

func LoadConfig() *Config {
	_ = godotenv.Overload()
	var configLoader Config

	if err := envconfig.Process("BaseConfig", &configLoader); err != nil {
		log.Printf("error load config: %v", err)
	}

	BaseConfig = &configLoader
	spew.Dump(configLoader)
	return &configLoader
}

func IsDevEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvDev
}

func IsProdEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvProd
}

func IsTestEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvTest
}
