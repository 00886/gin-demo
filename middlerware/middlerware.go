package middlerware

import (
	"gin-demo/util/jwt"
	"gin-demo/util/logging"
	"gin-demo/util/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JwtAuth(r *gin.Context) {
	requestUrl := r.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		r.Next()
		return
	}
	if token := r.GetHeader("Authorization"); token != "" {
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Ng(r, "解析token失败")
			r.Abort()
			return
		}
		logging.Info(logrus.Fields{
			"用户名":  claims.Username,
			"过期时间": claims.ExpiresAt,
		}, "解析token成功")
	} else {
		response.Ng(r, "用户未登录，请登录用户")
		r.Abort()
	}
}
