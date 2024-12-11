package models

type Playlist struct {
	ID          uint            `gorm:"primaryKey"`
	Title       string          `gorm:"not null"`
	Description string          `gorm:"size:255"`
	UserID      uint            `gorm:"not null"`
	User        User            `gorm:"foreignKey:UserID"`
	Videos      []PlaylistVideo `gorm:"foreignKey:PlaylistID"`
	// Videos      []Video `gorm:"many2many:playlist_videos;"`
	CreatedAt int64
}
