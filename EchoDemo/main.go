package main

import (
	"GoSql/EchoDemo/controllers"
	"GoSql/EchoDemo/utils"
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Router 路由
var Router *echo.Echo

func main() {

	//中间件
	Router.Use(middleware.Gzip())
	Router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
	}))
	//跨域配置

	if corsConfig, err := utils.GetCORSConfig(); err == nil {
		Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: corsConfig,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	} else {
		fmt.Println("corsconfig error", err)
	}

	Router.POST("/login", controllers.Login)
	Router.POST("/register", controllers.Login)
	Router.GET("/welcome", controllers.Welcome)

	//user
	userRouter := Router.Group("/user", utils.JWTConfig())
	{
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
