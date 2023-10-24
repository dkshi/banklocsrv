package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) offices(c *gin.Context) {
	var filter map[string]interface{}

	if err := c.BindJSON(&filter); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	offices, err := h.services.Offices.GetOffices(&filter)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, offices)
}
