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

// ChannelsInviteToChannelRequest represents TL type `channels.inviteToChannel#199f3a6c`.
// Invite users to a channel/supergroup
//
// See https://core.telegram.org/method/channels.inviteToChannel for reference.
type ChannelsInviteToChannelRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// Users to invite
	Users []InputUserClass
}

// ChannelsInviteToChannelRequestTypeID is TL type id of ChannelsInviteToChannelRequest.
const ChannelsInviteToChannelRequestTypeID = 0x199f3a6c

// String implements fmt.Stringer.
func (i *ChannelsInviteToChannelRequest) String() string {
	if i == nil {
		return "ChannelsInviteToChannelRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsInviteToChannelRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(i.Channel.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range i.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *ChannelsInviteToChannelRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode channels.inviteToChannel#199f3a6c as nil")
	}
	b.PutID(ChannelsInviteToChannelRequestTypeID)
	if i.Channel == nil {
		return fmt.Errorf("unable to encode channels.inviteToChannel#199f3a6c: field channel is nil")
	}
	if err := i.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.inviteToChannel#199f3a6c: field channel: %w", err)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.inviteToChannel#199f3a6c: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.inviteToChannel#199f3a6c: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ChannelsInviteToChannelRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode channels.inviteToChannel#199f3a6c to nil")
	}
	if err := b.ConsumeID(ChannelsInviteToChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.inviteToChannel#199f3a6c: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.inviteToChannel#199f3a6c: field channel: %w", err)
		}
		i.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.inviteToChannel#199f3a6c: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.inviteToChannel#199f3a6c: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsInviteToChannelRequest.
var (
	_ bin.Encoder = &ChannelsInviteToChannelRequest{}
	_ bin.Decoder = &ChannelsInviteToChannelRequest{}
)

// ChannelsInviteToChannel invokes method channels.inviteToChannel#199f3a6c returning error if any.
// Invite users to a channel/supergroup
//
// Possible errors:
//  400 BOTS_TOO_MUCH: There are too many bots in this chat/channel
//  400 BOT_GROUPS_BLOCKED: This bot can't be added to groups
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_INVALID: Invalid chat
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
//  400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels
//  400 USER_BLOCKED: User blocked
//  400 USER_BOT: Bots can only be admins in channels.
//  403 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups
//  400 USER_ID_INVALID: The provided user ID is invalid
//  400 USER_KICKED: This user was kicked from this supergroup/channel
//  400 USER_NOT_MUTUAL_CONTACT: The provided user is not a mutual contact
//  403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this
//
// See https://core.telegram.org/method/channels.inviteToChannel for reference.
func (c *Client) ChannelsInviteToChannel(ctx context.Context, request *ChannelsInviteToChannelRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
