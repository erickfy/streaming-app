package models

type User struct {
	ID            uint           `gorm:"primaryKey"`
	Username      string         `gorm:"unique;not null"`
	Email         string         `gorm:"unique;not null"`
	Password      string         `gorm:"not null"`
	ProfilePicURL string         `gorm:"size:255"`
	Subscribers   []Subscription `gorm:"foreignKey:SubscribedToID"`
	Videos    []Video        	`gorm:"foreignKey:UploadedBy"`

	CreatedAt     int64
	UpdatedAt     int64
}
