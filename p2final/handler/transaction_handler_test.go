package handler

import (
	"net/http"
	"net/http/httptest"
	"p2final/model"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockTransaction struct{}

func (m *MockTransaction) GetUserTransactions(userID uint) ([]model.TransactionHistory, error) {
	return []model.TransactionHistory{
		{
			SenderID:    1,
			ReceiverID:  2,
			Amount:      100000,
			Description: "Test transaction",
		},
	}, nil
}

func TestGetUserTransactions(t *testing.T) {
	e := echo.New()
	handler := &TransactionHandler{service: &MockTransaction{}}

	req := httptest.NewRequest(http.MethodGet, "/transactionhistories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
		Valid: true,
	})

	if assert.NoError(t, handler.GetMyTransactions(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test transaction")
	}
}
