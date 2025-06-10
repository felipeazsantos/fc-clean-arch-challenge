package configs

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"db_driver"`
	DBHost            string `mapstructure:"db_host"`
	DBPort            string `mapstructure:"db_port"`
	DBUser            string `mapstructure:"db_user"`
	DBPass            string `mapstructure:"db_pass"`
	DBName            string `mapstructure:"db_name"`
	WebServerPort     string `mapstructure:"web_server_port"`
	GrpcServerPort    string `mapstructure:"grpc_server_port"`
	GraphqlServerPort string `mapstructure:"graphql_server_port"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ConnectDB(cfg *conf) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
