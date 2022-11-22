package apis

import (
	"net/http"
	"telegraph/api_handler"

	"github.com/gin-gonic/gin"
)

func RegisterApis(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/createAccount", api_handler.CreateAccount)
	r.POST("/createAccount", api_handler.CreateAccount)
	r.GET("/getAccountInfo", api_handler.GetAccountInfo)
	r.POST("/getAccountInfo", api_handler.GetAccountInfo)
	r.GET("/revokeAccessToken", api_handler.RevokeAccessToken)
	r.POST("/revokeAccessToken", api_handler.RevokeAccessToken)
	return r
}
