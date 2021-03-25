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
)

// MessagesSearchCounter represents TL type `messages.searchCounter#e844ebff`.
// Indicates how many results would be found by a messages.search¹ call with the same
// parameters
//
// Links:
//  1) https://core.telegram.org/method/messages.search
//
// See https://core.telegram.org/constructor/messages.searchCounter for reference.
type MessagesSearchCounter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, the results may be inexact
	Inexact bool
	// Provided message filter
	Filter MessagesFilterClass
	// Number of results that were found server-side
	Count int
}

// MessagesSearchCounterTypeID is TL type id of MessagesSearchCounter.
const MessagesSearchCounterTypeID = 0xe844ebff

func (s *MessagesSearchCounter) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Inexact == false) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchCounter) String() string {
	if s == nil {
		return "MessagesSearchCounter(nil)"
	}
	type Alias MessagesSearchCounter
	return fmt.Sprintf("MessagesSearchCounter%+v", Alias(*s))
}

// FillFrom fills MessagesSearchCounter from given interface.
func (s *MessagesSearchCounter) FillFrom(from interface {
	GetInexact() (value bool)
	GetFilter() (value MessagesFilterClass)
	GetCount() (value int)
}) {
	s.Inexact = from.GetInexact()
	s.Filter = from.GetFilter()
	s.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchCounter) TypeID() uint32 {
	return MessagesSearchCounterTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchCounter) TypeName() string {
	return "messages.searchCounter"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchCounter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchCounter",
		ID:   MessagesSearchCounterTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inexact",
			SchemaName: "inexact",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSearchCounter) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchCounter#e844ebff as nil")
	}
	b.PutID(MessagesSearchCounterTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchCounter) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchCounter#e844ebff as nil")
	}
	if !(s.Inexact == false) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchCounter#e844ebff: field flags: %w", err)
	}
	if s.Filter == nil {
		return fmt.Errorf("unable to encode messages.searchCounter#e844ebff: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchCounter#e844ebff: field filter: %w", err)
	}
	b.PutInt(s.Count)
	return nil
}

// SetInexact sets value of Inexact conditional field.
func (s *MessagesSearchCounter) SetInexact(value bool) {
	if value {
		s.Flags.Set(1)
		s.Inexact = true
	} else {
		s.Flags.Unset(1)
		s.Inexact = false
	}
}

// GetInexact returns value of Inexact conditional field.
func (s *MessagesSearchCounter) GetInexact() (value bool) {
	return s.Flags.Has(1)
}

// GetFilter returns value of Filter field.
func (s *MessagesSearchCounter) GetFilter() (value MessagesFilterClass) {
	return s.Filter
}

// GetCount returns value of Count field.
func (s *MessagesSearchCounter) GetCount() (value int) {
	return s.Count
}

// Decode implements bin.Decoder.
func (s *MessagesSearchCounter) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchCounter#e844ebff to nil")
	}
	if err := b.ConsumeID(MessagesSearchCounterTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchCounter#e844ebff: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchCounter) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchCounter#e844ebff to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.searchCounter#e844ebff: field flags: %w", err)
		}
	}
	s.Inexact = s.Flags.Has(1)
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchCounter#e844ebff: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchCounter#e844ebff: field count: %w", err)
		}
		s.Count = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSearchCounter.
var (
	_ bin.Encoder     = &MessagesSearchCounter{}
	_ bin.Decoder     = &MessagesSearchCounter{}
	_ bin.BareEncoder = &MessagesSearchCounter{}
	_ bin.BareDecoder = &MessagesSearchCounter{}
)
