package models

type Payment struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	Amount    float64 `gorm:"not null"`
	Status    string  `gorm:"default:'Pending'"` // "Pending", "Completed", etc.
	CreatedAt int64
}
