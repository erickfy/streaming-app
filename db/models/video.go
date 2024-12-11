package models

type Video struct {
	ID            uint    `gorm:"primaryKey"`
	Title         string  `gorm:"not null"`
	Description   string  `gorm:"size:500"`
	VideoPath     string  `gorm:"type:varchar(512);not null"`
	ThumbnailPath string  `gorm:"type:varchar(512);not null"`
	UploadedBy    uint    `gorm:"not null"`
	IsFree        bool    `gorm:"default:true"`
	Price         float64 `gorm:"type:decimal(10,2);default:0.00"`

	User           User           `gorm:"foreignKey:UploadedBy"`
	Comments       []Comment      `gorm:"foreignKey:VideoID"`
	WatchHistories []WatchHistory `gorm:"foreignKey:VideoID"`
	Likes          []Like         `gorm:"foreignKey:VideoID"`

	Views uint `gorm:"default:0"`

	Playlists []Playlist `gorm:"many2many:playlist_videos;"`
	CreatedAt int64
	UpdatedAt int64
}
