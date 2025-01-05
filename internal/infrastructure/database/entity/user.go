package entity

import "time"

// ユーザーエンティティのDBモデル
type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique"`
	Birthday time.Time
}
