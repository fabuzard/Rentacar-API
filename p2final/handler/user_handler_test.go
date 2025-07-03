package handler

import (
	"net/http"
	"net/http/httptest"
	"p2final/dto"
	"p2final/model"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockUserService implements the UserService interface for testing
type MockUserService struct{}

func (m *MockUserService) GetByID(userID uint) (model.User, error) {
	return model.User{
		ID:      1,
		Name:    "Test User",
		Email:   "test@example.com",
		Balance: 50000,
	}, nil
}

func (m *MockUserService) TopUp(userID uint, amount int) (model.User, error) {
	return model.User{
		ID:      1,
		Name:    "Test User",
		Email:   "test@example.com",
		Balance: 100000 + amount,
	}, nil
}

// ✅ Fix: Accept dto.RegisterRequest
func (m *MockUserService) CreateUser(user dto.RegisterRequest) (model.User, error) {
	return model.User{}, nil
}

func (m *MockUserService) GetByEmail(email string) (model.User, error) {
	return model.User{}, nil
}

func TestGetMe(t *testing.T) {
	e := echo.New()
	handler := &UserHandler{service: &MockUserService{}}

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.GetMe(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test User")
		assert.Contains(t, rec.Body.String(), "test@example.com")
	}
}

func TestTopUp(t *testing.T) {
	e := echo.New()
	handler := &UserHandler{service: &MockUserService{}}

	body := `{"amount": 50000}`
	req := httptest.NewRequest(http.MethodPost, "/topup", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.TopUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Balance added successfully")
		assert.Contains(t, rec.Body.String(), "Test User")
	}
}
