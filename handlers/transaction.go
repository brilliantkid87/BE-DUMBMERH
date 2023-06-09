package handlers

import (
	dto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *transactionHandler {
	return &transactionHandler{TransactionRepository}
}

func (h *transactionHandler) FIndTransaction(c echo.Context) error {
	transactions, err := h.TransactionRepository.FIndTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *transactionHandler) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	trip, _ := h.TransactionRepository.GetUpdateTripByID(request.TripId)
	user, _ := h.TransactionRepository.GetUserByID(request.UserID)
	// country, _ := h.TransactionRepository.GetCountriesByID(request)

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripId:     request.TripId,
		UserID:     request.UserID,
		Trip:       models.TripsResponse(trip),
		User:       models.UserResponse(user),
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *transactionHandler) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(user)})
}

func (h *transactionHandler) UpdateTransaction(c echo.Context) error {
	request := new(transactiondto.UpdateTransactionRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	trip, _ := h.TransactionRepository.GetUpdateTripByID(request.TripId)

	if request.CounterQty != 0 {
		transaction.CounterQty = request.CounterQty
	}

	transaction.Trip = trip

	if request.Total != 0 {
		transaction.Total = request.Total
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	if request.Attachment != "" {
		transaction.Attachment = request.Attachment
	}

	if request.TripId != 0 {
		transaction.TripId = request.TripId
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func (h *transactionHandler) GetTransactionByUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransactionByUser(id)

	// fmt.Println(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID:         u.ID,
		CounterQty: u.CounterQty,
		Total:      u.Total,
		Status:     u.Status,
		Attachment: u.Attachment,
		TripId:     u.TripId,
	}
}
