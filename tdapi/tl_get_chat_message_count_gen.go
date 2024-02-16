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

// GetChatMessageCountRequest represents TL type `getChatMessageCount#38f78909`.
type GetChatMessageCountRequest struct {
	// Identifier of the chat in which to count messages
	ChatID int64
	// Filter for message content; searchMessagesFilterEmpty is unsupported in this function
	Filter SearchMessagesFilterClass
	// If not 0, only messages in the specified Saved Messages topic will be counted; pass 0
	// to count all messages, or for chats other than Saved Messages
	SavedMessagesTopicID int64
	// Pass true to get the number of messages without sending network requests, or -1 if the
	// number of messages is unknown locally
	ReturnLocal bool
}

// GetChatMessageCountRequestTypeID is TL type id of GetChatMessageCountRequest.
const GetChatMessageCountRequestTypeID = 0x38f78909

// Ensuring interfaces in compile-time for GetChatMessageCountRequest.
var (
	_ bin.Encoder     = &GetChatMessageCountRequest{}
	_ bin.Decoder     = &GetChatMessageCountRequest{}
	_ bin.BareEncoder = &GetChatMessageCountRequest{}
	_ bin.BareDecoder = &GetChatMessageCountRequest{}
)

func (g *GetChatMessageCountRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.SavedMessagesTopicID == 0) {
		return false
	}
	if !(g.ReturnLocal == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatMessageCountRequest) String() string {
	if g == nil {
		return "GetChatMessageCountRequest(nil)"
	}
	type Alias GetChatMessageCountRequest
	return fmt.Sprintf("GetChatMessageCountRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatMessageCountRequest) TypeID() uint32 {
	return GetChatMessageCountRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatMessageCountRequest) TypeName() string {
	return "getChatMessageCount"
}

// TypeInfo returns info about TL type.
func (g *GetChatMessageCountRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatMessageCount",
		ID:   GetChatMessageCountRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "SavedMessagesTopicID",
			SchemaName: "saved_messages_topic_id",
		},
		{
			Name:       "ReturnLocal",
			SchemaName: "return_local",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatMessageCountRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCount#38f78909 as nil")
	}
	b.PutID(GetChatMessageCountRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatMessageCountRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCount#38f78909 as nil")
	}
	b.PutInt53(g.ChatID)
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessageCount#38f78909: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessageCount#38f78909: field filter: %w", err)
	}
	b.PutInt53(g.SavedMessagesTopicID)
	b.PutBool(g.ReturnLocal)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatMessageCountRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCount#38f78909 to nil")
	}
	if err := b.ConsumeID(GetChatMessageCountRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatMessageCount#38f78909: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatMessageCountRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCount#38f78909 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field saved_messages_topic_id: %w", err)
		}
		g.SavedMessagesTopicID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field return_local: %w", err)
		}
		g.ReturnLocal = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatMessageCountRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCount#38f78909 as nil")
	}
	b.ObjStart()
	b.PutID("getChatMessageCount")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessageCount#38f78909: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessageCount#38f78909: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("saved_messages_topic_id")
	b.PutInt53(g.SavedMessagesTopicID)
	b.Comma()
	b.FieldStart("return_local")
	b.PutBool(g.ReturnLocal)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatMessageCountRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCount#38f78909 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatMessageCount"); err != nil {
				return fmt.Errorf("unable to decode getChatMessageCount#38f78909: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field chat_id: %w", err)
			}
			g.ChatID = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field filter: %w", err)
			}
			g.Filter = value
		case "saved_messages_topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field saved_messages_topic_id: %w", err)
			}
			g.SavedMessagesTopicID = value
		case "return_local":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCount#38f78909: field return_local: %w", err)
			}
			g.ReturnLocal = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatMessageCountRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetFilter returns value of Filter field.
func (g *GetChatMessageCountRequest) GetFilter() (value SearchMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetSavedMessagesTopicID returns value of SavedMessagesTopicID field.
func (g *GetChatMessageCountRequest) GetSavedMessagesTopicID() (value int64) {
	if g == nil {
		return
	}
	return g.SavedMessagesTopicID
}

// GetReturnLocal returns value of ReturnLocal field.
func (g *GetChatMessageCountRequest) GetReturnLocal() (value bool) {
	if g == nil {
		return
	}
	return g.ReturnLocal
}

// GetChatMessageCount invokes method getChatMessageCount#38f78909 returning error if any.
func (c *Client) GetChatMessageCount(ctx context.Context, request *GetChatMessageCountRequest) (*Count, error) {
	var result Count

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
