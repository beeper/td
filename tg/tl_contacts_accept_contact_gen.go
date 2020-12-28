// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// ContactsAcceptContactRequest represents TL type `contacts.acceptContact#f831a20f`.
// If the peer settings¹ of a new user allow us to add him as contact, add that user as contact
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
type ContactsAcceptContactRequest struct {
	// The user to add as contact
	ID InputUserClass
}

// ContactsAcceptContactRequestTypeID is TL type id of ContactsAcceptContactRequest.
const ContactsAcceptContactRequestTypeID = 0xf831a20f

// String implements fmt.Stringer.
func (a *ContactsAcceptContactRequest) String() string {
	if a == nil {
		return "ContactsAcceptContactRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsAcceptContactRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(a.ID.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *ContactsAcceptContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.acceptContact#f831a20f as nil")
	}
	b.PutID(ContactsAcceptContactRequestTypeID)
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *ContactsAcceptContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.acceptContact#f831a20f to nil")
	}
	if err := b.ConsumeID(ContactsAcceptContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: field id: %w", err)
		}
		a.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsAcceptContactRequest.
var (
	_ bin.Encoder = &ContactsAcceptContactRequest{}
	_ bin.Decoder = &ContactsAcceptContactRequest{}
)

// ContactsAcceptContact invokes method contacts.acceptContact#f831a20f returning error if any.
// If the peer settings¹ of a new user allow us to add him as contact, add that user as contact
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// Possible errors:
//  400 CONTACT_ADD_MISSING: Contact to add is missing
//  400 CONTACT_ID_INVALID: The provided contact ID is invalid
//  400 CONTACT_REQ_MISSING: Missing contact request
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
func (c *Client) ContactsAcceptContact(ctx context.Context, id InputUserClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ContactsAcceptContactRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
