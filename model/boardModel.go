package model

import (
	"time"

	"gorm.io/gorm"
)

// 게시판 정보를 나타내는 Board구조체
// gorm.Model을 포함하여 GORM의 기본 모델 필드(ID, CreatedAt 등)도 사용
type Board struct {
	gorm.Model           // GORM의 기본 모델을 임베딩
	No         int       `gorm:"primarykey"` // 고유 번호, 기본 키로 사용된다
	Subject    string    // 게시글 제목
	Contents   string    // 게시글 내용
	Author     string    // 작성자
	CreatedAt  time.Time // 글 작성 시간
	UpdatedAt  time.Time // 글 수정 시간
	DeletedAt  time.Time // 글 삭제 시간
}

// TableName 함수는 GORM에 이 구조체가 사용할 데이터베이스 테이블의 이름을 명시한다
// 'board' 테이블을 사용합니다.
func (Board) TableName() string {
	return "board"
}
