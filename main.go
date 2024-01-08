package main

import (
	"net/http"

	"github.com/HIG52/subwayDatingCourse/config"
	"github.com/HIG52/subwayDatingCourse/model"
	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DbConnection()

	r := gin.Default()

	r.GET("/")

	r.GET("/members", func(c *gin.Context) {
		var members []model.Member
		result := db.Find(&members)
		if result.Error != nil { // 에러 체크
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, members)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
