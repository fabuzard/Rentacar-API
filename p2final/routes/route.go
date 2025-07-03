package routes

import (
	"net/http"
	"os"
	"p2final/handler"
	"p2final/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	carHandler *handler.CarHandler,
	bookingHandler *handler.BookingHandler,
	rentalHandler *handler.RentalHandler,
	transactionHandler *handler.TransactionHandler,
) {
	jwtMiddleware := middleware.JWTMiddleware(os.Getenv("JWT_SECRET"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "🚀 Server running and DB connected!")
	})

	// Auth routes
	e.POST("/auth/register", authHandler.Register)
	e.POST("/auth/login", authHandler.Login)

	// User routes
	e.GET("/users/me", userHandler.GetMe, jwtMiddleware)
	e.POST("/users/topup", userHandler.TopUp, jwtMiddleware)

	// Car routes
	e.POST("/cars", carHandler.CreateCar, jwtMiddleware)
	e.GET("/cars/available", carHandler.GetAvailableCars, jwtMiddleware)
	e.GET("/cars/mine", carHandler.GetMyCars, jwtMiddleware)
	e.DELETE("/cars/:id", carHandler.DeleteCar, jwtMiddleware)

	// Booking routes
	e.POST("/bookings", bookingHandler.BookCar, jwtMiddleware)
	e.POST("/bookings/return", bookingHandler.ReturnCar, jwtMiddleware)

	// Rental history
	e.GET("/users/rentalhistory", rentalHandler.GetUserRentalHistories, jwtMiddleware)

	// Transactions
	e.GET("/users/transactionhistory", transactionHandler.GetMyTransactions, jwtMiddleware)
}
