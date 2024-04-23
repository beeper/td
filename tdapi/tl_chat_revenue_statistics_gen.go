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

// ChatRevenueStatistics represents TL type `chatRevenueStatistics#a5504cd4`.
type ChatRevenueStatistics struct {
	// A graph containing amount of revenue in a given hour
	RevenueByHourGraph StatisticalGraphClass
	// A graph containing amount of revenue
	RevenueGraph StatisticalGraphClass
	// Cryptocurrency in which revenue is calculated
	Cryptocurrency string
	// Total amount of the cryptocurrency earned, in the smallest units of the cryptocurrency
	CryptocurrencyTotalAmount int64
	// Amount of the cryptocurrency that isn't withdrawn yet, in the smallest units of the
	// cryptocurrency
	CryptocurrencyBalanceAmount int64
	// Amount of the cryptocurrency available for withdrawal, in the smallest units of the
	// cryptocurrency
	CryptocurrencyAvailableAmount int64
	// Current conversion rate of the cryptocurrency to USD
	UsdRate float64
}

// ChatRevenueStatisticsTypeID is TL type id of ChatRevenueStatistics.
const ChatRevenueStatisticsTypeID = 0xa5504cd4

// Ensuring interfaces in compile-time for ChatRevenueStatistics.
var (
	_ bin.Encoder     = &ChatRevenueStatistics{}
	_ bin.Decoder     = &ChatRevenueStatistics{}
	_ bin.BareEncoder = &ChatRevenueStatistics{}
	_ bin.BareDecoder = &ChatRevenueStatistics{}
)

func (c *ChatRevenueStatistics) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RevenueByHourGraph == nil) {
		return false
	}
	if !(c.RevenueGraph == nil) {
		return false
	}
	if !(c.Cryptocurrency == "") {
		return false
	}
	if !(c.CryptocurrencyTotalAmount == 0) {
		return false
	}
	if !(c.CryptocurrencyBalanceAmount == 0) {
		return false
	}
	if !(c.CryptocurrencyAvailableAmount == 0) {
		return false
	}
	if !(c.UsdRate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueStatistics) String() string {
	if c == nil {
		return "ChatRevenueStatistics(nil)"
	}
	type Alias ChatRevenueStatistics
	return fmt.Sprintf("ChatRevenueStatistics%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueStatistics) TypeID() uint32 {
	return ChatRevenueStatisticsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueStatistics) TypeName() string {
	return "chatRevenueStatistics"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueStatistics) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueStatistics",
		ID:   ChatRevenueStatisticsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RevenueByHourGraph",
			SchemaName: "revenue_by_hour_graph",
		},
		{
			Name:       "RevenueGraph",
			SchemaName: "revenue_graph",
		},
		{
			Name:       "Cryptocurrency",
			SchemaName: "cryptocurrency",
		},
		{
			Name:       "CryptocurrencyTotalAmount",
			SchemaName: "cryptocurrency_total_amount",
		},
		{
			Name:       "CryptocurrencyBalanceAmount",
			SchemaName: "cryptocurrency_balance_amount",
		},
		{
			Name:       "CryptocurrencyAvailableAmount",
			SchemaName: "cryptocurrency_available_amount",
		},
		{
			Name:       "UsdRate",
			SchemaName: "usd_rate",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueStatistics) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueStatistics#a5504cd4 as nil")
	}
	b.PutID(ChatRevenueStatisticsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueStatistics) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueStatistics#a5504cd4 as nil")
	}
	if c.RevenueByHourGraph == nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph is nil")
	}
	if err := c.RevenueByHourGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph: %w", err)
	}
	if c.RevenueGraph == nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_graph is nil")
	}
	if err := c.RevenueGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_graph: %w", err)
	}
	b.PutString(c.Cryptocurrency)
	b.PutLong(c.CryptocurrencyTotalAmount)
	b.PutLong(c.CryptocurrencyBalanceAmount)
	b.PutLong(c.CryptocurrencyAvailableAmount)
	b.PutDouble(c.UsdRate)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueStatistics) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueStatistics#a5504cd4 to nil")
	}
	if err := b.ConsumeID(ChatRevenueStatisticsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueStatistics) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueStatistics#a5504cd4 to nil")
	}
	{
		value, err := DecodeStatisticalGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph: %w", err)
		}
		c.RevenueByHourGraph = value
	}
	{
		value, err := DecodeStatisticalGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field revenue_graph: %w", err)
		}
		c.RevenueGraph = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency: %w", err)
		}
		c.Cryptocurrency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_total_amount: %w", err)
		}
		c.CryptocurrencyTotalAmount = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_balance_amount: %w", err)
		}
		c.CryptocurrencyBalanceAmount = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_available_amount: %w", err)
		}
		c.CryptocurrencyAvailableAmount = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field usd_rate: %w", err)
		}
		c.UsdRate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueStatistics) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueStatistics#a5504cd4 as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueStatistics")
	b.Comma()
	b.FieldStart("revenue_by_hour_graph")
	if c.RevenueByHourGraph == nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph is nil")
	}
	if err := c.RevenueByHourGraph.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph: %w", err)
	}
	b.Comma()
	b.FieldStart("revenue_graph")
	if c.RevenueGraph == nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_graph is nil")
	}
	if err := c.RevenueGraph.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueStatistics#a5504cd4: field revenue_graph: %w", err)
	}
	b.Comma()
	b.FieldStart("cryptocurrency")
	b.PutString(c.Cryptocurrency)
	b.Comma()
	b.FieldStart("cryptocurrency_total_amount")
	b.PutLong(c.CryptocurrencyTotalAmount)
	b.Comma()
	b.FieldStart("cryptocurrency_balance_amount")
	b.PutLong(c.CryptocurrencyBalanceAmount)
	b.Comma()
	b.FieldStart("cryptocurrency_available_amount")
	b.PutLong(c.CryptocurrencyAvailableAmount)
	b.Comma()
	b.FieldStart("usd_rate")
	b.PutDouble(c.UsdRate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueStatistics) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueStatistics#a5504cd4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueStatistics"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: %w", err)
			}
		case "revenue_by_hour_graph":
			value, err := DecodeTDLibJSONStatisticalGraph(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field revenue_by_hour_graph: %w", err)
			}
			c.RevenueByHourGraph = value
		case "revenue_graph":
			value, err := DecodeTDLibJSONStatisticalGraph(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field revenue_graph: %w", err)
			}
			c.RevenueGraph = value
		case "cryptocurrency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency: %w", err)
			}
			c.Cryptocurrency = value
		case "cryptocurrency_total_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_total_amount: %w", err)
			}
			c.CryptocurrencyTotalAmount = value
		case "cryptocurrency_balance_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_balance_amount: %w", err)
			}
			c.CryptocurrencyBalanceAmount = value
		case "cryptocurrency_available_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field cryptocurrency_available_amount: %w", err)
			}
			c.CryptocurrencyAvailableAmount = value
		case "usd_rate":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueStatistics#a5504cd4: field usd_rate: %w", err)
			}
			c.UsdRate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRevenueByHourGraph returns value of RevenueByHourGraph field.
func (c *ChatRevenueStatistics) GetRevenueByHourGraph() (value StatisticalGraphClass) {
	if c == nil {
		return
	}
	return c.RevenueByHourGraph
}

// GetRevenueGraph returns value of RevenueGraph field.
func (c *ChatRevenueStatistics) GetRevenueGraph() (value StatisticalGraphClass) {
	if c == nil {
		return
	}
	return c.RevenueGraph
}

// GetCryptocurrency returns value of Cryptocurrency field.
func (c *ChatRevenueStatistics) GetCryptocurrency() (value string) {
	if c == nil {
		return
	}
	return c.Cryptocurrency
}

// GetCryptocurrencyTotalAmount returns value of CryptocurrencyTotalAmount field.
func (c *ChatRevenueStatistics) GetCryptocurrencyTotalAmount() (value int64) {
	if c == nil {
		return
	}
	return c.CryptocurrencyTotalAmount
}

// GetCryptocurrencyBalanceAmount returns value of CryptocurrencyBalanceAmount field.
func (c *ChatRevenueStatistics) GetCryptocurrencyBalanceAmount() (value int64) {
	if c == nil {
		return
	}
	return c.CryptocurrencyBalanceAmount
}

// GetCryptocurrencyAvailableAmount returns value of CryptocurrencyAvailableAmount field.
func (c *ChatRevenueStatistics) GetCryptocurrencyAvailableAmount() (value int64) {
	if c == nil {
		return
	}
	return c.CryptocurrencyAvailableAmount
}

// GetUsdRate returns value of UsdRate field.
func (c *ChatRevenueStatistics) GetUsdRate() (value float64) {
	if c == nil {
		return
	}
	return c.UsdRate
}
