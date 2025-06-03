package beget

import (
	"fmt"
	"strings"
)

type response[T any] struct {
	Status    string    `json:"status"`
	ErrorText string    `json:"error_text"`
	ErrorCode int       `json:"error_code"`
	Answer    answer[T] `json:"answer"`
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
	return fmt.Sprintf("code=%d: text=%q", e.Code, e.Text)
}

type Errors []*Error

func (es Errors) Error() string {
	if len(es) == 1 {
		return "beget: " + es[0].Error()
	}

	var b strings.Builder
	b.WriteString("beget: ")
	b.WriteString(es[0].Error())
	for _, err := range es[1:] {
		b.WriteString(", ")
		b.WriteString(err.Error())
	}
	return b.String()
}

func (es Errors) Unwrap() []error {
	errs := make([]error, len(es))
	for i, err := range es {
		errs[i] = err
	}
	return errs
}
