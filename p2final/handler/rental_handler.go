package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

type RentalHandler struct {
	service service.RentalService
}

func NewRentalHandler(s service.RentalService) *RentalHandler {
	return &RentalHandler{service: s}
}

func (h *RentalHandler) GetUserRentalHistories(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing JWT", err.Error())
	}

	histories, err := h.service.GetUserRentalHistories(userID)
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "FETCH_ERROR", "Failed to fetch rental history", err.Error())
	}

	var response []dto.RentalHistoryResponse
	for _, history := range histories {
		r := dto.RentalHistoryResponse{
			ID:       history.ID,
			CarName:  history.Car.Name,
			Cost:     history.Cost,
			RentedAt: history.RentedAt,
			ReturnAt: history.ReturnAt,
		}
		response = append(response, r)
	}

	return helper.SendSuccess(c, http.StatusOK, "Rental history fetched", response)
}
