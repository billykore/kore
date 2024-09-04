package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

// New returns new Validator.
func New() *Validator {
	return &Validator{
		v: validator.New(),
	}
}

// Validate any request. The request is expected to be a struct.
func (v *Validator) Validate(req any) error {
	err := v.v.Struct(req)
	return validationErr(err)
}

// validationErr returns new validation error with custom message.
func validationErr(err error) error {
	var ve validator.ValidationErrors
	var errs error
	if errors.As(err, &ve) {
		for _, fe := range ve {
			errs = errors.Join(errors.New(errMessage(fe)))
		}
	}
	return errs
}

// errMessage return custom error message base on the validation tag.
func errMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", fe.Field())
	case "len":
		return fmt.Sprintf("%s length must be %s", fe.Field(), fe.Param())
	case "number":
		return fmt.Sprintf("%s must be number", fe.Field())
	}
	return fmt.Sprintf("%s format is invalid", fe.Field())
}
