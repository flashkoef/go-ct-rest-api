package errorhandler

import (
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	CheckError(err error, ctx *gin.Context) bool
}

type CheckErr struct{}

func New() *CheckErr {
	return &CheckErr{}
}

func (c *CheckErr) CheckError(err error, ctx *gin.Context) bool {
	if err != nil {
		switch e := err.(type) {
		case *NotFoundError:
			ctx.JSON(http.StatusNotFound, model.NewErrorResponse("Resource not found.", "NOT_FOUND", err))

			return true
		case *InternalError:
			ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse("An unexpected error occurred.", InternalErr, err))
			return true
		case *CtpError:
			ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse("An unexpected error occurred.", InternalErr, err))

			return true
		default:
			log.Printf("Oops, this was unexpected: %s", e)
			ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse("Oops, this was unexpected.", InternalErr, err))

			return true
		}
	}

	return false
}
