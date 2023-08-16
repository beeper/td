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

// CanSendStoryRequest represents TL type `canSendStory#ede53a66`.
type CanSendStoryRequest struct {
}

// CanSendStoryRequestTypeID is TL type id of CanSendStoryRequest.
const CanSendStoryRequestTypeID = 0xede53a66

// Ensuring interfaces in compile-time for CanSendStoryRequest.
var (
	_ bin.Encoder     = &CanSendStoryRequest{}
	_ bin.Decoder     = &CanSendStoryRequest{}
	_ bin.BareEncoder = &CanSendStoryRequest{}
	_ bin.BareDecoder = &CanSendStoryRequest{}
)

func (c *CanSendStoryRequest) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CanSendStoryRequest) String() string {
	if c == nil {
		return "CanSendStoryRequest(nil)"
	}
	type Alias CanSendStoryRequest
	return fmt.Sprintf("CanSendStoryRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CanSendStoryRequest) TypeID() uint32 {
	return CanSendStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CanSendStoryRequest) TypeName() string {
	return "canSendStory"
}

// TypeInfo returns info about TL type.
func (c *CanSendStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "canSendStory",
		ID:   CanSendStoryRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CanSendStoryRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#ede53a66 as nil")
	}
	b.PutID(CanSendStoryRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CanSendStoryRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#ede53a66 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CanSendStoryRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#ede53a66 to nil")
	}
	if err := b.ConsumeID(CanSendStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode canSendStory#ede53a66: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CanSendStoryRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#ede53a66 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CanSendStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#ede53a66 as nil")
	}
	b.ObjStart()
	b.PutID("canSendStory")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CanSendStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#ede53a66 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("canSendStory"); err != nil {
				return fmt.Errorf("unable to decode canSendStory#ede53a66: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CanSendStory invokes method canSendStory#ede53a66 returning error if any.
func (c *Client) CanSendStory(ctx context.Context) (CanSendStoryResultClass, error) {
	var result CanSendStoryResultBox

	request := &CanSendStoryRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CanSendStoryResult, nil
}