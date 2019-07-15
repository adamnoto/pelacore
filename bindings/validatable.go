package bindings

import "errors"

// Validatable is an interface for validatable bindings, to check
// if there is some missing attributes, or misformatting in the
// request payload
type Validatable interface {
	Validate() error
}

// ErrNotValidatable is given when an object is not Validatable
var ErrNotValidatable = errors.New("Unvalidatable object")

// Validator is the validator
type Validator struct{}

// Validate validates the object
func (v *Validator) Validate(i interface{}) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return ErrNotValidatable
}
