package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type GinHandler = func(ctx *gin.Context)
type ApiHandler = func(ctx *gin.Context) (*ApiResponse, error)

type ApiResponse struct {
	StatusCode int
	Result     any
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func FormatErrorResponse(message, details string) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Details: details,
	}
}
func ApiAdapter(handler ApiHandler) GinHandler {
	return func(ctx *gin.Context) {
		result, err := handler(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				FormatErrorResponse("Programming error", err.Error()),
			)
			return
		}
		ctx.JSON(result.StatusCode, result.Result)
	}
}

func HandleDBErrors(err error) (*ApiResponse, error) {
	if err == pgx.ErrNoRows {
		return &ApiResponse{
				StatusCode: http.StatusNotFound,
				Result: FormatErrorResponse(
					"Resource not found",
					"The requested resource was not found in remote server"),
			},
			nil
	}

	if pgxErr, ok := err.(*pgconn.PgError); ok {
		switch pgxErr.Code {
		case "23505": // unique_violation
			return &ApiResponse{
				StatusCode: http.StatusConflict,
				Result: FormatErrorResponse(
					"Resource with the same unique identifier already exists",
					"Resource conflict",
				),
			}, nil
		}
	}
	return nil, err

}
