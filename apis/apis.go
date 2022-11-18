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
	r.GET("/createPage", api_handler.CreatePage)
	r.POST("/createPage", api_handler.CreatePage)
	r.GET("/getPage/:path", api_handler.GetPage)
	r.POST("/getPage/:path", api_handler.GetPage)
	r.GET("/editPage/:path", api_handler.EditPage)
	r.POST("/editPage/:path", api_handler.EditPage)
	r.GET("/getPageList", api_handler.GetPageList)
	r.POST("/getPageList", api_handler.GetPageList)
	r.GET("/editAccountInfo", api_handler.EditAccountInfo)
	r.POST("/editAccountInfo", api_handler.EditAccountInfo)
	r.GET("/getViews/:path", api_handler.GetViews)
	r.POST("/getViews/:path", api_handler.GetViews)
	return r
}
