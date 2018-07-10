package router

import (
	"github.com/gin-gonic/gin"
	"web/api/v1/user"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	router(engine)
	return engine
}

//路由
func router(router *gin.Engine) {
	//v1 := router.Group("v1")
	//{
	//	v1.GET("/", func(context *gin.Context) {
	//		shopAdmin := new(models.ShopAdmin)
	//		core.DB.ID(12).Get(shopAdmin)
	//		//cmd, err := core.Redis.Ping().Result()
	//		//if err != nil{
	//		//	fmt.Println(err.Error())
	//		//}
	//		fmt.Println(shopAdmin)
	//
	//	})
	//	v1.GET("a", func(context *gin.Context) {
	//
	//	})
	//}
	v1 := router.Group("v1")
	{
		v1.POST("login", user.Login)
	}

}
