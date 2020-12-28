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

// AccountChangePhoneRequest represents TL type `account.changePhone#70c32edb`.
// Change the phone number of the current account
//
// See https://core.telegram.org/method/account.changePhone for reference.
type AccountChangePhoneRequest struct {
	// New phone number
	PhoneNumber string
	// Phone code hash received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCodeHash string
	// Phone code received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCode string
}

// AccountChangePhoneRequestTypeID is TL type id of AccountChangePhoneRequest.
const AccountChangePhoneRequestTypeID = 0x70c32edb

// String implements fmt.Stringer.
func (c *AccountChangePhoneRequest) String() string {
	if c == nil {
		return "AccountChangePhoneRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountChangePhoneRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPhoneNumber: ")
	sb.WriteString(fmt.Sprint(c.PhoneNumber))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoneCodeHash: ")
	sb.WriteString(fmt.Sprint(c.PhoneCodeHash))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoneCode: ")
	sb.WriteString(fmt.Sprint(c.PhoneCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *AccountChangePhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changePhone#70c32edb as nil")
	}
	b.PutID(AccountChangePhoneRequestTypeID)
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountChangePhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changePhone#70c32edb to nil")
	}
	if err := b.ConsumeID(AccountChangePhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.changePhone#70c32edb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code: %w", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountChangePhoneRequest.
var (
	_ bin.Encoder = &AccountChangePhoneRequest{}
	_ bin.Decoder = &AccountChangePhoneRequest{}
)

// AccountChangePhone invokes method account.changePhone#70c32edb returning error if any.
// Change the phone number of the current account
//
// Possible errors:
//  400 PHONE_CODE_EMPTY: phone_code is missing
//  400 PHONE_NUMBER_INVALID: The phone number is invalid
//
// See https://core.telegram.org/method/account.changePhone for reference.
func (c *Client) AccountChangePhone(ctx context.Context, request *AccountChangePhoneRequest) (UserClass, error) {
	var result UserBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
