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

	// 게시글 리스트 조회
	r.GET("/boards", handlers.GetBoards)
	// 특정 게시글 조회
	r.GET("/boards/:id", handlers.GetBoardByID)
	// 게시글 작성
	r.POST("/boards", handlers.CreateBoard)
	// 특정 게시글 수정
	r.PUT("/boards/:id", handlers.UpdateBoard)
	// 특정 게시글 삭제
	r.DELETE("/boards/:id", handlers.DeleteBoard)

	r.Run() // listen and serve on 0.0.0.0:8080
}
