package models

type DocumentContribute struct {
	UserID     string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DocumentID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UpdateTime int64  `gorm:"type:bigint"`
}
