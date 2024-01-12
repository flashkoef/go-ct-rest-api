package validator

import (
	"fmt"

	"github.com/flashkoef/go-ct-rest-api/customerror"
	"github.com/golodash/galidator"
)

type Validator interface {
	ValidateEmail(email string) error
}

type Val struct{}

func New() *Val {
	return &Val{}
}

func (v *Val) ValidateEmail(email string) error {
	var g = galidator.New()
	var emailValidator = g.Validator(g.RuleSet().Email())
	output := emailValidator.Validate(email)

	if output != nil {
		return customerror.NewValidationError(
			"email validation failed",
			fmt.Errorf("the provided email '%s' is not valid: %s", email, output),
		)
	}

	return nil
}
