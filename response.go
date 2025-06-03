package beget

import (
	"fmt"
	"unsafe"
)

type response[T any] struct {
	Status string `json:"status"`
	Error
	Answer answer[T] `json:"answer"`
}

func (r response[T]) err() error {
	switch {
	case r.Status == "error":
		return Errors{&r.Error}
	case r.Answer.Status == "error":
		return r.Answer.Errors
	default:
		return nil
	}
}

type answer[T any] struct {
	Status string `json:"status"`
	Result T      `json:"result"`
	Errors Errors `json:"errors"`
}

// Error represents a beget error and
// implements the error interface.
type Error struct {
	Code any    `json:"error_code"` // Execellent beget API can return int or string.
	Text string `json:"error_text"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%v: text=%q", e.Code, e.Text)
}

type Errors []*Error

func (es Errors) Error() string {
	switch len(es) {
	case 0:
		return "beget: no errors"
	case 1:
		return "beget: " + es[0].Error()
	default:
		// errors.Join realization
		b := []byte("beget: ")
		b = append(b, es[0].Error()...)
		for _, err := range es[1:] {
			b = append(b, '\n')
			b = append(b, err.Error()...)
		}

		return unsafe.String(&b[0], len(b))
	}
}

func (es Errors) Unwrap() []error {
	errs := make([]error, len(es))
	for i, err := range es {
		errs[i] = err
	}
	return errs
}
