package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status int            `json:"status"`
	Data   map[string]any `json:"data"`
	Error  error          `json:"error"`
}

type Responder struct {
}

const UNAUTHORIZED_MESSAGE = "Authorization required!!"

func (r *Responder) ResponseSuccess(c *gin.Context, data *map[string]interface{}) {
	c.JSON(http.StatusOK, ApiResponse{Status: http.StatusOK, Data: *data})
}

func (r *Responder) ResponseCreateSuccess(c *gin.Context) {
	c.JSON(http.StatusCreated, ApiResponse{Status: http.StatusCreated})
}
func (r *Responder) ResponseError(c *gin.Context, str string) {
	c.JSON(http.StatusBadRequest, ApiResponse{Status: http.StatusBadRequest, Error: errors.New(str)})
}

func (r *Responder) ResponseServerError(c *gin.Context, str string) {
	c.JSON(http.StatusInternalServerError, ApiResponse{Status: http.StatusInternalServerError, Error: errors.New(str)})
}

func (r *Responder) ResponseUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ApiResponse{Status: http.StatusUnauthorized, Error: errors.New(UNAUTHORIZED_MESSAGE)})
}
