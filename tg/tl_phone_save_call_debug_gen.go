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

// PhoneSaveCallDebugRequest represents TL type `phone.saveCallDebug#277add7e`.
// Send phone call debug data to server
//
// See https://core.telegram.org/method/phone.saveCallDebug for reference.
type PhoneSaveCallDebugRequest struct {
	// Phone call
	Peer InputPhoneCall
	// Debug statistics obtained from libtgvoip
	Debug DataJSON
}

// PhoneSaveCallDebugRequestTypeID is TL type id of PhoneSaveCallDebugRequest.
const PhoneSaveCallDebugRequestTypeID = 0x277add7e

// String implements fmt.Stringer.
func (s *PhoneSaveCallDebugRequest) String() string {
	if s == nil {
		return "PhoneSaveCallDebugRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneSaveCallDebugRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(s.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDebug: ")
	sb.WriteString(s.Debug.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *PhoneSaveCallDebugRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallDebug#277add7e as nil")
	}
	b.PutID(PhoneSaveCallDebugRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field peer: %w", err)
	}
	if err := s.Debug.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field debug: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSaveCallDebugRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallDebug#277add7e to nil")
	}
	if err := b.ConsumeID(PhoneSaveCallDebugRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field peer: %w", err)
		}
	}
	{
		if err := s.Debug.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field debug: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneSaveCallDebugRequest.
var (
	_ bin.Encoder = &PhoneSaveCallDebugRequest{}
	_ bin.Decoder = &PhoneSaveCallDebugRequest{}
)

// PhoneSaveCallDebug invokes method phone.saveCallDebug#277add7e returning error if any.
// Send phone call debug data to server
//
// Possible errors:
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//  400 DATA_JSON_INVALID: The provided JSON data is invalid
//
// See https://core.telegram.org/method/phone.saveCallDebug for reference.
func (c *Client) PhoneSaveCallDebug(ctx context.Context, request *PhoneSaveCallDebugRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
