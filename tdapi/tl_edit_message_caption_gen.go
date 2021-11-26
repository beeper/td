// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// EditMessageCaptionRequest represents TL type `editMessageCaption#44d2f92e`.
type EditMessageCaptionRequest struct {
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message reply markup; pass null if none; for bots only
	ReplyMarkup ReplyMarkupClass
	// New message content caption; 0-GetOption("message_caption_length_max") characters;
	// pass null to remove caption
	Caption FormattedText
}

// EditMessageCaptionRequestTypeID is TL type id of EditMessageCaptionRequest.
const EditMessageCaptionRequestTypeID = 0x44d2f92e

// Ensuring interfaces in compile-time for EditMessageCaptionRequest.
var (
	_ bin.Encoder     = &EditMessageCaptionRequest{}
	_ bin.Decoder     = &EditMessageCaptionRequest{}
	_ bin.BareEncoder = &EditMessageCaptionRequest{}
	_ bin.BareDecoder = &EditMessageCaptionRequest{}
)

func (e *EditMessageCaptionRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.Caption.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditMessageCaptionRequest) String() string {
	if e == nil {
		return "EditMessageCaptionRequest(nil)"
	}
	type Alias EditMessageCaptionRequest
	return fmt.Sprintf("EditMessageCaptionRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditMessageCaptionRequest) TypeID() uint32 {
	return EditMessageCaptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditMessageCaptionRequest) TypeName() string {
	return "editMessageCaption"
}

// TypeInfo returns info about TL type.
func (e *EditMessageCaptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editMessageCaption",
		ID:   EditMessageCaptionRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditMessageCaptionRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageCaption#44d2f92e as nil")
	}
	b.PutID(EditMessageCaptionRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditMessageCaptionRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageCaption#44d2f92e as nil")
	}
	b.PutLong(e.ChatID)
	b.PutLong(e.MessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field reply_markup: %w", err)
	}
	if err := e.Caption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field caption: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditMessageCaptionRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageCaption#44d2f92e to nil")
	}
	if err := b.ConsumeID(EditMessageCaptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditMessageCaptionRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageCaption#44d2f92e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		if err := e.Caption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field caption: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditMessageCaptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageCaption#44d2f92e as nil")
	}
	b.ObjStart()
	b.PutID("editMessageCaption")
	b.FieldStart("chat_id")
	b.PutLong(e.ChatID)
	b.FieldStart("message_id")
	b.PutLong(e.MessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field reply_markup: %w", err)
	}
	b.FieldStart("caption")
	if err := e.Caption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageCaption#44d2f92e: field caption: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditMessageCaptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageCaption#44d2f92e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editMessageCaption"); err != nil {
				return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field message_id: %w", err)
			}
			e.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "caption":
			if err := e.Caption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editMessageCaption#44d2f92e: field caption: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditMessageCaptionRequest) GetChatID() (value int64) {
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditMessageCaptionRequest) GetMessageID() (value int64) {
	return e.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditMessageCaptionRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	return e.ReplyMarkup
}

// GetCaption returns value of Caption field.
func (e *EditMessageCaptionRequest) GetCaption() (value FormattedText) {
	return e.Caption
}

// EditMessageCaption invokes method editMessageCaption#44d2f92e returning error if any.
func (c *Client) EditMessageCaption(ctx context.Context, request *EditMessageCaptionRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}