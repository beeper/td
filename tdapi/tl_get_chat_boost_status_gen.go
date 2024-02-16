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

// GetChatBoostStatusRequest represents TL type `getChatBoostStatus#cfac8acf`.
type GetChatBoostStatusRequest struct {
	// Identifier of the chat
	ChatID int64
}

// GetChatBoostStatusRequestTypeID is TL type id of GetChatBoostStatusRequest.
const GetChatBoostStatusRequestTypeID = 0xcfac8acf

// Ensuring interfaces in compile-time for GetChatBoostStatusRequest.
var (
	_ bin.Encoder     = &GetChatBoostStatusRequest{}
	_ bin.Decoder     = &GetChatBoostStatusRequest{}
	_ bin.BareEncoder = &GetChatBoostStatusRequest{}
	_ bin.BareDecoder = &GetChatBoostStatusRequest{}
)

func (g *GetChatBoostStatusRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatBoostStatusRequest) String() string {
	if g == nil {
		return "GetChatBoostStatusRequest(nil)"
	}
	type Alias GetChatBoostStatusRequest
	return fmt.Sprintf("GetChatBoostStatusRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatBoostStatusRequest) TypeID() uint32 {
	return GetChatBoostStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatBoostStatusRequest) TypeName() string {
	return "getChatBoostStatus"
}

// TypeInfo returns info about TL type.
func (g *GetChatBoostStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatBoostStatus",
		ID:   GetChatBoostStatusRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatBoostStatusRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostStatus#cfac8acf as nil")
	}
	b.PutID(GetChatBoostStatusRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatBoostStatusRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostStatus#cfac8acf as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatBoostStatusRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostStatus#cfac8acf to nil")
	}
	if err := b.ConsumeID(GetChatBoostStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatBoostStatus#cfac8acf: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatBoostStatusRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostStatus#cfac8acf to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoostStatus#cfac8acf: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatBoostStatusRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostStatus#cfac8acf as nil")
	}
	b.ObjStart()
	b.PutID("getChatBoostStatus")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatBoostStatusRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostStatus#cfac8acf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatBoostStatus"); err != nil {
				return fmt.Errorf("unable to decode getChatBoostStatus#cfac8acf: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoostStatus#cfac8acf: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatBoostStatusRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatBoostStatus invokes method getChatBoostStatus#cfac8acf returning error if any.
func (c *Client) GetChatBoostStatus(ctx context.Context, chatid int64) (*ChatBoostStatus, error) {
	var result ChatBoostStatus

	request := &GetChatBoostStatusRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
