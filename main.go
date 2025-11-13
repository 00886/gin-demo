package main

import (
	"gin-demo/config"
	"gin-demo/middlerware"
	"gin-demo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlerware.JwtAuth)
	router.RegisterRouter(r)

	r.Run(config.Port)

}
