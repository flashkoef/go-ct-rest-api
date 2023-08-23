package customers

import (
	"fmt"
	"log"

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

	ctx.JSON(200, gin.H{
		"customer": customer.Results[0],
	})
}
