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

// ContactsBlockRequest represents TL type `contacts.block#68cc1411`.
// Adds the user to the blacklist.
//
// See https://core.telegram.org/method/contacts.block for reference.
type ContactsBlockRequest struct {
	// User ID
	ID InputPeerClass
}

// ContactsBlockRequestTypeID is TL type id of ContactsBlockRequest.
const ContactsBlockRequestTypeID = 0x68cc1411

// String implements fmt.Stringer.
func (b *ContactsBlockRequest) String() string {
	if b == nil {
		return "ContactsBlockRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsBlockRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(b.ID.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *ContactsBlockRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.block#68cc1411 as nil")
	}
	buf.PutID(ContactsBlockRequestTypeID)
	if b.ID == nil {
		return fmt.Errorf("unable to encode contacts.block#68cc1411: field id is nil")
	}
	if err := b.ID.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.block#68cc1411: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *ContactsBlockRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.block#68cc1411 to nil")
	}
	if err := buf.ConsumeID(ContactsBlockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.block#68cc1411: %w", err)
	}
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.block#68cc1411: field id: %w", err)
		}
		b.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsBlockRequest.
var (
	_ bin.Encoder = &ContactsBlockRequest{}
	_ bin.Decoder = &ContactsBlockRequest{}
)

// ContactsBlock invokes method contacts.block#68cc1411 returning error if any.
// Adds the user to the blacklist.
//
// Possible errors:
//  400 CONTACT_ID_INVALID: The provided contact ID is invalid
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/contacts.block for reference.
func (c *Client) ContactsBlock(ctx context.Context, id InputPeerClass) (bool, error) {
	var result BoolBox

	request := &ContactsBlockRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
