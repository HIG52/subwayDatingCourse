package handlers

import (
	"net/http"

	"github.com/HIG52/subwayDatingCourse/config"
	"github.com/HIG52/subwayDatingCourse/model"
	"github.com/gin-gonic/gin"
)

// GetMembers 함수는 '/members' 경로에 대한 HTTP GET 요청을 처리한다.
func GetMembers(c *gin.Context) {

	var members []model.Member // Member 타입의 슬라이스 선언
	db := config.DB
	// GORM을 사용하여 모든 멤버를 데이터베이스에서 조회
	if err := db.Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 오류 발생 시 반환
		return
	}

	c.JSON(http.StatusOK, members) // 정상 조회 시 멤버 데이터 반환
}
