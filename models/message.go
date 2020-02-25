package models

import (
	"archie/connection/postgres_conn"
	"github.com/jinzhu/gorm"
)

type Message struct {
	ID          string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Owner       string `gorm:"type:varchar(100)"json:"-"`
	Type        int    `gorm:"type:int"json:"-"`
	From        string `gorm:"type:varchar(100)"json:"from"`
	Users       []User `gorm:"many2many:user_messages;"json:"-"`
	SendTime    int64  `gorm:"type:bigint"json:"sendTime"`
	MessageType int    `gorm:"type:int"json:"messageType"`
	IsRead      bool   `gorm:"bool;default:FALSE"json:"isRead"`
	IsDelete    bool   `gorm:"bool;default:FALSE"json:"isDelete"`
	Main        string `gorm:"type:varchar(5000)"json:"main"` // contain body and title (marshal channel_message)
}

func (message *Message) Create(to []string) error {
	return postgres_conn.Transaction(func(db *gorm.DB) error {
		query := db.Model(User{})
		var users []User

		for _, id := range to {
			query = query.Or("id = ?", id)
		}

		if err := query.Find(&users).Error; err != nil {
			return err
		}

		message.Users = users
		return db.Create(message).Error
	})
}
