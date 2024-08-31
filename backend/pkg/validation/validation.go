package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{
		v: validator.New(),
	}
}

func (v *Validator) Validate(in any) error {
	err := v.v.Struct(in)
	return validationErr(err)
}

func validationErr(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			return errors.New(errMessage(fe))
		}
	}
	return nil
}

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
