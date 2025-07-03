package handler

import (
	"net/http"
	"net/http/httptest"
	"p2final/model"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// ---- Mock Service ----
type MockBookingService struct{}

func (m *MockBookingService) BookCar(renterID, carID uint) (model.RentalHistory, error) {
	return model.RentalHistory{
		ID:    1,
		CarID: carID,
		Car: model.Car{
			Name:     "Civic",
			Category: "Sedan",
		},
		User: model.User{
			ID:      renterID,
			Balance: 500000,
		},
		Cost:     350000,
		RentedAt: time.Now(),
	}, nil
}

func (m *MockBookingService) ReturnCar(userID, rentalID uint) error {
	return nil
}

// ---- Test: BookCar ----
func TestBookCar(t *testing.T) {
	e := echo.New()
	handler := &BookingHandler{bookingService: &MockBookingService{}}

	body := `{"car_id": 2}`
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.BookCar(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Civic")
		assert.Contains(t, rec.Body.String(), "Sedan")
	}
}

// ---- Test: ReturnCar ----
func TestReturnCar(t *testing.T) {
	e := echo.New()
	handler := &BookingHandler{bookingService: &MockBookingService{}}

	body := `{"rental_id": 1}`
	req := httptest.NewRequest(http.MethodPost, "/returncar", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.ReturnCar(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "returned")
	}
}
