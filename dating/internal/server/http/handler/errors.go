package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sterligov/otus_highload/dating/internal/domain"
)

var (
	ErrBadParam  = fmt.Errorf("bad request param")
	ErrBadClaims = fmt.Errorf("bad jwt claims")
)

func JSONError(c *gin.Context, err error) {
	var status int

	switch {
	case errors.Is(err, domain.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrBadClaims):
		status = http.StatusBadRequest
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(status, gin.H{"error": err.Error()})
}
