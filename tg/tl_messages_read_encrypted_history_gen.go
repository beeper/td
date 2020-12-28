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

// MessagesReadEncryptedHistoryRequest represents TL type `messages.readEncryptedHistory#7f4b690a`.
// Marks message history within a secret chat as read.
//
// See https://core.telegram.org/method/messages.readEncryptedHistory for reference.
type MessagesReadEncryptedHistoryRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// Maximum date value for received messages in history
	MaxDate int
}

// MessagesReadEncryptedHistoryRequestTypeID is TL type id of MessagesReadEncryptedHistoryRequest.
const MessagesReadEncryptedHistoryRequestTypeID = 0x7f4b690a

// String implements fmt.Stringer.
func (r *MessagesReadEncryptedHistoryRequest) String() string {
	if r == nil {
		return "MessagesReadEncryptedHistoryRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesReadEncryptedHistoryRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(r.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMaxDate: ")
	sb.WriteString(fmt.Sprint(r.MaxDate))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *MessagesReadEncryptedHistoryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readEncryptedHistory#7f4b690a as nil")
	}
	b.PutID(MessagesReadEncryptedHistoryRequestTypeID)
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readEncryptedHistory#7f4b690a: field peer: %w", err)
	}
	b.PutInt(r.MaxDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadEncryptedHistoryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readEncryptedHistory#7f4b690a to nil")
	}
	if err := b.ConsumeID(MessagesReadEncryptedHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readEncryptedHistory#7f4b690a: %w", err)
	}
	{
		if err := r.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.readEncryptedHistory#7f4b690a: field peer: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readEncryptedHistory#7f4b690a: field max_date: %w", err)
		}
		r.MaxDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReadEncryptedHistoryRequest.
var (
	_ bin.Encoder = &MessagesReadEncryptedHistoryRequest{}
	_ bin.Decoder = &MessagesReadEncryptedHistoryRequest{}
)

// MessagesReadEncryptedHistory invokes method messages.readEncryptedHistory#7f4b690a returning error if any.
// Marks message history within a secret chat as read.
//
// Possible errors:
//  400 MSG_WAIT_FAILED: A waiting call returned an error
//
// See https://core.telegram.org/method/messages.readEncryptedHistory for reference.
func (c *Client) MessagesReadEncryptedHistory(ctx context.Context, request *MessagesReadEncryptedHistoryRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
