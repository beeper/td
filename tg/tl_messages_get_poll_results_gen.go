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

// MessagesGetPollResultsRequest represents TL type `messages.getPollResults#73bb643b`.
// Get poll results
//
// See https://core.telegram.org/method/messages.getPollResults for reference.
type MessagesGetPollResultsRequest struct {
	// Peer where the poll was found
	Peer InputPeerClass
	// Message ID of poll message
	MsgID int
}

// MessagesGetPollResultsRequestTypeID is TL type id of MessagesGetPollResultsRequest.
const MessagesGetPollResultsRequestTypeID = 0x73bb643b

// String implements fmt.Stringer.
func (g *MessagesGetPollResultsRequest) String() string {
	if g == nil {
		return "MessagesGetPollResultsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetPollResultsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(g.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(g.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetPollResultsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPollResults#73bb643b as nil")
	}
	b.PutID(MessagesGetPollResultsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getPollResults#73bb643b: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPollResults#73bb643b: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPollResultsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPollResults#73bb643b to nil")
	}
	if err := b.ConsumeID(MessagesGetPollResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetPollResultsRequest.
var (
	_ bin.Encoder = &MessagesGetPollResultsRequest{}
	_ bin.Decoder = &MessagesGetPollResultsRequest{}
)

// MessagesGetPollResults invokes method messages.getPollResults#73bb643b returning error if any.
// Get poll results
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/messages.getPollResults for reference.
func (c *Client) MessagesGetPollResults(ctx context.Context, request *MessagesGetPollResultsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
