package auth

import (
	"gin-demo/controller/auth"
	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

func RegisterSubRouter(group *gin.RouterGroup) {
	authGroup := group.Group("/auth")
	login(authGroup)
	logout(authGroup)

}
