package customers

import (
	"fmt"
	"log"

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
	customer, err := c.projectClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Query("email"))}).
		Execute(ctx)

	if err != nil {
		log.Fatalf("error while execute request to ctp %s", err)
	}

	if len(customer.Results) == 0 {
		msg := fmt.Sprintf("can't found customer with email %s", ctx.Query("email"))
		log.Println(msg)
		ctx.JSON(404, models.NewNotFoundRes(msg))
		return
	}

	ctx.JSON(200, customer.Results)
}
