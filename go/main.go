package main

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/config"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
	"github.com/gin-gonic/gin"
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
	api.Router.Current.LoadHTMLGlob("docs/*.html")
	api.Router.Current.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/html")
		// ctx.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		ctx.HTML(200, "index.html", gin.H{})
	})
	api.Router.Run()

}
