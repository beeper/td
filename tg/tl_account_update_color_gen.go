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

// AccountUpdateColorRequest represents TL type `account.updateColor#a001cc43`.
//
// See https://core.telegram.org/method/account.updateColor for reference.
type AccountUpdateColorRequest struct {
	// Flags field of AccountUpdateColorRequest.
	Flags bin.Fields
	// Color field of AccountUpdateColorRequest.
	Color int
	// BackgroundEmojiID field of AccountUpdateColorRequest.
	//
	// Use SetBackgroundEmojiID and GetBackgroundEmojiID helpers.
	BackgroundEmojiID int64
}

// AccountUpdateColorRequestTypeID is TL type id of AccountUpdateColorRequest.
const AccountUpdateColorRequestTypeID = 0xa001cc43

// Ensuring interfaces in compile-time for AccountUpdateColorRequest.
var (
	_ bin.Encoder     = &AccountUpdateColorRequest{}
	_ bin.Decoder     = &AccountUpdateColorRequest{}
	_ bin.BareEncoder = &AccountUpdateColorRequest{}
	_ bin.BareDecoder = &AccountUpdateColorRequest{}
)

func (u *AccountUpdateColorRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Color == 0) {
		return false
	}
	if !(u.BackgroundEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateColorRequest) String() string {
	if u == nil {
		return "AccountUpdateColorRequest(nil)"
	}
	type Alias AccountUpdateColorRequest
	return fmt.Sprintf("AccountUpdateColorRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateColorRequest from given interface.
func (u *AccountUpdateColorRequest) FillFrom(from interface {
	GetColor() (value int)
	GetBackgroundEmojiID() (value int64, ok bool)
}) {
	u.Color = from.GetColor()
	if val, ok := from.GetBackgroundEmojiID(); ok {
		u.BackgroundEmojiID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateColorRequest) TypeID() uint32 {
	return AccountUpdateColorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateColorRequest) TypeName() string {
	return "account.updateColor"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateColorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateColor",
		ID:   AccountUpdateColorRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Color",
			SchemaName: "color",
		},
		{
			Name:       "BackgroundEmojiID",
			SchemaName: "background_emoji_id",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateColorRequest) SetFlags() {
	if !(u.BackgroundEmojiID == 0) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateColorRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateColor#a001cc43 as nil")
	}
	b.PutID(AccountUpdateColorRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateColorRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateColor#a001cc43 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateColor#a001cc43: field flags: %w", err)
	}
	b.PutInt(u.Color)
	if u.Flags.Has(0) {
		b.PutLong(u.BackgroundEmojiID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateColorRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateColor#a001cc43 to nil")
	}
	if err := b.ConsumeID(AccountUpdateColorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateColor#a001cc43: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateColorRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateColor#a001cc43 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateColor#a001cc43: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateColor#a001cc43: field color: %w", err)
		}
		u.Color = value
	}
	if u.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateColor#a001cc43: field background_emoji_id: %w", err)
		}
		u.BackgroundEmojiID = value
	}
	return nil
}

// GetColor returns value of Color field.
func (u *AccountUpdateColorRequest) GetColor() (value int) {
	if u == nil {
		return
	}
	return u.Color
}

// SetBackgroundEmojiID sets value of BackgroundEmojiID conditional field.
func (u *AccountUpdateColorRequest) SetBackgroundEmojiID(value int64) {
	u.Flags.Set(0)
	u.BackgroundEmojiID = value
}

// GetBackgroundEmojiID returns value of BackgroundEmojiID conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateColorRequest) GetBackgroundEmojiID() (value int64, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.BackgroundEmojiID, true
}

// AccountUpdateColor invokes method account.updateColor#a001cc43 returning error if any.
//
// See https://core.telegram.org/method/account.updateColor for reference.
func (c *Client) AccountUpdateColor(ctx context.Context, request *AccountUpdateColorRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
