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

// ReorderBotMediaPreviewsRequest represents TL type `reorderBotMediaPreviews#89ea0cc6`.
type ReorderBotMediaPreviewsRequest struct {
	// Identifier of the target bot. The bot must be owned and must have the main Web App
	BotUserID int64
	// Language code of the media previews to reorder
	LanguageCode string
	// File identifiers of the media in the new order
	FileIDs []int32
}

// ReorderBotMediaPreviewsRequestTypeID is TL type id of ReorderBotMediaPreviewsRequest.
const ReorderBotMediaPreviewsRequestTypeID = 0x89ea0cc6

// Ensuring interfaces in compile-time for ReorderBotMediaPreviewsRequest.
var (
	_ bin.Encoder     = &ReorderBotMediaPreviewsRequest{}
	_ bin.Decoder     = &ReorderBotMediaPreviewsRequest{}
	_ bin.BareEncoder = &ReorderBotMediaPreviewsRequest{}
	_ bin.BareDecoder = &ReorderBotMediaPreviewsRequest{}
)

func (r *ReorderBotMediaPreviewsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.BotUserID == 0) {
		return false
	}
	if !(r.LanguageCode == "") {
		return false
	}
	if !(r.FileIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReorderBotMediaPreviewsRequest) String() string {
	if r == nil {
		return "ReorderBotMediaPreviewsRequest(nil)"
	}
	type Alias ReorderBotMediaPreviewsRequest
	return fmt.Sprintf("ReorderBotMediaPreviewsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReorderBotMediaPreviewsRequest) TypeID() uint32 {
	return ReorderBotMediaPreviewsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReorderBotMediaPreviewsRequest) TypeName() string {
	return "reorderBotMediaPreviews"
}

// TypeInfo returns info about TL type.
func (r *ReorderBotMediaPreviewsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reorderBotMediaPreviews",
		ID:   ReorderBotMediaPreviewsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "FileIDs",
			SchemaName: "file_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReorderBotMediaPreviewsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderBotMediaPreviews#89ea0cc6 as nil")
	}
	b.PutID(ReorderBotMediaPreviewsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReorderBotMediaPreviewsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderBotMediaPreviews#89ea0cc6 as nil")
	}
	b.PutInt53(r.BotUserID)
	b.PutString(r.LanguageCode)
	b.PutInt(len(r.FileIDs))
	for _, v := range r.FileIDs {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReorderBotMediaPreviewsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderBotMediaPreviews#89ea0cc6 to nil")
	}
	if err := b.ConsumeID(ReorderBotMediaPreviewsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReorderBotMediaPreviewsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderBotMediaPreviews#89ea0cc6 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field bot_user_id: %w", err)
		}
		r.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field language_code: %w", err)
		}
		r.LanguageCode = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field file_ids: %w", err)
		}

		if headerLen > 0 {
			r.FileIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field file_ids: %w", err)
			}
			r.FileIDs = append(r.FileIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReorderBotMediaPreviewsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderBotMediaPreviews#89ea0cc6 as nil")
	}
	b.ObjStart()
	b.PutID("reorderBotMediaPreviews")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(r.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(r.LanguageCode)
	b.Comma()
	b.FieldStart("file_ids")
	b.ArrStart()
	for _, v := range r.FileIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReorderBotMediaPreviewsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderBotMediaPreviews#89ea0cc6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reorderBotMediaPreviews"); err != nil {
				return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field bot_user_id: %w", err)
			}
			r.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field language_code: %w", err)
			}
			r.LanguageCode = value
		case "file_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field file_ids: %w", err)
				}
				r.FileIDs = append(r.FileIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reorderBotMediaPreviews#89ea0cc6: field file_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (r *ReorderBotMediaPreviewsRequest) GetBotUserID() (value int64) {
	if r == nil {
		return
	}
	return r.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (r *ReorderBotMediaPreviewsRequest) GetLanguageCode() (value string) {
	if r == nil {
		return
	}
	return r.LanguageCode
}

// GetFileIDs returns value of FileIDs field.
func (r *ReorderBotMediaPreviewsRequest) GetFileIDs() (value []int32) {
	if r == nil {
		return
	}
	return r.FileIDs
}

// ReorderBotMediaPreviews invokes method reorderBotMediaPreviews#89ea0cc6 returning error if any.
func (c *Client) ReorderBotMediaPreviews(ctx context.Context, request *ReorderBotMediaPreviewsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}