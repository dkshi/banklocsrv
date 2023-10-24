package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) atms(c *gin.Context) {
	var filter map[string]interface{}
	if err := c.BindJSON(&filter); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	atms, err := h.services.Atms.GetAtms(&filter)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, atms)
}
