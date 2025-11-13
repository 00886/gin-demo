package config

import (
	"gin-demo/util/logging"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	Port          string
	JwtSignKey    string
	JwtExpireTime int64
	Username      string
	Password      string
)

func init() {
	logging.Info(nil, "开始加载程序配置")
	err := godotenv.Load()
	if err != nil {
		logging.Error(nil, ".env 文件加载出错")
	}

	//获取应用端口
	viper.SetDefault("PORT", ":8080")
	//获取日志级别
	viper.SetDefault("LOG_LEVEL", "info")
	//获取jwt加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "fanjiale")
	//获取jwt过期时间
	viper.SetDefault("JWT_EXPIRE_TIME", 3600)
	//获取用户名和密码
	viper.SetDefault("USERNAME", "guest")
	viper.SetDefault("PASSWORD", "guest")
	viper.AutomaticEnv()

	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	initLogConfig(logLevel)

}

// 配置程序的日志输出级别
func initLogConfig(logLevel string) {
	var level logrus.Level
	switch logLevel {
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warn":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	default:
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
	})

}
