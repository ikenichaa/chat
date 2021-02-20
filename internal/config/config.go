package config

import (
	"log"
	"os"
	"strconv"
)

//LocalEnvironment...
const (
	LocalEnvironment = "local"
)

//Config ...
type Config struct {
	Env                                                 string
	ServerPort                                          int
	MysqlUsername, MysqlPassword, MysqlDB, MysqlAddress string
	MysqlPort                                           int
}

//Init ...
func Init() (_ *Config, err error) {
	var cfg = &Config{
		Env:           GetEnv("ENVIRONMENT", ""),
		ServerPort:    GetEnvInt("PORT", 9999),
		MysqlUsername: GetEnv("MYSQL_USERNAME", ""),
		MysqlPassword: GetEnv("MYSQL_PASSWORD", ""),
		MysqlDB:       GetEnv("MYSQL_DATABASE", ""),
		MysqlPort:     GetEnvInt("MYSQL_PORT", 3306),
		MysqlAddress:  GetEnv("MYSQL_ADDRESS", ""),
	}
	return cfg, nil
}

//GetEnv ...
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// GetEnvInt return to int
func GetEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("cannot convert port to int")
	}
	return v
}
