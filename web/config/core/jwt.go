package core

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
	"web/config"
)

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var jwtObj *jwt.GinJWTMiddleware

func init() {
	initJwt()
}

func Jwt() *jwt.GinJWTMiddleware {
	if jwtObj == nil {
		initJwt()
	}
	return jwtObj
}

func initJwt() {
	if jwtObj == nil {
		jwtObj = &jwt.GinJWTMiddleware{
			Realm:         config.GetConfig().Jwt.Name,
			Key:           []byte(config.GetConfig().Jwt.Secret),
			Timeout:       time.Duration(config.GetConfig().Jwt.Duration) * time.Second,
			TokenLookup:   "header:Authorization",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
			Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			//用于登录
			Authenticator: func(userID string, password string, c *gin.Context) (interface{}, bool) { //登录
				//查询数据库，并登录
				//测试用例
				if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
					return &User{
						UserName:  userID,
						LastName:  "Bo-Yi",
						FirstName: "Wu",
					}, true
				}
				return nil, false
			},
			//用于认证
			Authorizator: func(user interface{}, c *gin.Context) bool { //认证
				if v, ok := user.(string); ok && v == "admin" {
					return true
				}
				return false
			},
		}
	}
}
