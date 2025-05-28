package beget_test

import (
	"context"
	"errors"

	"github.com/KurosawaAngel/beget"
)

func Example() {
	ctx := context.Background()
	// create a new Beget API client.
	c := beget.New("username", "password")
	// use api methods as you need
	err := c.DropMailbox(ctx, "domain", "mailbox")
	if err != nil {
		var e *beget.Error
		if errors.As(err, &e) {
			// handle beget error
		}
		// handle other errors
	}
}
