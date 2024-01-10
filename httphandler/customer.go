package httphandler

import (
	"github.com/flashkoef/go-ct-rest-api/errorhandler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
)

type CustomerHandler struct {
	customerService service.CustomerServicer
	checkError      errorhandler.ErrorHandler
	mapper          mapper.CustomerMapper
}

func NewCustomersHandler(
	customerService service.CustomerServicer,
	checkError errorhandler.ErrorHandler,
	mapper mapper.CustomerMapper,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		checkError:      checkError,
		mapper:          mapper,
	}
}
