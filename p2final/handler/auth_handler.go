package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var u dto.RegisterRequest
	if err := c.Bind(&u); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_JSON", "Invalid request format", err.Error())
	}

	user, err := h.Service.CreateUser(u)
	if err != nil {
		return helper.SendError(c, http.StatusBadRequest, "REGISTER_FAILED", err.Error(), nil)
	}

	response := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return helper.SendSuccess(c, http.StatusCreated, "User registered successfully", response)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var u dto.LoginRequest
	if err := c.Bind(&u); err != nil {
		return helper.SendError(c, http.StatusBadRequest, "INVALID_JSON", "Invalid request format", err.Error())
	}

	user, err := h.Service.GetByEmail(u.Email)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "EMAIL_NOT_FOUND", "Email not found", nil)
	}

	if !helper.CheckPasswordHash(user.Password, u.Password) {
		return helper.SendError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Wrong password", nil)
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "TOKEN_ERROR", "Failed generating token", err.Error())
	}

	return helper.SendSuccess(c, http.StatusOK, "Login successful", echo.Map{
		"token": token,
	})
}
