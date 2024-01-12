package errorhandler

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/customerror"
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
		case *customerror.NotFoundError:
			ctx.JSON(http.StatusNotFound, model.NewErrorResponse("Resource not found", customerror.NotFoundErr, e))

			return true
		case *customerror.InternalError:
			ctx.JSON(
				http.StatusInternalServerError,
				model.NewErrorResponse("Internal server error occurred", customerror.InternalErr, e),
			)

			return true
		case *customerror.CtpError:
			ctx.JSON(
				http.StatusInternalServerError,
				model.NewErrorResponse(
					"An error has occurred when interacting with the commercetools platform",
					customerror.CtpErr,
					e,
				),
			)

			return true
		case *customerror.ValidationError:
			ctx.JSON(
				http.StatusBadRequest,
				model.NewErrorResponse("Bad Request", customerror.ValidationErr, e),
			)

			return true
		default:
			ctx.JSON(
				http.StatusInternalServerError,
				model.NewErrorResponse("Oops, this was unexpected", customerror.InternalErr, e),
			)

			return true
		}
	}

	return false
}
