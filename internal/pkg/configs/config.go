package configs

import (
	"flag"
	"io"
	"sync"
	"os"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"github.com/pkg/errors"
)
const (
	configFileKey = "configFile"
	defaultConfigFile = ""
	configFileUsage = "Path to the configuration file"
)

var (
	once sync.Once
	cachedConfig *AppConfig
)

type ClientConfig struct {
	ClientName string `mapstructure:"clientName"`
	LogLevel string `mapstructure:"logLevel"`
	ServerAddress string `mapstructure:"serverAddress"`
}

type DatabaseConfig struct {
	Dbname string `mapstructure:"dbname"`
	Schema string `mapstructure:"schema"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	LogMode bool `mapstructure:"logMode"`
	SslMode string `mapstructure:"sslMode"`
	Connection ConnectionPool `mapstructure:"ConnectionPool"`
	MigrationPath string `mapstructure:"migrationPath"`
}

type ConnectionPool struct {
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxIdleTime int `mapstructure:"maxIdleTime"`
	MaxLifetime int `mapstructure:"maxLifetime"`
	TimeOut int `mapstructure:"timeOut"`
}

type ServerConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	LogLevel string `mapstructure:"logLevel"`
}

type AppConfig struct {
	ClientConfig ClientConfig `mapstructure:"clientConfig"`
	DBConfig DatabaseConfig `mapstructure:"databaseConfig"`
	ServerConfig ServerConfig `mapstructure:"serverConfig"`
}

func ProvideAppConfig() (c *AppConfig, err error) {
	once.Do(func() {
		var configFile string
		flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
		flag.Parse()

	var configReader io.ReadCloser
	configReader, err = os.Open(configFile)
	defer configReader.Close() //nolint:errcheck
	if err != nil {
		err = errors.Wrap(err, "failed to open config file")
		return
	}
	c, err = LoadConfig(configReader)
	if err != nil {
		err = errors.Wrap(err, "failed to load config")
		return
	}
	cachedConfig = c
	})

	return cachedConfig, err
}

func LoadConfig(r io.Reader) (*AppConfig, error) {
	var  appConfig AppConfig
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	keysToEnvironmentVariables := map[string]string{
		"app.port": "APP_PORT",
		"db.name": "DB_NAME",
		"db.user": "DB_USER",
		"db.password": "DB_PASSWORD",
		"db.host": "DB_HOST",
		"db.port": "DB_PORT",
	}	
	err := bind(keysToEnvironmentVariables)
	if err != nil {
		return nil, err
	}
	if err := viper.ReadConfig(r); err != nil {
		return nil, errors.Wrap(err, "failed to read load file")
	}
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}
	return &appConfig, nil
}

func bind(keysToEnvironmentVariables map[string]string) error {
	var bindError error
	for key, envVar := range keysToEnvironmentVariables {
		if err := viper.BindEnv(key, envVar); err != nil {
			bindError = multierror.Append(bindError, err)
		}
	}
	return bindError
}