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

// UploadGetFileHashesRequest represents TL type `upload.getFileHashes#c7025931`.
// Get SHA256 hashes for verifying downloaded files
//
// See https://core.telegram.org/method/upload.getFileHashes for reference.
type UploadGetFileHashesRequest struct {
	// File
	Location InputFileLocationClass
	// Offset from which to get file hashes
	Offset int
}

// UploadGetFileHashesRequestTypeID is TL type id of UploadGetFileHashesRequest.
const UploadGetFileHashesRequestTypeID = 0xc7025931

// String implements fmt.Stringer.
func (g *UploadGetFileHashesRequest) String() string {
	if g == nil {
		return "UploadGetFileHashesRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadGetFileHashesRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tLocation: ")
	sb.WriteString(g.Location.String())
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *UploadGetFileHashesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFileHashes#c7025931 as nil")
	}
	b.PutID(UploadGetFileHashesRequestTypeID)
	if g.Location == nil {
		return fmt.Errorf("unable to encode upload.getFileHashes#c7025931: field location is nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFileHashes#c7025931: field location: %w", err)
	}
	b.PutInt(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetFileHashesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFileHashes#c7025931 to nil")
	}
	if err := b.ConsumeID(UploadGetFileHashesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: %w", err)
	}
	{
		value, err := DecodeInputFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: field location: %w", err)
		}
		g.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadGetFileHashesRequest.
var (
	_ bin.Encoder = &UploadGetFileHashesRequest{}
	_ bin.Decoder = &UploadGetFileHashesRequest{}
)

// UploadGetFileHashes invokes method upload.getFileHashes#c7025931 returning error if any.
// Get SHA256 hashes for verifying downloaded files
//
// Possible errors:
//  400 LOCATION_INVALID: The provided location is invalid
//
// See https://core.telegram.org/method/upload.getFileHashes for reference.
// Can be used by bots.
func (c *Client) UploadGetFileHashes(ctx context.Context, request *UploadGetFileHashesRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
