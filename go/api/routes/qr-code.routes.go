package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
)

func GenerateQRCode(group *gin.RouterGroup) {
	qr := group.Group("/qrcode")
	{
		qr.GET("/generate", func(ctx *gin.Context) {
			params := ctx.Request.URL.Query()
			redirect_url := params.Get("redirect_url")
			if redirect_url == "" {
				ctx.JSON(http.StatusBadRequest, utils.FailedResponse("no redirect_url provided"))
				return
			}
			point := params.Get("point")
			if point == "" {
				ctx.JSON(http.StatusBadRequest, utils.FailedResponse("no point provided"))
				return
			}
			access_token := params.Get("access_token")
			if access_token == "" {
				ctx.JSON(http.StatusBadRequest, utils.FailedResponse("no access_token provided"))
				return
			}
			var r = redirect_url + "&reward_point=" + point + "&access_token=" + access_token
			response, err := http.Get("https://api.qrserver.com/v1/create-qr-code/?size=512x512&data=" + r)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.FailedResponse("Error while generating QR code"))
				return
			} else {
				ctx.Writer.Header().Set("Content-Type", "image/png")
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					//   log.Fatalln(err)
				}
				ctx.Writer.Write(body)

			}
			defer response.Body.Close()
		})
	}
}
