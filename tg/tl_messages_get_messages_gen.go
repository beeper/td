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

// MessagesGetMessagesRequest represents TL type `messages.getMessages#63c66506`.
// Returns the list of messages by their IDs.
//
// See https://core.telegram.org/method/messages.getMessages for reference.
type MessagesGetMessagesRequest struct {
	// Message ID list
	ID []InputMessageClass
}

// MessagesGetMessagesRequestTypeID is TL type id of MessagesGetMessagesRequest.
const MessagesGetMessagesRequestTypeID = 0x63c66506

// String implements fmt.Stringer.
func (g *MessagesGetMessagesRequest) String() string {
	if g == nil {
		return "MessagesGetMessagesRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetMessagesRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range g.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessages#63c66506 as nil")
	}
	b.PutID(MessagesGetMessagesRequestTypeID)
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode messages.getMessages#63c66506: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getMessages#63c66506: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessages#63c66506 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessages#63c66506: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessages#63c66506: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getMessages#63c66506: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetMessagesRequest.
var (
	_ bin.Encoder = &MessagesGetMessagesRequest{}
	_ bin.Decoder = &MessagesGetMessagesRequest{}
)

// MessagesGetMessages invokes method messages.getMessages#63c66506 returning error if any.
// Returns the list of messages by their IDs.
//
// See https://core.telegram.org/method/messages.getMessages for reference.
// Can be used by bots.
func (c *Client) MessagesGetMessages(ctx context.Context, id []InputMessageClass) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	request := &MessagesGetMessagesRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
