package main

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/config"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
)

func init() {
	config.LoadEnv()
}

func main() {
	config.GetDBConfig()
	database.Client.ConnectDb()
}
