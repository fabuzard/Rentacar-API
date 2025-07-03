package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"p2final/model"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockCarService implements the CarService interface for testing
type MockCarService struct{}

func (m *MockCarService) GetAllAvailable() ([]model.Car, error) {
	return []model.Car{
		{Name: "Avanza", Code: "AV01", Category: "Family", RentalCost: 300000, IsAvailable: true},
		{Name: "Brio", Code: "BR01", Category: "City", RentalCost: 250000, IsAvailable: true},
	}, nil
}

func (m *MockCarService) CreateCar(car model.Car) (model.Car, error) {
	return model.Car{
		Name:        "Civic",
		Code:        "CV01",
		Category:    "Sedan",
		RentalCost:  350000,
		IsAvailable: true,
	}, nil
}

func (m *MockCarService) GetOwnedCars(userID uint) ([]model.Car, error) {
	return []model.Car{
		{Name: "Jazz", Code: "JZ01", Category: "Hatchback", RentalCost: 200000, IsAvailable: true},
	}, nil
}
func (m *MockCarService) GetByID(id uint) (model.Car, error) {
	// Return a car with UserID set to match the token (1)
	return model.Car{
		ID:      id,
		OwnerID: 1,
		Name:    "Avanza",
	}, nil
}

func (m *MockCarService) DeleteCar(id uint) error                    { return nil }
func (m *MockCarService) UpdateCar(car model.Car) (model.Car, error) { return model.Car{}, nil }

func TestGetAvailableCars(t *testing.T) {
	e := echo.New()

	// Setup mock service and handler

	handler := &CarHandler{service: &MockCarService{}}

	req := httptest.NewRequest(http.MethodGet, "/cars/available", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAvailableCars(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Avanza")
	}
}

func TestCreateCar(t *testing.T) {
	e := echo.New()
	handler := &CarHandler{service: &MockCarService{}}

	carJSON := `{
		"name": "Civic",
		"code": "CV01",
		"category": "Sedan",
		"rental_cost": 350000,
		"is_available": true
	}`

	req := httptest.NewRequest(http.MethodPost, "/cars", strings.NewReader(carJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.CreateCar(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Civic")
	}
}

func TestGetMyCars(t *testing.T) {
	e := echo.New()
	handler := &CarHandler{service: &MockCarService{}}

	req := httptest.NewRequest(http.MethodGet, "/cars/mine", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Inject JWT token with user_id = 1
	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.GetMyCars(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Jazz")
	}
}

func TestDeleteCar(t *testing.T) {
	e := echo.New()
	handler := &CarHandler{service: &MockCarService{}}

	req := httptest.NewRequest(http.MethodDelete, "/cars/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set path parameter :id to "1"
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Inject JWT token with user_id = 1
	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.DeleteCar(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success") // optional: change this based on your actual message
	}
}
