// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// ContactsGetTopPeersRequest represents TL type `contacts.getTopPeers#d4982db5`.
// Get most used peers
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
type ContactsGetTopPeersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Users we've chatted most frequently with
	Correspondents bool
	// Most used bots
	BotsPm bool
	// Most used inline bots
	BotsInline bool
	// Most frequently called users
	PhoneCalls bool
	// Users to which the users often forwards messages to
	ForwardUsers bool
	// Chats to which the users often forwards messages to
	ForwardChats bool
	// Often-opened groups and supergroups
	Groups bool
	// Most frequently visited channels
	Channels bool
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// ContactsGetTopPeersRequestTypeID is TL type id of ContactsGetTopPeersRequest.
const ContactsGetTopPeersRequestTypeID = 0xd4982db5

// String implements fmt.Stringer.
func (g *ContactsGetTopPeersRequest) String() string {
	if g == nil {
		return "ContactsGetTopPeersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsGetTopPeersRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(g.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(g.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ContactsGetTopPeersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getTopPeers#d4982db5 as nil")
	}
	b.PutID(ContactsGetTopPeersRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getTopPeers#d4982db5: field flags: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// SetCorrespondents sets value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) SetCorrespondents(value bool) {
	if value {
		g.Flags.Set(0)
		g.Correspondents = true
	} else {
		g.Flags.Unset(0)
		g.Correspondents = false
	}
}

// SetBotsPm sets value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsPm(value bool) {
	if value {
		g.Flags.Set(1)
		g.BotsPm = true
	} else {
		g.Flags.Unset(1)
		g.BotsPm = false
	}
}

// SetBotsInline sets value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsInline(value bool) {
	if value {
		g.Flags.Set(2)
		g.BotsInline = true
	} else {
		g.Flags.Unset(2)
		g.BotsInline = false
	}
}

// SetPhoneCalls sets value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) SetPhoneCalls(value bool) {
	if value {
		g.Flags.Set(3)
		g.PhoneCalls = true
	} else {
		g.Flags.Unset(3)
		g.PhoneCalls = false
	}
}

// SetForwardUsers sets value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardUsers(value bool) {
	if value {
		g.Flags.Set(4)
		g.ForwardUsers = true
	} else {
		g.Flags.Unset(4)
		g.ForwardUsers = false
	}
}

// SetForwardChats sets value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardChats(value bool) {
	if value {
		g.Flags.Set(5)
		g.ForwardChats = true
	} else {
		g.Flags.Unset(5)
		g.ForwardChats = false
	}
}

// SetGroups sets value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) SetGroups(value bool) {
	if value {
		g.Flags.Set(10)
		g.Groups = true
	} else {
		g.Flags.Unset(10)
		g.Groups = false
	}
}

// SetChannels sets value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) SetChannels(value bool) {
	if value {
		g.Flags.Set(15)
		g.Channels = true
	} else {
		g.Flags.Unset(15)
		g.Channels = false
	}
}

// Decode implements bin.Decoder.
func (g *ContactsGetTopPeersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getTopPeers#d4982db5 to nil")
	}
	if err := b.ConsumeID(ContactsGetTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field flags: %w", err)
		}
	}
	g.Correspondents = g.Flags.Has(0)
	g.BotsPm = g.Flags.Has(1)
	g.BotsInline = g.Flags.Has(2)
	g.PhoneCalls = g.Flags.Has(3)
	g.ForwardUsers = g.Flags.Has(4)
	g.ForwardChats = g.Flags.Has(5)
	g.Groups = g.Flags.Has(10)
	g.Channels = g.Flags.Has(15)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetTopPeersRequest.
var (
	_ bin.Encoder = &ContactsGetTopPeersRequest{}
	_ bin.Decoder = &ContactsGetTopPeersRequest{}
)

// ContactsGetTopPeers invokes method contacts.getTopPeers#d4982db5 returning error if any.
// Get most used peers
//
// Possible errors:
//  400 TYPES_EMPTY: No top peer type was provided
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
func (c *Client) ContactsGetTopPeers(ctx context.Context, request *ContactsGetTopPeersRequest) (ContactsTopPeersClass, error) {
	var result ContactsTopPeersBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TopPeers, nil
}
