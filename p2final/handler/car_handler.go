package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/model"
	"p2final/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(s service.CarService) *CarHandler {
	return &CarHandler{service: s}
}

func (h *CarHandler) CreateCar(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing token", err.Error())
	}

	var req dto.CreateCarRequest

	if err := c.Bind(&req); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_JSON", "Invalid request format", err.Error())

	}

	car := model.Car{
		Name:       req.Name,
		Code:       req.Code,
		Category:   req.Category,
		RentalCost: req.RentalCost,
		OwnerID:    userID,
	}

	createdCar, err := h.service.CreateCar(car)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "Car registration failed", err.Error(), nil)
	}

	response := dto.CreateCarResponse{
		ID:       createdCar.ID,
		Name:     createdCar.Name,
		Code:     createdCar.Code,
		Category: createdCar.Category,
	}

	return helper.SendSuccess(c, http.StatusCreated, "Car registered successfully", response)
}

func (h *CarHandler) GetAvailableCars(c echo.Context) error {
	_, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing token", err.Error())
	}

	cars, err := h.service.GetAllAvailable()
	if err != nil {
		return helper.SendError(
			c,
			http.StatusInternalServerError,
			"FETCH_ERROR",
			"Failed to fetch available cars",
			err.Error(),
		)
	}

	// Map model.Car slice to []dto.GetCarResponse
	var response []dto.GetCarResponse
	for _, car := range cars {
		response = append(response, dto.GetCarResponse{
			ID:          car.ID,
			Name:        car.Name,
			Code:        car.Code,
			Category:    car.Category,
			RentalCost:  car.RentalCost,
			IsAvailable: car.IsAvailable,
		})
	}

	return helper.SendSuccess(c, http.StatusOK, "Available cars", response)
}

func (h *CarHandler) GetMyCars(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", err.Error())
	}

	cars, err := h.service.GetOwnedCars(userID)
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "FETCH_ERROR", "Failed to fetch your cars", err.Error())
	}

	var response []dto.GetCarResponse
	for _, car := range cars {
		response = append(response, dto.GetCarResponse{
			ID:          car.ID,
			Name:        car.Name,
			Code:        car.Code,
			Category:    car.Category,
			RentalCost:  car.RentalCost,
			IsAvailable: car.IsAvailable,
		})
	}
	return helper.SendSuccess(c, http.StatusOK, "Your cars", response)
}
func (h *CarHandler) DeleteCar(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", err.Error())
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid car ID", err.Error())
	}

	// Fetch the car to verify ownership
	car, err := h.service.GetByID(uint(id))
	if err != nil {
		return helper.SendError(c, http.StatusNotFound, "NOT_FOUND", "Car not found", err.Error())
	}

	if car.OwnerID != userID {
		return helper.SendError(c, http.StatusForbidden, "FORBIDDEN", "You are not the owner of this car", nil)
	}

	// Proceed with deletion
	err = h.service.DeleteCar(uint(id))
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "DELETE_FAILED", "Failed to delete car", err.Error())
	}
	response := dto.DeletedCarResponse{
		ID:   car.ID,
		Name: car.Name,
		Code: car.Code,
	}

	return helper.SendSuccess(c, http.StatusOK, "Car deleted successfully", response)
}
