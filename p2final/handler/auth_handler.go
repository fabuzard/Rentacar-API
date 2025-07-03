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

// @Summary      Register a new user
// @Description  Creates a new user account with name, email, and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequest true "User registration data"
// @Success      201  {object} dto.UserResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router       /auth/register [post]
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

// @Summary      User login
// @Description  Authenticates user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "User login credentials"
// @Success      200  {object} map[string]string
// @Failure      400  {object} dto.ErrorResponse
// @Failure      401  {object} dto.ErrorResponse
// @Failure      500  {object} dto.ErrorResponse
// @Router       /auth/login [post]
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
