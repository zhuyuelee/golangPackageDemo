package main

import (
	"GoSql/EchoDemo/controllers"
	"GoSql/EchoDemo/utils"
	"net/http"

	"github.com/labstack/echo"
)

//Router 路由
var Router *echo.Echo

func main() {

	Router.GET("/", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.String(http.StatusOK, "hello echo demo")
	})
	Router.POST("/login", controllers.Login)
	Router.GET("/welcome", controllers.Welcome)

	userRouter := Router.Group("/user")
	{
		//JWTConfig
		userRouter.Use(utils.JWTConfig())
		userRouter.GET("/welcome", controllers.Welcome)
		//获取会员资料
		userRouter.GET("/:id", controllers.GetUser)
		userRouter.GET("/list", controllers.GetUserList)
		userRouter.POST("/add", controllers.AddUser)
		userRouter.PUT("/update", controllers.UpdateUser)
		userRouter.DELETE("/delete/:id", controllers.DeleteUser)
	}
	Router.Start(":9090")
}

func init() {
	Router = echo.New()
}
