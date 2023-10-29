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

// MessageReplyHeader represents TL type `messageReplyHeader#6eebcabd`.
// Message replies and thread¹ information
//
// Links:
//  1. https://core.telegram.org/api/threads
//
// See https://core.telegram.org/constructor/messageReplyHeader for reference.
type MessageReplyHeader struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// This is a reply to a scheduled message.
	ReplyToScheduled bool
	// Whether this message was sent in a forum topic¹ (except for the General topic).
	//
	// Links:
	//  1) https://core.telegram.org/api/forum#forum-topics
	ForumTopic bool
	// Quote field of MessageReplyHeader.
	Quote bool
	// ID of message to which this message is replying
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// For replies sent in channel discussion threads¹ of which the current user is not a
	// member, the discussion group ID
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToPeerID and GetReplyToPeerID helpers.
	ReplyToPeerID PeerClass
	// ReplyFrom field of MessageReplyHeader.
	//
	// Use SetReplyFrom and GetReplyFrom helpers.
	ReplyFrom MessageFwdHeader
	// ReplyMedia field of MessageReplyHeader.
	//
	// Use SetReplyMedia and GetReplyMedia helpers.
	ReplyMedia MessageMediaClass
	// ID of the message that started this message thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToTopID and GetReplyToTopID helpers.
	ReplyToTopID int
	// QuoteText field of MessageReplyHeader.
	//
	// Use SetQuoteText and GetQuoteText helpers.
	QuoteText string
	// QuoteEntities field of MessageReplyHeader.
	//
	// Use SetQuoteEntities and GetQuoteEntities helpers.
	QuoteEntities []MessageEntityClass
}

// MessageReplyHeaderTypeID is TL type id of MessageReplyHeader.
const MessageReplyHeaderTypeID = 0x6eebcabd

// construct implements constructor of MessageReplyHeaderClass.
func (m MessageReplyHeader) construct() MessageReplyHeaderClass { return &m }

// Ensuring interfaces in compile-time for MessageReplyHeader.
var (
	_ bin.Encoder     = &MessageReplyHeader{}
	_ bin.Decoder     = &MessageReplyHeader{}
	_ bin.BareEncoder = &MessageReplyHeader{}
	_ bin.BareDecoder = &MessageReplyHeader{}

	_ MessageReplyHeaderClass = &MessageReplyHeader{}
)

func (m *MessageReplyHeader) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.ReplyToScheduled == false) {
		return false
	}
	if !(m.ForumTopic == false) {
		return false
	}
	if !(m.Quote == false) {
		return false
	}
	if !(m.ReplyToMsgID == 0) {
		return false
	}
	if !(m.ReplyToPeerID == nil) {
		return false
	}
	if !(m.ReplyFrom.Zero()) {
		return false
	}
	if !(m.ReplyMedia == nil) {
		return false
	}
	if !(m.ReplyToTopID == 0) {
		return false
	}
	if !(m.QuoteText == "") {
		return false
	}
	if !(m.QuoteEntities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyHeader) String() string {
	if m == nil {
		return "MessageReplyHeader(nil)"
	}
	type Alias MessageReplyHeader
	return fmt.Sprintf("MessageReplyHeader%+v", Alias(*m))
}

// FillFrom fills MessageReplyHeader from given interface.
func (m *MessageReplyHeader) FillFrom(from interface {
	GetReplyToScheduled() (value bool)
	GetForumTopic() (value bool)
	GetQuote() (value bool)
	GetReplyToMsgID() (value int, ok bool)
	GetReplyToPeerID() (value PeerClass, ok bool)
	GetReplyFrom() (value MessageFwdHeader, ok bool)
	GetReplyMedia() (value MessageMediaClass, ok bool)
	GetReplyToTopID() (value int, ok bool)
	GetQuoteText() (value string, ok bool)
	GetQuoteEntities() (value []MessageEntityClass, ok bool)
}) {
	m.ReplyToScheduled = from.GetReplyToScheduled()
	m.ForumTopic = from.GetForumTopic()
	m.Quote = from.GetQuote()
	if val, ok := from.GetReplyToMsgID(); ok {
		m.ReplyToMsgID = val
	}

	if val, ok := from.GetReplyToPeerID(); ok {
		m.ReplyToPeerID = val
	}

	if val, ok := from.GetReplyFrom(); ok {
		m.ReplyFrom = val
	}

	if val, ok := from.GetReplyMedia(); ok {
		m.ReplyMedia = val
	}

	if val, ok := from.GetReplyToTopID(); ok {
		m.ReplyToTopID = val
	}

	if val, ok := from.GetQuoteText(); ok {
		m.QuoteText = val
	}

	if val, ok := from.GetQuoteEntities(); ok {
		m.QuoteEntities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyHeader) TypeID() uint32 {
	return MessageReplyHeaderTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyHeader) TypeName() string {
	return "messageReplyHeader"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyHeader) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyHeader",
		ID:   MessageReplyHeaderTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReplyToScheduled",
			SchemaName: "reply_to_scheduled",
			Null:       !m.Flags.Has(2),
		},
		{
			Name:       "ForumTopic",
			SchemaName: "forum_topic",
			Null:       !m.Flags.Has(3),
		},
		{
			Name:       "Quote",
			SchemaName: "quote",
			Null:       !m.Flags.Has(9),
		},
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
			Null:       !m.Flags.Has(4),
		},
		{
			Name:       "ReplyToPeerID",
			SchemaName: "reply_to_peer_id",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "ReplyFrom",
			SchemaName: "reply_from",
			Null:       !m.Flags.Has(5),
		},
		{
			Name:       "ReplyMedia",
			SchemaName: "reply_media",
			Null:       !m.Flags.Has(8),
		},
		{
			Name:       "ReplyToTopID",
			SchemaName: "reply_to_top_id",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "QuoteText",
			SchemaName: "quote_text",
			Null:       !m.Flags.Has(6),
		},
		{
			Name:       "QuoteEntities",
			SchemaName: "quote_entities",
			Null:       !m.Flags.Has(7),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessageReplyHeader) SetFlags() {
	if !(m.ReplyToScheduled == false) {
		m.Flags.Set(2)
	}
	if !(m.ForumTopic == false) {
		m.Flags.Set(3)
	}
	if !(m.Quote == false) {
		m.Flags.Set(9)
	}
	if !(m.ReplyToMsgID == 0) {
		m.Flags.Set(4)
	}
	if !(m.ReplyToPeerID == nil) {
		m.Flags.Set(0)
	}
	if !(m.ReplyFrom.Zero()) {
		m.Flags.Set(5)
	}
	if !(m.ReplyMedia == nil) {
		m.Flags.Set(8)
	}
	if !(m.ReplyToTopID == 0) {
		m.Flags.Set(1)
	}
	if !(m.QuoteText == "") {
		m.Flags.Set(6)
	}
	if !(m.QuoteEntities == nil) {
		m.Flags.Set(7)
	}
}

// Encode implements bin.Encoder.
func (m *MessageReplyHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyHeader#6eebcabd as nil")
	}
	b.PutID(MessageReplyHeaderTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyHeader) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyHeader#6eebcabd as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field flags: %w", err)
	}
	if m.Flags.Has(4) {
		b.PutInt(m.ReplyToMsgID)
	}
	if m.Flags.Has(0) {
		if m.ReplyToPeerID == nil {
			return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field reply_to_peer_id is nil")
		}
		if err := m.ReplyToPeerID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field reply_to_peer_id: %w", err)
		}
	}
	if m.Flags.Has(5) {
		if err := m.ReplyFrom.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field reply_from: %w", err)
		}
	}
	if m.Flags.Has(8) {
		if m.ReplyMedia == nil {
			return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field reply_media is nil")
		}
		if err := m.ReplyMedia.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field reply_media: %w", err)
		}
	}
	if m.Flags.Has(1) {
		b.PutInt(m.ReplyToTopID)
	}
	if m.Flags.Has(6) {
		b.PutString(m.QuoteText)
	}
	if m.Flags.Has(7) {
		b.PutVectorHeader(len(m.QuoteEntities))
		for idx, v := range m.QuoteEntities {
			if v == nil {
				return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field quote_entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messageReplyHeader#6eebcabd: field quote_entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyHeader#6eebcabd to nil")
	}
	if err := b.ConsumeID(MessageReplyHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyHeader) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyHeader#6eebcabd to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field flags: %w", err)
		}
	}
	m.ReplyToScheduled = m.Flags.Has(2)
	m.ForumTopic = m.Flags.Has(3)
	m.Quote = m.Flags.Has(9)
	if m.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field reply_to_msg_id: %w", err)
		}
		m.ReplyToMsgID = value
	}
	if m.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field reply_to_peer_id: %w", err)
		}
		m.ReplyToPeerID = value
	}
	if m.Flags.Has(5) {
		if err := m.ReplyFrom.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field reply_from: %w", err)
		}
	}
	if m.Flags.Has(8) {
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field reply_media: %w", err)
		}
		m.ReplyMedia = value
	}
	if m.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field reply_to_top_id: %w", err)
		}
		m.ReplyToTopID = value
	}
	if m.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field quote_text: %w", err)
		}
		m.QuoteText = value
	}
	if m.Flags.Has(7) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field quote_entities: %w", err)
		}

		if headerLen > 0 {
			m.QuoteEntities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyHeader#6eebcabd: field quote_entities: %w", err)
			}
			m.QuoteEntities = append(m.QuoteEntities, value)
		}
	}
	return nil
}

// SetReplyToScheduled sets value of ReplyToScheduled conditional field.
func (m *MessageReplyHeader) SetReplyToScheduled(value bool) {
	if value {
		m.Flags.Set(2)
		m.ReplyToScheduled = true
	} else {
		m.Flags.Unset(2)
		m.ReplyToScheduled = false
	}
}

// GetReplyToScheduled returns value of ReplyToScheduled conditional field.
func (m *MessageReplyHeader) GetReplyToScheduled() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(2)
}

// SetForumTopic sets value of ForumTopic conditional field.
func (m *MessageReplyHeader) SetForumTopic(value bool) {
	if value {
		m.Flags.Set(3)
		m.ForumTopic = true
	} else {
		m.Flags.Unset(3)
		m.ForumTopic = false
	}
}

// GetForumTopic returns value of ForumTopic conditional field.
func (m *MessageReplyHeader) GetForumTopic() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(3)
}

// SetQuote sets value of Quote conditional field.
func (m *MessageReplyHeader) SetQuote(value bool) {
	if value {
		m.Flags.Set(9)
		m.Quote = true
	} else {
		m.Flags.Unset(9)
		m.Quote = false
	}
}

// GetQuote returns value of Quote conditional field.
func (m *MessageReplyHeader) GetQuote() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(9)
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (m *MessageReplyHeader) SetReplyToMsgID(value int) {
	m.Flags.Set(4)
	m.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToMsgID() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(4) {
		return value, false
	}
	return m.ReplyToMsgID, true
}

// SetReplyToPeerID sets value of ReplyToPeerID conditional field.
func (m *MessageReplyHeader) SetReplyToPeerID(value PeerClass) {
	m.Flags.Set(0)
	m.ReplyToPeerID = value
}

// GetReplyToPeerID returns value of ReplyToPeerID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToPeerID() (value PeerClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.ReplyToPeerID, true
}

// SetReplyFrom sets value of ReplyFrom conditional field.
func (m *MessageReplyHeader) SetReplyFrom(value MessageFwdHeader) {
	m.Flags.Set(5)
	m.ReplyFrom = value
}

// GetReplyFrom returns value of ReplyFrom conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyFrom() (value MessageFwdHeader, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(5) {
		return value, false
	}
	return m.ReplyFrom, true
}

// SetReplyMedia sets value of ReplyMedia conditional field.
func (m *MessageReplyHeader) SetReplyMedia(value MessageMediaClass) {
	m.Flags.Set(8)
	m.ReplyMedia = value
}

// GetReplyMedia returns value of ReplyMedia conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyMedia() (value MessageMediaClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(8) {
		return value, false
	}
	return m.ReplyMedia, true
}

// SetReplyToTopID sets value of ReplyToTopID conditional field.
func (m *MessageReplyHeader) SetReplyToTopID(value int) {
	m.Flags.Set(1)
	m.ReplyToTopID = value
}

// GetReplyToTopID returns value of ReplyToTopID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToTopID() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.ReplyToTopID, true
}

// SetQuoteText sets value of QuoteText conditional field.
func (m *MessageReplyHeader) SetQuoteText(value string) {
	m.Flags.Set(6)
	m.QuoteText = value
}

// GetQuoteText returns value of QuoteText conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetQuoteText() (value string, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(6) {
		return value, false
	}
	return m.QuoteText, true
}

// SetQuoteEntities sets value of QuoteEntities conditional field.
func (m *MessageReplyHeader) SetQuoteEntities(value []MessageEntityClass) {
	m.Flags.Set(7)
	m.QuoteEntities = value
}

// GetQuoteEntities returns value of QuoteEntities conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetQuoteEntities() (value []MessageEntityClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(7) {
		return value, false
	}
	return m.QuoteEntities, true
}

// MapQuoteEntities returns field QuoteEntities wrapped in MessageEntityClassArray helper.
func (m *MessageReplyHeader) MapQuoteEntities() (value MessageEntityClassArray, ok bool) {
	if !m.Flags.Has(7) {
		return value, false
	}
	return MessageEntityClassArray(m.QuoteEntities), true
}

// MessageReplyStoryHeader represents TL type `messageReplyStoryHeader#9c98bfc1`.
//
// See https://core.telegram.org/constructor/messageReplyStoryHeader for reference.
type MessageReplyStoryHeader struct {
	// UserID field of MessageReplyStoryHeader.
	UserID int64
	// StoryID field of MessageReplyStoryHeader.
	StoryID int
}

// MessageReplyStoryHeaderTypeID is TL type id of MessageReplyStoryHeader.
const MessageReplyStoryHeaderTypeID = 0x9c98bfc1

// construct implements constructor of MessageReplyHeaderClass.
func (m MessageReplyStoryHeader) construct() MessageReplyHeaderClass { return &m }

// Ensuring interfaces in compile-time for MessageReplyStoryHeader.
var (
	_ bin.Encoder     = &MessageReplyStoryHeader{}
	_ bin.Decoder     = &MessageReplyStoryHeader{}
	_ bin.BareEncoder = &MessageReplyStoryHeader{}
	_ bin.BareDecoder = &MessageReplyStoryHeader{}

	_ MessageReplyHeaderClass = &MessageReplyStoryHeader{}
)

func (m *MessageReplyStoryHeader) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyStoryHeader) String() string {
	if m == nil {
		return "MessageReplyStoryHeader(nil)"
	}
	type Alias MessageReplyStoryHeader
	return fmt.Sprintf("MessageReplyStoryHeader%+v", Alias(*m))
}

// FillFrom fills MessageReplyStoryHeader from given interface.
func (m *MessageReplyStoryHeader) FillFrom(from interface {
	GetUserID() (value int64)
	GetStoryID() (value int)
}) {
	m.UserID = from.GetUserID()
	m.StoryID = from.GetStoryID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyStoryHeader) TypeID() uint32 {
	return MessageReplyStoryHeaderTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyStoryHeader) TypeName() string {
	return "messageReplyStoryHeader"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyStoryHeader) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyStoryHeader",
		ID:   MessageReplyStoryHeaderTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReplyStoryHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyStoryHeader#9c98bfc1 as nil")
	}
	b.PutID(MessageReplyStoryHeaderTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyStoryHeader) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyStoryHeader#9c98bfc1 as nil")
	}
	b.PutLong(m.UserID)
	b.PutInt(m.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyStoryHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyStoryHeader#9c98bfc1 to nil")
	}
	if err := b.ConsumeID(MessageReplyStoryHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyStoryHeader#9c98bfc1: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyStoryHeader) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyStoryHeader#9c98bfc1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyStoryHeader#9c98bfc1: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyStoryHeader#9c98bfc1: field story_id: %w", err)
		}
		m.StoryID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageReplyStoryHeader) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}

// GetStoryID returns value of StoryID field.
func (m *MessageReplyStoryHeader) GetStoryID() (value int) {
	if m == nil {
		return
	}
	return m.StoryID
}

// MessageReplyHeaderClassName is schema name of MessageReplyHeaderClass.
const MessageReplyHeaderClassName = "MessageReplyHeader"

// MessageReplyHeaderClass represents MessageReplyHeader generic type.
//
// See https://core.telegram.org/type/MessageReplyHeader for reference.
//
// Example:
//
//	g, err := tg.DecodeMessageReplyHeader(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessageReplyHeader: // messageReplyHeader#6eebcabd
//	case *tg.MessageReplyStoryHeader: // messageReplyStoryHeader#9c98bfc1
//	default: panic(v)
//	}
type MessageReplyHeaderClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageReplyHeaderClass

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
}

// DecodeMessageReplyHeader implements binary de-serialization for MessageReplyHeaderClass.
func DecodeMessageReplyHeader(buf *bin.Buffer) (MessageReplyHeaderClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageReplyHeaderTypeID:
		// Decoding messageReplyHeader#6eebcabd.
		v := MessageReplyHeader{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyHeaderClass: %w", err)
		}
		return &v, nil
	case MessageReplyStoryHeaderTypeID:
		// Decoding messageReplyStoryHeader#9c98bfc1.
		v := MessageReplyStoryHeader{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyHeaderClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageReplyHeaderClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageReplyHeader boxes the MessageReplyHeaderClass providing a helper.
type MessageReplyHeaderBox struct {
	MessageReplyHeader MessageReplyHeaderClass
}

// Decode implements bin.Decoder for MessageReplyHeaderBox.
func (b *MessageReplyHeaderBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageReplyHeaderBox to nil")
	}
	v, err := DecodeMessageReplyHeader(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageReplyHeader = v
	return nil
}

// Encode implements bin.Encode for MessageReplyHeaderBox.
func (b *MessageReplyHeaderBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageReplyHeader == nil {
		return fmt.Errorf("unable to encode MessageReplyHeaderClass as nil")
	}
	return b.MessageReplyHeader.Encode(buf)
}
