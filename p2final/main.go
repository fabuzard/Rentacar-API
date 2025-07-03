package main

import (
	"fmt"
	"net/http"
	"os"
	"p2final/config"
	"p2final/handler"
	"p2final/middleware"
	"p2final/model"
	"p2final/repository"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	db := config.DBInit()

	e := echo.New()
	config.DB.AutoMigrate(&model.User{}, &model.Car{}, &model.RentalHistory{}, &model.TransactionHistory{})

	// Repositories
	userRepo := repository.NewUserRepository(db)
	carRepo := repository.NewCarRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	carService := service.NewCarService(carRepo)
	bookingService := service.NewBookingService(bookingRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandler(userService)
	carHandler := handler.NewCarHandler(carService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	rentalRepo := repository.NewRentalRepository(db)
	rentalService := service.NewRentalService(rentalRepo)
	rentalHandler := handler.NewRentalHandler(rentalService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "🚀 Server running and DB connected!")
	})

	// Auth
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// User
	e.GET("/me", userHandler.GetMe, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	e.POST("/topup", userHandler.TopUp, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	// Cars
	e.POST("/cars", carHandler.CreateCar, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	e.GET("/cars", carHandler.GetAvailableCars, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	e.GET("/cars/mine", carHandler.GetMyCars, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	e.DELETE("/cars/:id", carHandler.DeleteCar, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	// Booking
	e.POST("/bookings", bookingHandler.BookCar, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	e.POST("/returncar", bookingHandler.ReturnCar, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	// e.POST("/bookings/:id/pay", bookingHandler.PayBooking, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	// rental
	e.GET("/rentalhistories", rentalHandler.GetUserRentalHistories, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	// Transaction
	e.GET("/transactionhistories", transactionHandler.GetMyTransactions, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	fmt.Println("Connected to DB")

	e.Logger.Fatal(e.Start(":8080"))
}
