//go:build example

package beget_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/KurosawaAngel/beget"
)

func Example() {
	ctx := context.Background()
	// create a new Beget API client.
	c := beget.New("username", "password")
	// use api methods as you need
	res, err := c.DropMailbox(ctx, "domain", "mailbox")
	if err != nil {
		var e beget.Errors
		if errors.As(err, &e) {
			// handle beget errors
		}
		// handle other errors
	}
	fmt.Println(res)
	// Output: true
}
