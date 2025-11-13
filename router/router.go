package router

import (
	"gin-demo/router/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)

}
