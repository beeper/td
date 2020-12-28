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

// ChannelsToggleSignaturesRequest represents TL type `channels.toggleSignatures#1f69b606`.
// Enable/disable message signatures in channels
//
// See https://core.telegram.org/method/channels.toggleSignatures for reference.
type ChannelsToggleSignaturesRequest struct {
	// Channel
	Channel InputChannelClass
	// Value
	Enabled bool
}

// ChannelsToggleSignaturesRequestTypeID is TL type id of ChannelsToggleSignaturesRequest.
const ChannelsToggleSignaturesRequestTypeID = 0x1f69b606

// String implements fmt.Stringer.
func (t *ChannelsToggleSignaturesRequest) String() string {
	if t == nil {
		return "ChannelsToggleSignaturesRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsToggleSignaturesRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(t.Channel.String())
	sb.WriteString(",\n")
	sb.WriteString("\tEnabled: ")
	sb.WriteString(fmt.Sprint(t.Enabled))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleSignaturesRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSignatures#1f69b606 as nil")
	}
	b.PutID(ChannelsToggleSignaturesRequestTypeID)
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleSignatures#1f69b606: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleSignatures#1f69b606: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleSignaturesRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSignatures#1f69b606 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleSignaturesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsToggleSignaturesRequest.
var (
	_ bin.Encoder = &ChannelsToggleSignaturesRequest{}
	_ bin.Decoder = &ChannelsToggleSignaturesRequest{}
)

// ChannelsToggleSignatures invokes method channels.toggleSignatures#1f69b606 returning error if any.
// Enable/disable message signatures in channels
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//
// See https://core.telegram.org/method/channels.toggleSignatures for reference.
func (c *Client) ChannelsToggleSignatures(ctx context.Context, request *ChannelsToggleSignaturesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
