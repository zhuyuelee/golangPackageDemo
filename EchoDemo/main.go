package main

import (
	"GoSql/EchoDemo/controllers"
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
	//获取会员资料
	Router.GET("/user/:id", controllers.GetUser)
	Router.GET("/user/list", controllers.GetUserList)
	Router.POST("/user/add", controllers.AddUser)
	Router.PUT("/user/update", controllers.UpdateUser)
	Router.DELETE("/user/delete/:id", controllers.DeleteUser)
	Router.Start(":9090")
}

func init() {
	Router = echo.New()
}
