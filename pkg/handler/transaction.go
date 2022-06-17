package handler

import (
	"net/http"

	"github.com/callmehorhe/konstanta/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var input models.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid input")
		return
	}

	if err := h.services.CreateTransaction(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction created",
	})
}

type ID struct {
	Id int `json:"id"`
}

func (h *Handler) GetAllTransactionsByUserID(c *gin.Context) {
	var id ID
	if err := c.BindJSON(&id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	transactions, err := h.services.GetAllTransactionsByID(id.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user id":      id,
		"transactions": transactions,
	})
}

type Email struct {
	Email string `json:"email"`
}

func (h *Handler) GetAllTransactionsByEmail(c *gin.Context) {
	var email Email
	if err := c.BindJSON(&email); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	transactions, err := h.services.GetAllTransactionsByEmail(email.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"email":        email,
		"transactions": transactions,
	})
}
