package bindings

import "strings"

// PayloadErrors represent all errors in the payload
type PayloadErrors struct {
	errs []error
}

// Append registers a new error into the list of error
func (pe *PayloadErrors) Append(err error) {
	pe.errs = append(pe.errs, err)
}

// Len returns the number of errors
func (pe *PayloadErrors) Len() int {
	return len(pe.errs)
}

func (pe *PayloadErrors) Error() string {
	errstr := []string{}
	for _, e := range pe.errs {
		errstr = append(errstr, e.Error())
	}
	return strings.Join(errstr, ". ")
}
