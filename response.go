package beget

import (
	"fmt"
	"strings"
)

type response[T any] struct {
	Status string    `json:"status"`
	Answer answer[T] `json:"answer"`
}

func (r response[T]) hasErrors() bool {
	return len(r.Answer.Errors) > 0
}

type answer[T any] struct {
	Status string `json:"status"`
	Result T      `json:"result"`
	Errors Errors `json:"errors"`
}

// Error represents a beget error and
// implements the error interface.
type Error struct {
	Code int    `json:"error_code"`
	Text string `json:"error_text"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Text)
}

type Errors []*Error

func (e Errors) Error() string {
	switch len(e) {
	case 0:
		return "beget: no errors"
	case 1:
		return "beget: " + e[0].Error()
	default:
		errs := make([]string, len(e)+1)
		errs[0] = "beget: "
		for i := range e {
			errs[i+1] = e[i].Error()
		}
		return strings.Join(errs, ", ")
	}
}
