package beget

import "fmt"

type response[T any] struct {
	Status string    `json:"status"`
	Answer answer[T] `json:"answer"`
}

type answer[T any] struct {
	Status string   `json:"status"`
	Result T        `json:"result"`
	Errors []*Error `json:"errors"`
}

// Error represents a beget error and
// implements the error interface.
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorText string `json:"error_text"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("beget: %d: %s", e.ErrorCode, e.ErrorText)
}
