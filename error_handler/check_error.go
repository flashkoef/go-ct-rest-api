package error_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type ErrorHandler interface {
	CheckInternError(err error, ctx *gin.Context) bool
	CheckCtSdkErrorForNonPagedResponse(err error, ctx *gin.Context) (bool, error)
	CheckCtSdkErrorForPagedResponse(err error, result *platform.ProductProjectionPagedQueryResponse, ctx *gin.Context) (bool, error)
}

type CheckError struct{}

func New() *CheckError {
	return &CheckError{}
}

func (ce *CheckError) CheckInternError(err error, ctx *gin.Context) bool {
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

func (ce *CheckError) CheckCtSdkErrorForNonPagedResponse(err error, ctx *gin.Context) (bool, error) {
	if err != nil {
		log.Printf("error while execute request to ctp %s", err)
		if err.Error() == platform.ErrNotFound.Error() {
			msg := fmt.Sprintf("can't found customer with id %s", ctx.Param("customerID"))
			return true, NewNotFoundErrorWithOriginalErr(msg, err.Error())
		}
	}
	return false, nil
}

func (ce *CheckError) CheckCtSdkErrorForPagedResponse(err error, result *platform.ProductProjectionPagedQueryResponse, ctx *gin.Context) (bool, error) {
	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return true, NewInternalError(msg, err)
	}

	if len(result.Results) == 0 {
		msg := fmt.Sprintf("can't found product projection with slug %s", ctx.Query("slug"))
		log.Println(msg)
		return true, NewNotFoundError(msg)
	}
	return false, nil
}
