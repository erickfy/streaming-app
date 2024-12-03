package models

type Subscription struct {
	ID             uint `gorm:"primaryKey"`
	SubscriberID   uint
	Subscriber     User `gorm:"foreignKey:SubscriberID"`
	SubscribedToID uint
	SubscribedTo   User `gorm:"foreignKey:SubscribedToID"`
	CreatedAt      int64
}
