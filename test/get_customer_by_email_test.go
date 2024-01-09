// http test
package test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHttpGetCustomerByEmail(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	errorHandler := error_handler.New()
	customerService := service.NewCustomerService(connector.New().GetProjectClient(), errorHandler)
	customerHandler := http_handler.NewCustomersHandler(customerService, errorHandler, mapper.NewMapper())

	resRecorder := httptest.NewRecorder()
	ctx := GetTestGinContext(resRecorder)

	pathParams := []gin.Param{
		{
			Key:   "email",
			Value: "jane.doe@test.de",
		},
	}

	MockJsonGetCustomerByEmail(ctx, pathParams)

	customerHandler.GetCustomerByEmail(ctx)

	assert.EqualValues(t, http.StatusOK, resRecorder.Code)

	var customer model.Customer
	err := json.NewDecoder(resRecorder.Body).Decode(&customer)
	if err != nil {
		log.Fatalf("error while decoding response body %s ", err)
	}

	assert.Equal(t, expectedCustomer, customer)
}

func TestHttpGetCustomerByEmailNotFound(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	errorHandler := error_handler.New()
	customerService := service.NewCustomerService(connector.New().GetProjectClient(), errorHandler)
	customerHandler := http_handler.NewCustomersHandler(customerService, errorHandler, mapper.NewMapper())

	resRecorder := httptest.NewRecorder()
	ctx := GetTestGinContext(resRecorder)

	pathParams := []gin.Param{
		{
			Key:   "email",
			Value: "unknown@test.de",
		},
	}

	MockJsonGetCustomerByEmail(ctx, pathParams)

	customerHandler.GetCustomerByEmail(ctx)

	assert.EqualValues(t, http.StatusNotFound, resRecorder.Code)

	var errResponse model.ErrorResponse
	err := json.NewDecoder(resRecorder.Body).Decode(&errResponse)
	if err != nil {
		log.Fatalf("error while decoding response body %s ", err)
	}

	assert.Equal(t, expectedErrorResponse, errResponse)
}

func GetTestGinContext(resRecorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(resRecorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func MockJsonGetCustomerByEmail(ctx *gin.Context, pathParams gin.Params) {
	ctx.Request.Method = "GET"
	ctx.Request.Header.Set("Content-Type", "application/json")
	// ctx.Set("email", "jane.doe@test.de") // when invoking the endpoint
	ctx.Params = pathParams
}
