package models

type Video struct {
	ID           uint       `gorm:"primaryKey"`
	Title        string     `gorm:"not null"`
	Description  string     `gorm:"size:500"`
	URL          string     `gorm:"not null"`
	ThumbnailURL string     `gorm:"size:255"`
	Views        uint       `gorm:"default:0"`
	UserID       uint       `gorm:"not null"`
	User         User       `gorm:"foreignKey:UserID"`
	Likes        []Like     `gorm:"foreignKey:VideoID"`
	Comments     []Comment  `gorm:"foreignKey:VideoID"`
	Playlists    []Playlist `gorm:"many2many:playlist_videos;"`
	CreatedAt    int64
	UpdatedAt    int64
}
