package beget_test

import (
	"net/http"

	"github.com/h2non/gock"

	"github.com/KurosawaAngel/beget"
)

func setupGock(endpoint, input, output string) {
	gock.New(beget.DefaultBaseURL).
		Get(endpoint).
		MatchParam("login", "test").
		MatchParam("passwd", "test").
		MatchParam("input_format", "json").
		MatchParam("output_format", "json").
		MatchParam("input_data", input).
		Reply(http.StatusOK).
		JSON(output)
}

func newTestClient() *beget.Client {
	return beget.New("test", "test")
}
