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

// MessagesReportRequest represents TL type `messages.report#bd82b658`.
// Report a message in a chat for violation of telegram's Terms of Service
//
// See https://core.telegram.org/method/messages.report for reference.
type MessagesReportRequest struct {
	// Peer
	Peer InputPeerClass
	// IDs of messages to report
	ID []int
	// Why are these messages being reported
	Reason ReportReasonClass
}

// MessagesReportRequestTypeID is TL type id of MessagesReportRequest.
const MessagesReportRequestTypeID = 0xbd82b658

// String implements fmt.Stringer.
func (r *MessagesReportRequest) String() string {
	if r == nil {
		return "MessagesReportRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesReportRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(r.Peer.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range r.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tReason: ")
	sb.WriteString(r.Reason.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *MessagesReportRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.report#bd82b658 as nil")
	}
	b.PutID(MessagesReportRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.report#bd82b658: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.report#bd82b658: field peer: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode messages.report#bd82b658: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.report#bd82b658: field reason: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.report#bd82b658 to nil")
	}
	if err := b.ConsumeID(MessagesReportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.report#bd82b658: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#bd82b658: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#bd82b658: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.report#bd82b658: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#bd82b658: field reason: %w", err)
		}
		r.Reason = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReportRequest.
var (
	_ bin.Encoder = &MessagesReportRequest{}
	_ bin.Decoder = &MessagesReportRequest{}
)

// MessagesReport invokes method messages.report#bd82b658 returning error if any.
// Report a message in a chat for violation of telegram's Terms of Service
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.report for reference.
func (c *Client) MessagesReport(ctx context.Context, request *MessagesReportRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
