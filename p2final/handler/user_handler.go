package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// @Summary      Top up user balance
// @Description  Add balance to the authenticated user's account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.TopUpRequest true "Top up amount"
// @Success      200  {object} dto.TopupResponse
// @Failure      400  {object} dto.ErrorResponse
// @Failure      401  {object} dto.ErrorResponse
// @Router       /users/topup [post]
// @Security     ApiKeyAuth
func (h *UserHandler) TopUp(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing token", err.Error())
	}

	var req dto.TopUpRequest
	if err := c.Bind(&req); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_JSON", "Failed to parse JSON body", err.Error())
	}

	updatedUser, err := h.service.TopUp(userID, req.Amount)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "TOPUP_FAILED", "Failed to top up balance", err.Error())
	}

	response := dto.TopupResponse{
		ID:      updatedUser.ID,
		Name:    updatedUser.Name,
		Balance: updatedUser.Balance,
	}

	return helper.SendSuccess(c, http.StatusOK, "Balance added successfully", response)
}

// @Summary      Get user profile
// @Description  Retrieve the authenticated user's profile and balance
// @Tags         users
// @Produce      json
// @Success      200  {object} dto.MeResponse
// @Failure      401  {object} dto.ErrorResponse
// @Failure      500  {object} dto.ErrorResponse
// @Router       /users/me [get]
// @Security     ApiKeyAuth
func (h *UserHandler) GetMe(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing token", err.Error())
	}

	user, err := h.service.GetByID(userID)
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "DB_ERROR", "Failed to fetch user", err.Error())
	}

	response := dto.MeResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}

	return helper.SendSuccess(c, http.StatusOK, "User profile retrieved", response)
}
