package handler

import (
	"net/http"
	"p2final/dto"
	"p2final/helper"
	"p2final/service"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// @Summary      Get transaction history
// @Description  Retrieves the authenticated user's transaction history
// @Tags         transactions
// @Produce      json
// @Success      200  {array}  dto.TransactionResponse
// @Failure      401  {object} dto.ErrorResponse
// @Failure      500  {object} dto.ErrorResponse
// @Router       /users/transactionhistory [get]
// @Security     ApiKeyAuth
func (h *TransactionHandler) GetMyTransactions(c echo.Context) error {
	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or missing token", err.Error())
	}

	txs, err := h.service.GetUserTransactions(userID)
	if err != nil {
		return helper.SendError(c, http.StatusInternalServerError, "DB_ERROR", "Failed to fetch transactions", err.Error())
	}
	var response []dto.TransactionResponse
	for _, tx := range txs {
		response = append(response, dto.TransactionResponse{
			ID: tx.ID,
			Sender: dto.TransactionUserInfo{
				ID:    tx.Sender.ID,
				Name:  tx.Sender.Name,
				Email: tx.Sender.Email,
			},
			Receiver: dto.TransactionUserInfo{
				ID:    tx.Receiver.ID,
				Name:  tx.Receiver.Name,
				Email: tx.Receiver.Email,
			},
			Amount:      tx.Amount,
			Description: tx.Description,
			RentalID:    tx.RentalID,
			CreatedAt:   tx.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return helper.SendSuccess(c, http.StatusOK, "Transaction history retrieved", response)
}
