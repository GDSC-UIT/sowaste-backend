package config

import "os"

func GetDBConfig() *DBConfig {
	// LoadEnv()

	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	return &DBConfig{
		Connection: "mongobd",
		DbName:     dbName,
		Username:   dbUsername,
		Password:   dbPassword,
	}
}
