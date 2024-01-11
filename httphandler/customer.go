package httphandler

import (
	"github.com/flashkoef/go-ct-rest-api/errorhandler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
)

type CustomerHandler struct {
	customerService service.CustomerServicer
	errorHandler    errorhandler.ErrorHandler
	mapper          mapper.CustomerMapper
}

func NewCustomersHandler(
	customerService service.CustomerServicer,
	errorHandler errorhandler.ErrorHandler,
	mapper mapper.CustomerMapper,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		errorHandler:    errorHandler,
		mapper:          mapper,
	}
}
