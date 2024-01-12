package httphandler

import (
	"github.com/flashkoef/go-ct-rest-api/errorhandler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/flashkoef/go-ct-rest-api/validator"
)

type CustomerHandler struct {
	customerService service.CustomerServicer
	errorHandler    errorhandler.ErrorHandler
	mapper          mapper.CustomerMapper
	validator       validator.Validator
}

func NewCustomersHandler(
	customerService service.CustomerServicer,
	errorHandler errorhandler.ErrorHandler,
	mapper mapper.CustomerMapper,
	validator validator.Validator,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		errorHandler:    errorHandler,
		mapper:          mapper,
		validator:       validator,
	}
}
