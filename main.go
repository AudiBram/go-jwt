package main

import (
	"github.com/gin-gonic/gin"
	"jwt-gin/controllers"
	"jwt-gin/initializer"
	"jwt-gin/middlewares"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
	initializer.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()

}
