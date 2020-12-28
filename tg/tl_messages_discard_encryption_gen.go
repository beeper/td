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

// MessagesDiscardEncryptionRequest represents TL type `messages.discardEncryption#edd923c5`.
// Cancels a request for creation and/or delete info on secret chat.
//
// See https://core.telegram.org/method/messages.discardEncryption for reference.
type MessagesDiscardEncryptionRequest struct {
	// Secret chat ID
	ChatID int
}

// MessagesDiscardEncryptionRequestTypeID is TL type id of MessagesDiscardEncryptionRequest.
const MessagesDiscardEncryptionRequestTypeID = 0xedd923c5

// String implements fmt.Stringer.
func (d *MessagesDiscardEncryptionRequest) String() string {
	if d == nil {
		return "MessagesDiscardEncryptionRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesDiscardEncryptionRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(d.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *MessagesDiscardEncryptionRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.discardEncryption#edd923c5 as nil")
	}
	b.PutID(MessagesDiscardEncryptionRequestTypeID)
	b.PutInt(d.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDiscardEncryptionRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.discardEncryption#edd923c5 to nil")
	}
	if err := b.ConsumeID(MessagesDiscardEncryptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.discardEncryption#edd923c5: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discardEncryption#edd923c5: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDiscardEncryptionRequest.
var (
	_ bin.Encoder = &MessagesDiscardEncryptionRequest{}
	_ bin.Decoder = &MessagesDiscardEncryptionRequest{}
)

// MessagesDiscardEncryption invokes method messages.discardEncryption#edd923c5 returning error if any.
// Cancels a request for creation and/or delete info on secret chat.
//
// Possible errors:
//  400 CHAT_ID_EMPTY: The provided chat ID is empty
//  400 ENCRYPTION_ALREADY_DECLINED: The secret chat was already declined
//  400 ENCRYPTION_ID_INVALID: The provided secret chat ID is invalid
//
// See https://core.telegram.org/method/messages.discardEncryption for reference.
func (c *Client) MessagesDiscardEncryption(ctx context.Context, chatid int) (bool, error) {
	var result BoolBox

	request := &MessagesDiscardEncryptionRequest{
		ChatID: chatid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
