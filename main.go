package main

import (
	"telegraph/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apis.RegisterApis(r)
	r.Run(":8080")
}
