package beget_test

import (
	"testing"

	"github.com/KurosawaAngel/beget"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetMailboxList(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/getMailboxList",
		`{"domain":"test.com"}`,
		`{"status": "success", "answer": {"status": "success", "result": [{"mailbox": "testik", "domain": "test.com", "spam_filter_status": 1, "forward_mail_status": "no_forward"}]}}`,
	)
	responseCheck := []beget.Mailbox{
		{
			Mailbox:           "testik",
			Domain:            "test.com",
			SpamFilterStatus:  beget.SpamFilterStatusEnabled,
			ForwardMailStatus: beget.ForwardMailStatusNoForward,
		},
	}
	c := newTestClient()

	res, err := c.GetMailboxList(t.Context(), "test.com")

	assert.Nil(t, err)
	assert.Equal(t, responseCheck, res)
}

func TestClient_ChangeMailboxPassword(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/changeMailboxPassword",
		`{"domain":"test.com","mailbox":"testik","mailbox_password":"testik"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.ChangeMailboxPassword(t.Context(), "test.com", "testik", "testik")

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_CreateMailbox(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/createMailbox",
		`{"domain":"test.com","mailbox":"testik","mailbox_password":"testik"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.CreateMailbox(t.Context(), "test.com", "testik", "testik")

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_DropMailbox(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/dropMailbox",
		`{"domain":"test.com","mailbox":"testik"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.DropMailbox(t.Context(), "test.com", "testik")

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_ChangeMailboxSettings(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/changeMailboxSettings",
		`{"domain":"test.com","mailbox":"testik","spam_filter_status":1,"spam_filter":20,"forward_mail_status":"no_forward"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.ChangeMailboxSettings(
		t.Context(),
		"test.com",
		"testik",
		beget.SpamFilterStatusEnabled,
		beget.ForwardMailStatusNoForward,
	)

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_ForwardListAddMailbox(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/forwardListAddMailbox",
		`{"domain":"test.com","mailbox":"testik","forward_mailbox":"testik2@test.com"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.ForwardListAddMailbox(
		t.Context(),
		"test.com",
		"testik",
		"testik2@test.com",
	)

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_ForwardListDeleteMailbox(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/forwardListDeleteMailbox",
		`{"domain":"test.com","mailbox":"testik","forward_mailbox":"testik2@test.com"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.ForwardListDeleteMailbox(
		t.Context(),
		"test.com",
		"testik",
		"testik2@test.com",
	)

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_ForwardListShow(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/forwardListShow",
		`{"domain":"test.com","mailbox":"testik"}`,
		`{"status": "success", "answer": {"status": "success", "result": [{"forward_mailbox": "testik2@test.com"}]}}`,
	)
	responseCheck := []beget.ForwardMailbox{
		{
			ForwardMailbox: "testik2@test.com",
		},
	}
	c := newTestClient()

	res, err := c.ForwardListShow(
		t.Context(),
		"test.com",
		"testik",
	)

	assert.Nil(t, err)
	assert.Equal(t, responseCheck, res)
}

func TestClient_SetDomainMail(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/setDomainMail",
		`{"domain":"test.com","domain_mailbox":"testik"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.SetDomainMail(
		t.Context(),
		"test.com",
		"testik",
	)

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_ClearDomainMail(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/clearDomainMail",
		`{"domain":"test.com"}`,
		`{"status": "success", "answer": {"status": "success", "result": true}}`,
	)
	c := newTestClient()

	res, err := c.ClearDomainMail(
		t.Context(),
		"test.com",
	)

	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func TestClient_Error(t *testing.T) {
	defer gock.Off()
	setupGock(
		"mail/clearDomainMail",
		`{"domain":"test.com"}`,
		`{"status": "success", "answer": {"status": "error", "errors": [{"error_code": "INVALID_DATA", "error_text": "Incorrect input data"}]}}`,
	)
	responseCheck := beget.Errors{&beget.Error{Code: "INVALID_DATA", Text: "Incorrect input data"}}
	c := newTestClient()

	_, err := c.ClearDomainMail(
		t.Context(),
		"test.com",
	)

	assert.Equal(t, responseCheck, err)
}
