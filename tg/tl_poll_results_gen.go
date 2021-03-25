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

// PollResults represents TL type `pollResults#badcc1a3`.
// Results of poll
//
// See https://core.telegram.org/constructor/pollResults for reference.
type PollResults struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Similar to min¹ objects, used for poll constructors that are the same for all users
	// so they don't have option chosen by the current user (you can use messages
	// getPollResults² to get the full poll results).
	//
	// Links:
	//  1) https://core.telegram.org/api/min
	//  2) https://core.telegram.org/method/messages.getPollResults
	Min bool
	// Poll results
	//
	// Use SetResults and GetResults helpers.
	Results []PollAnswerVoters
	// Total number of people that voted in the poll
	//
	// Use SetTotalVoters and GetTotalVoters helpers.
	TotalVoters int
	// IDs of the last users that recently voted in the poll
	//
	// Use SetRecentVoters and GetRecentVoters helpers.
	RecentVoters []int
	// Explanation of quiz solution
	//
	// Use SetSolution and GetSolution helpers.
	Solution string
	// Message entities for styled text in quiz solution¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetSolutionEntities and GetSolutionEntities helpers.
	SolutionEntities []MessageEntityClass
}

// PollResultsTypeID is TL type id of PollResults.
const PollResultsTypeID = 0xbadcc1a3

func (p *PollResults) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Min == false) {
		return false
	}
	if !(p.Results == nil) {
		return false
	}
	if !(p.TotalVoters == 0) {
		return false
	}
	if !(p.RecentVoters == nil) {
		return false
	}
	if !(p.Solution == "") {
		return false
	}
	if !(p.SolutionEntities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollResults) String() string {
	if p == nil {
		return "PollResults(nil)"
	}
	type Alias PollResults
	return fmt.Sprintf("PollResults%+v", Alias(*p))
}

// FillFrom fills PollResults from given interface.
func (p *PollResults) FillFrom(from interface {
	GetMin() (value bool)
	GetResults() (value []PollAnswerVoters, ok bool)
	GetTotalVoters() (value int, ok bool)
	GetRecentVoters() (value []int, ok bool)
	GetSolution() (value string, ok bool)
	GetSolutionEntities() (value []MessageEntityClass, ok bool)
}) {
	p.Min = from.GetMin()
	if val, ok := from.GetResults(); ok {
		p.Results = val
	}

	if val, ok := from.GetTotalVoters(); ok {
		p.TotalVoters = val
	}

	if val, ok := from.GetRecentVoters(); ok {
		p.RecentVoters = val
	}

	if val, ok := from.GetSolution(); ok {
		p.Solution = val
	}

	if val, ok := from.GetSolutionEntities(); ok {
		p.SolutionEntities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollResults) TypeID() uint32 {
	return PollResultsTypeID
}

// TypeName returns name of type in TL schema.
func (*PollResults) TypeName() string {
	return "pollResults"
}

// TypeInfo returns info about TL type.
func (p *PollResults) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollResults",
		ID:   PollResultsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Min",
			SchemaName: "min",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Results",
			SchemaName: "results",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "TotalVoters",
			SchemaName: "total_voters",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "RecentVoters",
			SchemaName: "recent_voters",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "Solution",
			SchemaName: "solution",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "SolutionEntities",
			SchemaName: "solution_entities",
			Null:       !p.Flags.Has(4),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollResults) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollResults#badcc1a3 as nil")
	}
	b.PutID(PollResultsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollResults) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollResults#badcc1a3 as nil")
	}
	if !(p.Min == false) {
		p.Flags.Set(0)
	}
	if !(p.Results == nil) {
		p.Flags.Set(1)
	}
	if !(p.TotalVoters == 0) {
		p.Flags.Set(2)
	}
	if !(p.RecentVoters == nil) {
		p.Flags.Set(3)
	}
	if !(p.Solution == "") {
		p.Flags.Set(4)
	}
	if !(p.SolutionEntities == nil) {
		p.Flags.Set(4)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pollResults#badcc1a3: field flags: %w", err)
	}
	if p.Flags.Has(1) {
		b.PutVectorHeader(len(p.Results))
		for idx, v := range p.Results {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode pollResults#badcc1a3: field results element with index %d: %w", idx, err)
			}
		}
	}
	if p.Flags.Has(2) {
		b.PutInt(p.TotalVoters)
	}
	if p.Flags.Has(3) {
		b.PutVectorHeader(len(p.RecentVoters))
		for _, v := range p.RecentVoters {
			b.PutInt(v)
		}
	}
	if p.Flags.Has(4) {
		b.PutString(p.Solution)
	}
	if p.Flags.Has(4) {
		b.PutVectorHeader(len(p.SolutionEntities))
		for idx, v := range p.SolutionEntities {
			if v == nil {
				return fmt.Errorf("unable to encode pollResults#badcc1a3: field solution_entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode pollResults#badcc1a3: field solution_entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetMin sets value of Min conditional field.
func (p *PollResults) SetMin(value bool) {
	if value {
		p.Flags.Set(0)
		p.Min = true
	} else {
		p.Flags.Unset(0)
		p.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (p *PollResults) GetMin() (value bool) {
	return p.Flags.Has(0)
}

// SetResults sets value of Results conditional field.
func (p *PollResults) SetResults(value []PollAnswerVoters) {
	p.Flags.Set(1)
	p.Results = value
}

// GetResults returns value of Results conditional field and
// boolean which is true if field was set.
func (p *PollResults) GetResults() (value []PollAnswerVoters, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Results, true
}

// SetTotalVoters sets value of TotalVoters conditional field.
func (p *PollResults) SetTotalVoters(value int) {
	p.Flags.Set(2)
	p.TotalVoters = value
}

// GetTotalVoters returns value of TotalVoters conditional field and
// boolean which is true if field was set.
func (p *PollResults) GetTotalVoters() (value int, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.TotalVoters, true
}

// SetRecentVoters sets value of RecentVoters conditional field.
func (p *PollResults) SetRecentVoters(value []int) {
	p.Flags.Set(3)
	p.RecentVoters = value
}

// GetRecentVoters returns value of RecentVoters conditional field and
// boolean which is true if field was set.
func (p *PollResults) GetRecentVoters() (value []int, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.RecentVoters, true
}

// SetSolution sets value of Solution conditional field.
func (p *PollResults) SetSolution(value string) {
	p.Flags.Set(4)
	p.Solution = value
}

// GetSolution returns value of Solution conditional field and
// boolean which is true if field was set.
func (p *PollResults) GetSolution() (value string, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.Solution, true
}

// SetSolutionEntities sets value of SolutionEntities conditional field.
func (p *PollResults) SetSolutionEntities(value []MessageEntityClass) {
	p.Flags.Set(4)
	p.SolutionEntities = value
}

// GetSolutionEntities returns value of SolutionEntities conditional field and
// boolean which is true if field was set.
func (p *PollResults) GetSolutionEntities() (value []MessageEntityClass, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.SolutionEntities, true
}

// MapSolutionEntities returns field SolutionEntities wrapped in MessageEntityClassArray helper.
func (p *PollResults) MapSolutionEntities() (value MessageEntityClassArray, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return MessageEntityClassArray(p.SolutionEntities), true
}

// Decode implements bin.Decoder.
func (p *PollResults) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollResults#badcc1a3 to nil")
	}
	if err := b.ConsumeID(PollResultsTypeID); err != nil {
		return fmt.Errorf("unable to decode pollResults#badcc1a3: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollResults) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollResults#badcc1a3 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field flags: %w", err)
		}
	}
	p.Min = p.Flags.Has(0)
	if p.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field results: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PollAnswerVoters
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode pollResults#badcc1a3: field results: %w", err)
			}
			p.Results = append(p.Results, value)
		}
	}
	if p.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field total_voters: %w", err)
		}
		p.TotalVoters = value
	}
	if p.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field recent_voters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode pollResults#badcc1a3: field recent_voters: %w", err)
			}
			p.RecentVoters = append(p.RecentVoters, value)
		}
	}
	if p.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field solution: %w", err)
		}
		p.Solution = value
	}
	if p.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pollResults#badcc1a3: field solution_entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode pollResults#badcc1a3: field solution_entities: %w", err)
			}
			p.SolutionEntities = append(p.SolutionEntities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PollResults.
var (
	_ bin.Encoder     = &PollResults{}
	_ bin.Decoder     = &PollResults{}
	_ bin.BareEncoder = &PollResults{}
	_ bin.BareDecoder = &PollResults{}
)
