package handler

import (
	"net/http"
	"net/http/httptest"
	"p2final/model"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mock implementation of the RentalService
type MockRental struct{}

func (m *MockRental) GetUserRentalHistories(userID uint) ([]model.RentalHistory, error) {
	return []model.RentalHistory{
		{
			ID:       1,
			UserID:   1,
			CarID:    2,
			Cost:     100000,
			RentedAt: time.Now(),
		},
	}, nil
}

func TestGetUserRentalHistories(t *testing.T) {
	e := echo.New()
	handler := &RentalHandler{service: &MockRental{}}

	req := httptest.NewRequest(http.MethodGet, "/users/rentals", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.GetUserRentalHistories(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "100000")
	}
}
