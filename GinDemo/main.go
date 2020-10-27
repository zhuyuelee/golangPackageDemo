package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("starting ...")
	router := gin.Default()
	admin := router.Group("/api")
	{
		admin.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"temp":  "index",
				"title": "登录",
			})
		})
		admin.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"temp":  "login",
				"title": "登录",
			})
		})
		admin.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"temp":  "user",
				"title": "首页",
			})
		})
	}
	router.Run(":8090")
}
