package customers

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/core/models"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type Controller struct {
	projectClient *platform.ByProjectKeyRequestBuilder
}

func New(pc *platform.ByProjectKeyRequestBuilder) *Controller {
	return &Controller{
		projectClient: pc,
	}
}

func (c *Controller) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.executeGetCustomerByEmailRequest(ctx)

	shouldReturn := c.checkError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(200, customer)
}

func (*Controller) checkError(err error, ctx *gin.Context) bool {
	if err != nil {
		switch e := err.(type) {
		case *errors.NotFoundError:
			ctx.JSON(404, models.NewErrorResponse(404, errors.NotFoundErr, err))
			return true
		case *errors.InternalError:
			ctx.JSON(500, models.NewErrorResponse(500, errors.InternalErr, err))
			return true
		default:
			log.Printf("oops, this was unexpected: %s", e)
			ctx.JSON(500, models.NewErrorResponse(500, errors.InternalErr, err))
			return true
		}
	}
	return false
}

func (c *Controller) executeGetCustomerByEmailRequest(ctx *gin.Context) (platform.Customer, error) {
	customer, err := c.projectClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Query("email"))}).
		Execute(ctx)

	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return platform.Customer{}, errors.NewInternalError(msg, err)
	}

	if len(customer.Results) == 0 {
		msg := fmt.Sprintf("can't found customer with email %s", ctx.Query("email"))
		log.Println(msg)
		return platform.Customer{}, errors.NewNotFoundError(msg)
	}

	return customer.Results[0], nil
}
