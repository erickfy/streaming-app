package models

import "time"

type WatchHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	VideoID   uint      `gorm:"not null"`
	WatchedAt time.Time `gorm:"autoCreateTime"`
	User      User      `gorm:"foreignKey:UserID"`
	Video     Video     `gorm:"foreignKey:VideoID"`
}
