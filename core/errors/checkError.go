package errors

import (
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/models"
	"github.com/gin-gonic/gin"
)

func CheckError(err error, ctx *gin.Context) bool {
	if err != nil {
		switch e := err.(type) {
		case *NotFoundError:
			ctx.JSON(404, models.NewErrorResponse(404, NotFoundErr, err))
			return true
		case *InternalError:
			ctx.JSON(500, models.NewErrorResponse(500, InternalErr, err))
			return true
		default:
			log.Printf("oops, this was unexpected: %s", e)
			ctx.JSON(500, models.NewErrorResponse(500, InternalErr, err))
			return true
		}
	}
	return false
}
