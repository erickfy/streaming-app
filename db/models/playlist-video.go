package models

import "time"

type PlaylistVideo struct {
	ID         uint      `gorm:"primaryKey"`
	PlaylistID uint      `gorm:"not null"`
	VideoID    uint      `gorm:"not null"`
	AddedAt    time.Time `gorm:"autoCreateTime"`
	Playlist   Playlist  `gorm:"foreignKey:PlaylistID"`
	Video      Video     `gorm:"foreignKey:VideoID"`
}
