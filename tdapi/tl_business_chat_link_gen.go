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

// BusinessChatLink represents TL type `businessChatLink#8e998b83`.
type BusinessChatLink struct {
	// The HTTPS link
	Link string
	// Message draft text that will be added to the input field
	Text FormattedText
	// Link title
	Title string
	// Number of times the link was used
	ViewCount int32
}

// BusinessChatLinkTypeID is TL type id of BusinessChatLink.
const BusinessChatLinkTypeID = 0x8e998b83

// Ensuring interfaces in compile-time for BusinessChatLink.
var (
	_ bin.Encoder     = &BusinessChatLink{}
	_ bin.Decoder     = &BusinessChatLink{}
	_ bin.BareEncoder = &BusinessChatLink{}
	_ bin.BareDecoder = &BusinessChatLink{}
)

func (b *BusinessChatLink) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Link == "") {
		return false
	}
	if !(b.Text.Zero()) {
		return false
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.ViewCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessChatLink) String() string {
	if b == nil {
		return "BusinessChatLink(nil)"
	}
	type Alias BusinessChatLink
	return fmt.Sprintf("BusinessChatLink%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessChatLink) TypeID() uint32 {
	return BusinessChatLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessChatLink) TypeName() string {
	return "businessChatLink"
}

// TypeInfo returns info about TL type.
func (b *BusinessChatLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessChatLink",
		ID:   BusinessChatLinkTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "ViewCount",
			SchemaName: "view_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessChatLink) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLink#8e998b83 as nil")
	}
	buf.PutID(BusinessChatLinkTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessChatLink) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLink#8e998b83 as nil")
	}
	buf.PutString(b.Link)
	if err := b.Text.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessChatLink#8e998b83: field text: %w", err)
	}
	buf.PutString(b.Title)
	buf.PutInt32(b.ViewCount)
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessChatLink) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLink#8e998b83 to nil")
	}
	if err := buf.ConsumeID(BusinessChatLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode businessChatLink#8e998b83: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessChatLink) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLink#8e998b83 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode businessChatLink#8e998b83: field link: %w", err)
		}
		b.Link = value
	}
	{
		if err := b.Text.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessChatLink#8e998b83: field text: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode businessChatLink#8e998b83: field title: %w", err)
		}
		b.Title = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode businessChatLink#8e998b83: field view_count: %w", err)
		}
		b.ViewCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessChatLink) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLink#8e998b83 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessChatLink")
	buf.Comma()
	buf.FieldStart("link")
	buf.PutString(b.Link)
	buf.Comma()
	buf.FieldStart("text")
	if err := b.Text.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessChatLink#8e998b83: field text: %w", err)
	}
	buf.Comma()
	buf.FieldStart("title")
	buf.PutString(b.Title)
	buf.Comma()
	buf.FieldStart("view_count")
	buf.PutInt32(b.ViewCount)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessChatLink) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLink#8e998b83 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessChatLink"); err != nil {
				return fmt.Errorf("unable to decode businessChatLink#8e998b83: %w", err)
			}
		case "link":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode businessChatLink#8e998b83: field link: %w", err)
			}
			b.Link = value
		case "text":
			if err := b.Text.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessChatLink#8e998b83: field text: %w", err)
			}
		case "title":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode businessChatLink#8e998b83: field title: %w", err)
			}
			b.Title = value
		case "view_count":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode businessChatLink#8e998b83: field view_count: %w", err)
			}
			b.ViewCount = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetLink returns value of Link field.
func (b *BusinessChatLink) GetLink() (value string) {
	if b == nil {
		return
	}
	return b.Link
}

// GetText returns value of Text field.
func (b *BusinessChatLink) GetText() (value FormattedText) {
	if b == nil {
		return
	}
	return b.Text
}

// GetTitle returns value of Title field.
func (b *BusinessChatLink) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

// GetViewCount returns value of ViewCount field.
func (b *BusinessChatLink) GetViewCount() (value int32) {
	if b == nil {
		return
	}
	return b.ViewCount
}
