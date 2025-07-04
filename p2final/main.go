// @title        Car Rental API
// @version      1.0
// @description  Final Project Phase 2 - Car Rental System API
// @termsOfService http://swagger.io/terms/

// @contact.name   Fahreza Abuzard Alghifary
// @contact.email  fabuzard123@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

package main

import (
	"os"
	"p2final/config"
	"p2final/handler"
	"p2final/model"
	"p2final/repository"
	"p2final/routes"
	"p2final/service"

	_ "p2final/docs" // This line is necessary for swag to find docs

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.LoadEnv()
	db := config.DBInit()

	e := echo.New()
	config.DB.AutoMigrate(&model.User{}, &model.Car{}, &model.RentalHistory{}, &model.TransactionHistory{})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	carRepo := repository.NewCarRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	rentalRepo := repository.NewRentalRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	carService := service.NewCarService(carRepo)
	bookingService := service.NewBookingService(bookingRepo)
	rentalService := service.NewRentalService(rentalRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandler(userService)
	carHandler := handler.NewCarHandler(carService)
	bookingHandler := handler.NewBookingHandler(bookingService)
	rentalHandler := handler.NewRentalHandler(rentalService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Routes
	routes.SetupRoutes(e, authHandler, userHandler, carHandler, bookingHandler, rentalHandler, transactionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
