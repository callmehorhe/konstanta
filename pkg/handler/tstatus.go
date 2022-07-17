package handler

import (
	"net/http"
	"strconv"

	"github.com/callmehorhe/konstanta/models"
	"github.com/gin-gonic/gin"
)

const token = "mySecertToken"

type Status struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	var input Status
	auth := c.GetHeader("authorization")

	if auth != token {
		newErrorResponse(c, http.StatusUnauthorized, "unauth")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid input")
		return
	}
	var err error
	switch input.Status {
	case models.StatusSuccess:
		err = h.services.UpdateStatus(input.Id, models.StatusSuccess)
	case models.StatusNotSuccess:
		err = h.services.UpdateStatus(input.Id, models.StatusNotSuccess)
	default:
		newErrorResponse(c, http.StatusInternalServerError, "wrong status")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "status changed",
		"status":  input.Status,
	})
}

func (h *Handler) GetStatusByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	status, err := h.services.GetStatusByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "not such transaction")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": status,
	})
}

func (h *Handler) CancelTransaction(c *gin.Context) {
	var id ID
	if err := c.BindJSON(&id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.services.UpdateStatus(id.Id, models.StatusCanceled); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "status changed",
		"status":  models.StatusCanceled,
	})
}
