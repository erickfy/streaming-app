package confs

import (
	"streaming/db/models"
	"streaming/db/models/database"
)

func LoadConfDatabase() error {
	// database conf
	database.ConnectDB()

	// Migración de modelos
	err := database.DB.AutoMigrate(
		// &models.User{},
		// &models.Video{},
		// &models.Subscription{},
		// &models.Comment{},
		// &models.Playlist{},
		// &models.Payment{},
		// &models.Like{},
		&models.User{},
		&models.Video{},
		&models.Subscription{},
		&models.Payment{},
		&models.WatchHistory{},
		&models.Comment{},
		&models.Like{},
		&models.Playlist{},
		&models.PlaylistVideo{},
	)

	if err != nil {
		return err
	}

	// log.Println("Database migration completed")
	return nil
}
