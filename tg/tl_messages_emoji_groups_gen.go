// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesEmojiGroupsNotModified represents TL type `messages.emojiGroupsNotModified#6fb4ad87`.
// The list of emoji categories¹ hasn't changed.
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji#emoji-categories
//
// See https://core.telegram.org/constructor/messages.emojiGroupsNotModified for reference.
type MessagesEmojiGroupsNotModified struct {
}

// MessagesEmojiGroupsNotModifiedTypeID is TL type id of MessagesEmojiGroupsNotModified.
const MessagesEmojiGroupsNotModifiedTypeID = 0x6fb4ad87

// construct implements constructor of MessagesEmojiGroupsClass.
func (e MessagesEmojiGroupsNotModified) construct() MessagesEmojiGroupsClass { return &e }

// Ensuring interfaces in compile-time for MessagesEmojiGroupsNotModified.
var (
	_ bin.Encoder     = &MessagesEmojiGroupsNotModified{}
	_ bin.Decoder     = &MessagesEmojiGroupsNotModified{}
	_ bin.BareEncoder = &MessagesEmojiGroupsNotModified{}
	_ bin.BareDecoder = &MessagesEmojiGroupsNotModified{}

	_ MessagesEmojiGroupsClass = &MessagesEmojiGroupsNotModified{}
)

func (e *MessagesEmojiGroupsNotModified) Zero() bool {
	if e == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEmojiGroupsNotModified) String() string {
	if e == nil {
		return "MessagesEmojiGroupsNotModified(nil)"
	}
	type Alias MessagesEmojiGroupsNotModified
	return fmt.Sprintf("MessagesEmojiGroupsNotModified%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEmojiGroupsNotModified) TypeID() uint32 {
	return MessagesEmojiGroupsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEmojiGroupsNotModified) TypeName() string {
	return "messages.emojiGroupsNotModified"
}

// TypeInfo returns info about TL type.
func (e *MessagesEmojiGroupsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.emojiGroupsNotModified",
		ID:   MessagesEmojiGroupsNotModifiedTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEmojiGroupsNotModified) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.emojiGroupsNotModified#6fb4ad87 as nil")
	}
	b.PutID(MessagesEmojiGroupsNotModifiedTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEmojiGroupsNotModified) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.emojiGroupsNotModified#6fb4ad87 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEmojiGroupsNotModified) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.emojiGroupsNotModified#6fb4ad87 to nil")
	}
	if err := b.ConsumeID(MessagesEmojiGroupsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.emojiGroupsNotModified#6fb4ad87: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEmojiGroupsNotModified) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.emojiGroupsNotModified#6fb4ad87 to nil")
	}
	return nil
}

// MessagesEmojiGroups represents TL type `messages.emojiGroups#881fb94b`.
// Represents a list of emoji categories¹.
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji#emoji-categories
//
// See https://core.telegram.org/constructor/messages.emojiGroups for reference.
type MessagesEmojiGroups struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// A list of emoji categories¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/custom-emoji#emoji-categories
	Groups []EmojiGroup
}

// MessagesEmojiGroupsTypeID is TL type id of MessagesEmojiGroups.
const MessagesEmojiGroupsTypeID = 0x881fb94b

// construct implements constructor of MessagesEmojiGroupsClass.
func (e MessagesEmojiGroups) construct() MessagesEmojiGroupsClass { return &e }

// Ensuring interfaces in compile-time for MessagesEmojiGroups.
var (
	_ bin.Encoder     = &MessagesEmojiGroups{}
	_ bin.Decoder     = &MessagesEmojiGroups{}
	_ bin.BareEncoder = &MessagesEmojiGroups{}
	_ bin.BareDecoder = &MessagesEmojiGroups{}

	_ MessagesEmojiGroupsClass = &MessagesEmojiGroups{}
)

func (e *MessagesEmojiGroups) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Hash == 0) {
		return false
	}
	if !(e.Groups == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEmojiGroups) String() string {
	if e == nil {
		return "MessagesEmojiGroups(nil)"
	}
	type Alias MessagesEmojiGroups
	return fmt.Sprintf("MessagesEmojiGroups%+v", Alias(*e))
}

// FillFrom fills MessagesEmojiGroups from given interface.
func (e *MessagesEmojiGroups) FillFrom(from interface {
	GetHash() (value int)
	GetGroups() (value []EmojiGroup)
}) {
	e.Hash = from.GetHash()
	e.Groups = from.GetGroups()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEmojiGroups) TypeID() uint32 {
	return MessagesEmojiGroupsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEmojiGroups) TypeName() string {
	return "messages.emojiGroups"
}

// TypeInfo returns info about TL type.
func (e *MessagesEmojiGroups) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.emojiGroups",
		ID:   MessagesEmojiGroupsTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Groups",
			SchemaName: "groups",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEmojiGroups) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.emojiGroups#881fb94b as nil")
	}
	b.PutID(MessagesEmojiGroupsTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEmojiGroups) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.emojiGroups#881fb94b as nil")
	}
	b.PutInt(e.Hash)
	b.PutVectorHeader(len(e.Groups))
	for idx, v := range e.Groups {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.emojiGroups#881fb94b: field groups element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEmojiGroups) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.emojiGroups#881fb94b to nil")
	}
	if err := b.ConsumeID(MessagesEmojiGroupsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.emojiGroups#881fb94b: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEmojiGroups) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.emojiGroups#881fb94b to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.emojiGroups#881fb94b: field hash: %w", err)
		}
		e.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.emojiGroups#881fb94b: field groups: %w", err)
		}

		if headerLen > 0 {
			e.Groups = make([]EmojiGroup, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value EmojiGroup
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.emojiGroups#881fb94b: field groups: %w", err)
			}
			e.Groups = append(e.Groups, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (e *MessagesEmojiGroups) GetHash() (value int) {
	if e == nil {
		return
	}
	return e.Hash
}

// GetGroups returns value of Groups field.
func (e *MessagesEmojiGroups) GetGroups() (value []EmojiGroup) {
	if e == nil {
		return
	}
	return e.Groups
}

// MessagesEmojiGroupsClassName is schema name of MessagesEmojiGroupsClass.
const MessagesEmojiGroupsClassName = "messages.EmojiGroups"

// MessagesEmojiGroupsClass represents messages.EmojiGroups generic type.
//
// See https://core.telegram.org/type/messages.EmojiGroups for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesEmojiGroups(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesEmojiGroupsNotModified: // messages.emojiGroupsNotModified#6fb4ad87
//	case *tg.MessagesEmojiGroups: // messages.emojiGroups#881fb94b
//	default: panic(v)
//	}
type MessagesEmojiGroupsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesEmojiGroupsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map MessagesEmojiGroupsClass to MessagesEmojiGroups.
	AsModified() (*MessagesEmojiGroups, bool)
}

// AsModified tries to map MessagesEmojiGroupsNotModified to MessagesEmojiGroups.
func (e *MessagesEmojiGroupsNotModified) AsModified() (*MessagesEmojiGroups, bool) {
	return nil, false
}

// AsModified tries to map MessagesEmojiGroups to MessagesEmojiGroups.
func (e *MessagesEmojiGroups) AsModified() (*MessagesEmojiGroups, bool) {
	return e, true
}

// DecodeMessagesEmojiGroups implements binary de-serialization for MessagesEmojiGroupsClass.
func DecodeMessagesEmojiGroups(buf *bin.Buffer) (MessagesEmojiGroupsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesEmojiGroupsNotModifiedTypeID:
		// Decoding messages.emojiGroupsNotModified#6fb4ad87.
		v := MessagesEmojiGroupsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesEmojiGroupsClass: %w", err)
		}
		return &v, nil
	case MessagesEmojiGroupsTypeID:
		// Decoding messages.emojiGroups#881fb94b.
		v := MessagesEmojiGroups{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesEmojiGroupsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesEmojiGroupsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesEmojiGroups boxes the MessagesEmojiGroupsClass providing a helper.
type MessagesEmojiGroupsBox struct {
	EmojiGroups MessagesEmojiGroupsClass
}

// Decode implements bin.Decoder for MessagesEmojiGroupsBox.
func (b *MessagesEmojiGroupsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesEmojiGroupsBox to nil")
	}
	v, err := DecodeMessagesEmojiGroups(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiGroups = v
	return nil
}

// Encode implements bin.Encode for MessagesEmojiGroupsBox.
func (b *MessagesEmojiGroupsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiGroups == nil {
		return fmt.Errorf("unable to encode MessagesEmojiGroupsClass as nil")
	}
	return b.EmojiGroups.Encode(buf)
}