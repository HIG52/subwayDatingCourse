package main

import (
	"github.com/HIG52/subwayDatingCourse/config"
	"github.com/HIG52/subwayDatingCourse/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	config.DbConnection()

	r := gin.Default()

	r.GET("/")
	r.GET("/members", handlers.GetMembers)

	r.Run() // listen and serve on 0.0.0.0:8080
}
