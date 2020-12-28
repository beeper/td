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

// AccountSendConfirmPhoneCodeRequest represents TL type `account.sendConfirmPhoneCode#1b3faa88`.
// Send confirmation code to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// See https://core.telegram.org/method/account.sendConfirmPhoneCode for reference.
type AccountSendConfirmPhoneCodeRequest struct {
	// The hash from the service notification, for more info click here »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/account-deletion
	Hash string
	// Phone code settings
	Settings CodeSettings
}

// AccountSendConfirmPhoneCodeRequestTypeID is TL type id of AccountSendConfirmPhoneCodeRequest.
const AccountSendConfirmPhoneCodeRequestTypeID = 0x1b3faa88

// String implements fmt.Stringer.
func (s *AccountSendConfirmPhoneCodeRequest) String() string {
	if s == nil {
		return "AccountSendConfirmPhoneCodeRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountSendConfirmPhoneCodeRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(s.Hash))
	sb.WriteString(",\n")
	sb.WriteString("\tSettings: ")
	sb.WriteString(s.Settings.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AccountSendConfirmPhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendConfirmPhoneCode#1b3faa88 as nil")
	}
	b.PutID(AccountSendConfirmPhoneCodeRequestTypeID)
	b.PutString(s.Hash)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendConfirmPhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendConfirmPhoneCode#1b3faa88 to nil")
	}
	if err := b.ConsumeID(AccountSendConfirmPhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendConfirmPhoneCodeRequest.
var (
	_ bin.Encoder = &AccountSendConfirmPhoneCodeRequest{}
	_ bin.Decoder = &AccountSendConfirmPhoneCodeRequest{}
)

// AccountSendConfirmPhoneCode invokes method account.sendConfirmPhoneCode#1b3faa88 returning error if any.
// Send confirmation code to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// Possible errors:
//  400 HASH_INVALID: The provided hash is invalid
//
// See https://core.telegram.org/method/account.sendConfirmPhoneCode for reference.
func (c *Client) AccountSendConfirmPhoneCode(ctx context.Context, request *AccountSendConfirmPhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
