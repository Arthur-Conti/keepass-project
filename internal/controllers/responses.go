package controllers

import (
	"github.com/gin-gonic/gin"
)

func sendError(ginContext *gin.Context, statusCode int, errorMessage string) {
	ginContext.JSON(statusCode, BaseErrorResponse{
		Success: false,
		Message: errorMessage,
	})
}

func sendSuccess(ginContext *gin.Context, statusCode int, data any) {
	ginContext.JSON(statusCode, BaseSuccessResponse{
		Success: true,
		Data: data,
	})
}

type BaseSuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type BaseErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
