package main

import (
	"fmt"
	"telegraph/apis"
	"telegraph/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apis.RegisterApis(r)
	url := fmt.Sprintf("%s:%d", config.GetConfigs().Host, config.GetConfigs().Port)
	r.Run(url)
}
