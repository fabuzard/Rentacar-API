package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(s service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: s}
}

// @Summary      Book a car
// @Description  Allows a user to book a car by ID
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        request body dto.BookCarRequest true "Car ID to book"
// @Success      201  {object} dto.BookingResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router       /bookings [post]
// @Security     ApiKeyAuth
func (h *BookingHandler) BookCar(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token", err.Error())
	}

	type requestBody struct {
		CarID uint `json:"car_id" validate:"required"`
	}

	var req requestBody
	if err := c.Bind(&req); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "BIND_ERROR", "Invalid JSON", err.Error())
	}

	rental, err := h.bookingService.BookCar(userID, req.CarID)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "BOOKING_FAILED", err.Error(), nil)
	}

	response := dto.BookingResponse{
		RentalID:    rental.ID,
		CarID:       rental.CarID,
		CarName:     rental.Car.Name,
		Category:    rental.Car.Category,
		Cost:        rental.Cost,
		RentedAt:    rental.RentedAt,
		UserBalance: rental.User.Balance,
	}

	return helper.SendSuccess(c, http.StatusCreated, "Car booked successfully", response)
}

// GetAllUsers godoc
// @Summary      Return a rented car
// @Description  Allows a user to return a previously rented car by Rental ID
// @Tags         bookings
// @Produce      json
// @Param        request  body  dto.ReturnCarRequest  true  "Return payload"
// @Success      200   {object} dto.SuccessResponse
// @Failure      400  {object} dto.ErrorResponse
// @Failure      401  {object} dto.ErrorResponse
// @Router       /bookings/return [post]
// @Security     ApiKeyAuth
func (h *BookingHandler) ReturnCar(c echo.Context) error {
	var req dto.ReturnCarRequest

	if err := c.Bind(&req); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_JSON", "Invalid request body", err.Error())
	}

	if req.RentalID == 0 {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_ID", "Rental ID is required", nil)
	}

	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", err.Error())
	}

	err = h.bookingService.ReturnCar(userID, req.RentalID)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "RETURN_FAILED", "Failed to return car", err.Error())
	}

	return helper.SendSuccess(c, http.StatusOK, "Car returned successfully", nil)
}
