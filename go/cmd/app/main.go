package main

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/config"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
)

func init() {
	//** load .env file anywhere from any directory **//
	config.LoadEnv()
}

func main() {
	//** Database connection **/
	config.GetDBConfig()
	database.Client.ConnectDb()
	defer database.Client.DisconnetDb()

	//** API connect (router) **/
	api.Init()
	api.Router.RoutersEstablishment()
	api.Router.Run()

}
