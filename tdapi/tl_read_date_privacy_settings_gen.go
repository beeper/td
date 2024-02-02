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

// ReadDatePrivacySettings represents TL type `readDatePrivacySettings#62a2e628`.
type ReadDatePrivacySettings struct {
	// True, if message read date is shown to other users in private chats. If false and the
	// current user isn't a Telegram Premium user, then they will not be able to see other's
	// message read date.
	ShowReadDate bool
}

// ReadDatePrivacySettingsTypeID is TL type id of ReadDatePrivacySettings.
const ReadDatePrivacySettingsTypeID = 0x62a2e628

// Ensuring interfaces in compile-time for ReadDatePrivacySettings.
var (
	_ bin.Encoder     = &ReadDatePrivacySettings{}
	_ bin.Decoder     = &ReadDatePrivacySettings{}
	_ bin.BareEncoder = &ReadDatePrivacySettings{}
	_ bin.BareDecoder = &ReadDatePrivacySettings{}
)

func (r *ReadDatePrivacySettings) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ShowReadDate == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReadDatePrivacySettings) String() string {
	if r == nil {
		return "ReadDatePrivacySettings(nil)"
	}
	type Alias ReadDatePrivacySettings
	return fmt.Sprintf("ReadDatePrivacySettings%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReadDatePrivacySettings) TypeID() uint32 {
	return ReadDatePrivacySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ReadDatePrivacySettings) TypeName() string {
	return "readDatePrivacySettings"
}

// TypeInfo returns info about TL type.
func (r *ReadDatePrivacySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "readDatePrivacySettings",
		ID:   ReadDatePrivacySettingsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShowReadDate",
			SchemaName: "show_read_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReadDatePrivacySettings) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readDatePrivacySettings#62a2e628 as nil")
	}
	b.PutID(ReadDatePrivacySettingsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReadDatePrivacySettings) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readDatePrivacySettings#62a2e628 as nil")
	}
	b.PutBool(r.ShowReadDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReadDatePrivacySettings) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readDatePrivacySettings#62a2e628 to nil")
	}
	if err := b.ConsumeID(ReadDatePrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode readDatePrivacySettings#62a2e628: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReadDatePrivacySettings) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readDatePrivacySettings#62a2e628 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode readDatePrivacySettings#62a2e628: field show_read_date: %w", err)
		}
		r.ShowReadDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReadDatePrivacySettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode readDatePrivacySettings#62a2e628 as nil")
	}
	b.ObjStart()
	b.PutID("readDatePrivacySettings")
	b.Comma()
	b.FieldStart("show_read_date")
	b.PutBool(r.ShowReadDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReadDatePrivacySettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode readDatePrivacySettings#62a2e628 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("readDatePrivacySettings"); err != nil {
				return fmt.Errorf("unable to decode readDatePrivacySettings#62a2e628: %w", err)
			}
		case "show_read_date":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode readDatePrivacySettings#62a2e628: field show_read_date: %w", err)
			}
			r.ShowReadDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetShowReadDate returns value of ShowReadDate field.
func (r *ReadDatePrivacySettings) GetShowReadDate() (value bool) {
	if r == nil {
		return
	}
	return r.ShowReadDate
}