package main

import (
	"fmt"
	"log"
	"os"
	"streaming/controllers"

	// "streaming/db/models"
	// "streaming/db/models/database"

	"streaming/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// CONFS
	utils.InitConf()


	// // database conf
	// 	database.ConnectDB()

	// // Migración de modelos
	// err := database.DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Video{},
	// 	&models.Subscription{},
	// 	&models.Comment{},
	// 	&models.Playlist{},
	// 	&models.Payment{},
	// 	&models.Like{},
	// )
	// if err != nil {
	// 	log.Fatal("Failed to migrate models:", err)
	// }

	// log.Println("Database migration completed")

// Leer una variable de entorno
	// Leer las variables de entorno
	dbURL := os.Getenv("DATABASE_URL")
	DB_SCHEMA := os.Getenv("DB_SCHEMA")

	// Usar las variables de entorno
	fmt.Println("URL de la base de datos:", dbURL)
	fmt.Println("Entorno:", DB_SCHEMA)


	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
app.Use(logger.New(logger.Config{
    Format: "[${time}] ${ip} -> ${method} ${path}\n",
}))
	// Define a route for the home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// Define a route for streaming video
	app.Get("/stream", controllers.StreamVideo)

	// Start the Fiber server
    port := "3000"
    log.Printf("Server is running on http://localhost:%s", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("Unable to start the application: %v", err)
    }
}