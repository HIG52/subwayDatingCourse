package model

import "gorm.io/gorm"

// 회원 정보를 나타내는 Member구조체
// gorm.Model을 포함하여 GORM의 기본 모델 필드(ID, CreatedAt 등)도 사용
type Member struct {
	gorm.Model        // GORM의 기본 모델을 임베딩
	No         int    `gorm:"primarykey"` // 고유 번호, 기본 키로 사용된다
	Id         string // 회원의 ID
	Password   string // 회원의 비밀번호
	Name       string // 회원의 이름
}

// TableName 함수는 GORM에 이 구조체가 사용할 데이터베이스 테이블의 이름을 명시한다
// 'member' 테이블을 사용합니다.
func (Member) TableName() string {
	return "member"
}
