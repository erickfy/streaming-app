package main

import (
	"fmt"
	"log"
	"os"
	"streaming/controllers"
	"streaming/middleware"
	"streaming/services"

	// "streaming/db/models"
	// "streaming/db/models/database"

	"streaming/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// CONFS
	utils.InitConf()

	// Leer las variables de entorno
	dbURL := os.Getenv("DATABASE_URL")
	DB_SCHEMA := os.Getenv("DB_SCHEMA")
	port := os.Getenv("PORT")

	fmt.Println("URL de la base de datos:", dbURL)
	fmt.Println("Entorno:", DB_SCHEMA)

	// Configuración de la base de datos
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

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

	// ALL ROUTE CONTROLLERS
	// Registrar la ruta con el controlador

	// Crear el servicio de video
	videoService := services.NewVideoService() // Crea una instancia de tu servicio (implementa esta función)

	// Crear el controlador de video
	videoController := &controllers.VideoController{
		VideoService: videoService,
	}

	// Inicializa el servicio de usuario
	userService := services.NewUserService(db)

	// Inicializa el controlador de registro
	registerController := controllers.NewRegisterController(userService)

	app.Get("/stream/:id", videoController.StreamVideo)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/payment", controllers.ProcessPayment)
	app.Post("/login", controllers.Login)

	app.Post("/register", registerController.RegisterUser)

	// Ruta protegida por JWT
	app.Get("/profile", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{
			"status": "success",
			"user":   user,
		})
	})

	// Start the Fiber server
	log.Printf("Server is running on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Unable to start the application: %v", err)
	}
}
