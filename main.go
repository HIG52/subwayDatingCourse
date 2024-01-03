package main

import (
	"github.com/HIG52/subwayDatingCourse/config"
	"github.com/gin-gonic/gin"
)

type Member struct {
	No       int `gorm:"primarykey"`
	Id       string
	Password string
	Name     string
}

func main() {

	config.DbConnection()

	router := gin.Default()

	router.GET("/")

	router.Run("localhost:8080")
}
