package auth

import (
	"gin-demo/config"
	"gin-demo/util/jwt"
	"gin-demo/util/logging"
	"gin-demo/util/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	userInfo := UserInfo{}

	if err := r.ShouldBindJSON(&userInfo); err != nil {
		response.Ng(r, err.Error())
		return
	}
	logging.Debug(logrus.Fields{
		"用户名": userInfo.Username,
		"密码":  userInfo.Password,
	}, "开始验证登录信息")

	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		token, err := jwt.GenerToken(userInfo.Username)
		if err != nil {
			response.Ng(r, err.Error())
			return
		}

		response.Ok(r, token)
	} else {
		response.Ng(r, "登录失败，请输入正确的用户名和密码")
	}
}

func Logout(r *gin.Context) {
	response.Ok(r, "登出成功")
}
