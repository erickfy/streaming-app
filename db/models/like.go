package models

type Like struct {
	ID        uint  `gorm:"primaryKey"`
	IsLiked   bool  `gorm:"not null"`
	UserID    uint  `gorm:"not null"`
	User      User  `gorm:"foreignKey:UserID"`
	VideoID   uint  `gorm:"not null"`
	Video     Video `gorm:"foreignKey:VideoID"`
	CreatedAt int64
}
