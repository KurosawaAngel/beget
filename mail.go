package beget

import (
	"context"
)

type Mailbox struct {
	Mailbox           string            `json:"mailbox"`
	Domain            string            `json:"domain"`
	SpamFilterStatus  SpamFilterStatus  `json:"spam_filter_status"`
	ForwardMailStatus ForwardMailStatus `json:"forward_mail_status"`
}

type ForwardMailbox struct {
	ForwardMailbox string `json:"forward_mailbox"`
}

// SpamFilterStatus is the status of spam filter.
type SpamFilterStatus int

const (
	SpamFilterStatusDisabled SpamFilterStatus = 0
	SpamFilterStatusEnabled                   = 1
)

// ForwardMailStatus is the status of forwarding mail.
type ForwardMailStatus string

const (
	ForwardMailStatusNoForward        ForwardMailStatus = "no_forward"
	ForwardMailStatusForward                            = "forward"
	ForwardMailStatusForwardAndDelete                   = "forward_and_delete"
)

// GetMailboxList returns all mailboxes on the given domain.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#getmailboxlist
func (c *Client) GetMailboxList(ctx context.Context, domain string) ([]Mailbox, error) {
	var response response[[]Mailbox]
	data := map[string]string{"domain": domain}
	if err := c.do(ctx, "mail/mailboxes", data, &response); err != nil {
		return nil, err
	}

	if response.hasErrors() {
		return nil, response.Answer.Errors
	}

	return response.Answer.Result, nil
}

// ChangeMailboxPassword changes password for the given mailbox.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#changemailboxpassword
func (c *Client) ChangeMailboxPassword(ctx context.Context, domain, mailbox, password string) error {
	var response response[bool]
	data := map[string]string{
		"domain":           domain,
		"mailbox":          mailbox,
		"mailbox_password": password,
	}
	if err := c.do(ctx, "mail/changeMailboxPassword", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// CreateMailbox creates a new mailbox on the given domain.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#createmailbox
func (c *Client) CreateMailbox(ctx context.Context, domain, mailbox, password string) error {
	var response response[bool]
	data := map[string]string{
		"domain":           domain,
		"mailbox":          mailbox,
		"mailbox_password": password,
	}
	if err := c.do(ctx, "mail/createMailbox", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// DropMailbox deletes the specified mailbox.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#dropmailbox
func (c *Client) DropMailbox(ctx context.Context, domain, mailbox string) error {
	var response response[bool]
	data := map[string]string{
		"domain":  domain,
		"mailbox": mailbox,
	}
	if err := c.do(ctx, "mail/dropMailbox", data, &response); err != nil {
		return err
	}
	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// ChangeMailboxSettings updates settings for the specified mailbox.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#changemailboxsettings
func (c *Client) ChangeMailboxSettings(ctx context.Context, domain, mailbox string, spamFilterStatus SpamFilterStatus, forwardMailStatus ForwardMailStatus) error {
	var response response[bool]
	data := map[string]any{
		"domain":              domain,
		"mailbox":             mailbox,
		"spam_filter_status":  spamFilterStatus,
		"forward_mail_status": forwardMailStatus,
	}
	if err := c.do(ctx, "mail/changeMailboxSettings", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// ForwardListAddMailbox adds a mailbox to the forwarding list.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#forwardlistaddmailbox
func (c *Client) ForwardListAddMailbox(ctx context.Context, domain, mailbox, forwardMailbox string) error {
	var response response[bool]
	data := map[string]string{
		"domain":          domain,
		"mailbox":         mailbox,
		"forward_mailbox": forwardMailbox,
	}
	if err := c.do(ctx, "mail/forwardListAddMailbox", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// ForwardListDeleteMailbox removes a mailbox from the forwarding list.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#forwardlistdeletemailbox
func (c *Client) ForwardListDeleteMailbox(ctx context.Context, domain, mailbox, forwardMailbox string) error {
	var response response[bool]
	data := map[string]string{
		"domain":          domain,
		"mailbox":         mailbox,
		"forward_mailbox": forwardMailbox,
	}
	if err := c.do(ctx, "mail/forwardListDeleteMailbox", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// ForwardListShow returns the forwarding list for the specified mailbox.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#forwardlistshow
func (c *Client) ForwardListShow(ctx context.Context, domain, mailbox string) ([]ForwardMailbox, error) {
	const endpoint = "mail/forwardListShow"

	var response response[[]ForwardMailbox]
	data := map[string]string{
		"domain":  domain,
		"mailbox": mailbox,
	}

	if err := c.do(ctx, endpoint, data, &response); err != nil {
		return nil, err
	}

	if response.hasErrors() {
		return nil, response.Answer.Errors
	}

	return response.Answer.Result, nil
}

// SetDomainMail sets up domain mail for the specified domain.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#setdomainmail
func (c *Client) SetDomainMail(ctx context.Context, domain, domainMailbox string) error {
	var response response[bool]
	data := map[string]string{
		"domain":         domain,
		"domain_mailbox": domainMailbox,
	}
	if err := c.do(ctx, "mail/setDomainMail", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}

// ClearDomainMail resets domain mail for the specified domain.
//
// If there is an beget error, it will be of type [Errors].
//
// Beget API docs: https://beget.com/en/kb/api/functions-for-work-with-mail#cleardomainmail
func (c *Client) ClearDomainMail(ctx context.Context, domain string) error {
	var response response[bool]
	data := map[string]string{
		"domain": domain,
	}
	if err := c.do(ctx, "mail/clearDomainMail", data, &response); err != nil {
		return err
	}

	if response.hasErrors() {
		return response.Answer.Errors
	}

	return nil
}
