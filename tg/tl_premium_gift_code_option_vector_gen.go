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

// PremiumGiftCodeOptionVector is a box for Vector<PremiumGiftCodeOption>
type PremiumGiftCodeOptionVector struct {
	// Elements of Vector<PremiumGiftCodeOption>
	Elems []PremiumGiftCodeOption
}

// PremiumGiftCodeOptionVectorTypeID is TL type id of PremiumGiftCodeOptionVector.
const PremiumGiftCodeOptionVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for PremiumGiftCodeOptionVector.
var (
	_ bin.Encoder     = &PremiumGiftCodeOptionVector{}
	_ bin.Decoder     = &PremiumGiftCodeOptionVector{}
	_ bin.BareEncoder = &PremiumGiftCodeOptionVector{}
	_ bin.BareDecoder = &PremiumGiftCodeOptionVector{}
)

func (vec *PremiumGiftCodeOptionVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *PremiumGiftCodeOptionVector) String() string {
	if vec == nil {
		return "PremiumGiftCodeOptionVector(nil)"
	}
	type Alias PremiumGiftCodeOptionVector
	return fmt.Sprintf("PremiumGiftCodeOptionVector%+v", Alias(*vec))
}

// FillFrom fills PremiumGiftCodeOptionVector from given interface.
func (vec *PremiumGiftCodeOptionVector) FillFrom(from interface {
	GetElems() (value []PremiumGiftCodeOption)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGiftCodeOptionVector) TypeID() uint32 {
	return PremiumGiftCodeOptionVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGiftCodeOptionVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *PremiumGiftCodeOptionVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   PremiumGiftCodeOptionVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *PremiumGiftCodeOptionVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<PremiumGiftCodeOption> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *PremiumGiftCodeOptionVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<PremiumGiftCodeOption> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<PremiumGiftCodeOption>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *PremiumGiftCodeOptionVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<PremiumGiftCodeOption> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *PremiumGiftCodeOptionVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<PremiumGiftCodeOption> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<PremiumGiftCodeOption>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]PremiumGiftCodeOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PremiumGiftCodeOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<PremiumGiftCodeOption>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *PremiumGiftCodeOptionVector) GetElems() (value []PremiumGiftCodeOption) {
	if vec == nil {
		return
	}
	return vec.Elems
}