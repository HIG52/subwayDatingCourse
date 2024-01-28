package handlers

import (
	"net/http"
	"strconv"

	"github.com/HIG52/subwayDatingCourse/config"
	"github.com/HIG52/subwayDatingCourse/model"
	"github.com/gin-gonic/gin"
)

// GetBoards 함수는 '/boards' 경로에 대한 HTTP GET 요청을 처리한다.
func GetBoards(c *gin.Context) {
	var boards []model.Board // Board 타입의 슬라이스 선언
	db := config.DB

	// 페이징을 위한 파라미터를 쿼리에서 가져오기
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))          // 기본값은 1
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")) // 기본값은 10

	offset := (page - 1) * pageSize // 페이지 시작 위치 계산

	// GORM을 사용하여 페이징 처리된 게시글을 데이터베이스에서 조회
	if err := db.Offset(offset).Limit(pageSize).Find(&boards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 오류 발생 시 반환
		return
	}

	c.JSON(http.StatusOK, boards) // 정상 조회 시 게시글 데이터 반환
}

// GetBoardByID 함수는 특정 게시글을 조회한다.
func GetBoardByID(c *gin.Context) {
	var board model.Board
	db := config.DB

	// URL에서 id를 가져온다.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	// ID를 사용하여 데이터베이스에서 게시글을 조회한다.
	result := db.First(&board, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	c.JSON(http.StatusOK, board)
}

// CreateBoard 함수는 새 게시글을 추가한다.
func CreateBoard(c *gin.Context) {
	var board model.Board
	db := config.DB

	// 클라이언트의 요청 본문에서 게시글 데이터를 파싱한다.
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 데이터베이스에 게시글을 추가한다.
	if err := db.Create(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, board)
}

// UpdateBoard 함수는 특정 게시글을 수정한다.
func UpdateBoard(c *gin.Context) {
	var board model.Board
	db := config.DB
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	// 데이터베이스에서 해당 ID의 게시글을 찾는다.
	if err := db.First(&board, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	// 클라이언트의 요청 본문에서 수정된 데이터를 파싱한다.
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 데이터베이스에서 게시글을 수정한다.
	if err := db.Save(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, board)
}

// DeleteBoard 함수는 특정 게시글을 삭제한다.
func DeleteBoard(c *gin.Context) {
	var board model.Board
	db := config.DB
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	// 데이터베이스에서 해당 ID의 게시글을 삭제한다.
	if err := db.Delete(&board, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Board deleted"})
}
