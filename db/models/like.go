package models

type Like struct {
	ID        uint  `gorm:"primaryKey"`
	UserID    uint  `gorm:"not null"`
	User      User  `gorm:"foreignKey:UserID"`
	VideoID   uint  `gorm:"not null"`
	Video     Video `gorm:"foreignKey:VideoID"`
	CreatedAt int64
}
