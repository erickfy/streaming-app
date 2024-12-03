package models

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	VideoID   uint   `gorm:"not null"`
	Video     Video  `gorm:"foreignKey:VideoID"`
	CreatedAt int64
}
