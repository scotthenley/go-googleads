// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/scotthenley/go-googleads/pb/v8/resources"
	servicespb "github.com/scotthenley/go-googleads/pb/v8/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newFeedMappingClientHook clientHook

// FeedMappingCallOptions contains the retry settings for each method of FeedMappingClient.
type FeedMappingCallOptions struct {
	GetFeedMapping     []gax.CallOption
	MutateFeedMappings []gax.CallOption
}

func defaultFeedMappingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFeedMappingCallOptions() *FeedMappingCallOptions {
	return &FeedMappingCallOptions{
		GetFeedMapping: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MutateFeedMappings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalFeedMappingClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedMappingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedMapping(context.Context, *servicespb.GetFeedMappingRequest, ...gax.CallOption) (*resourcespb.FeedMapping, error)
	MutateFeedMappings(context.Context, *servicespb.MutateFeedMappingsRequest, ...gax.CallOption) (*servicespb.MutateFeedMappingsResponse, error)
}

// FeedMappingClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage feed mappings.
type FeedMappingClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedMappingClient

	// The call options for this service.
	CallOptions *FeedMappingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedMappingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedMappingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedMappingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedMapping returns the requested feed mapping in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedMappingClient) GetFeedMapping(ctx context.Context, req *servicespb.GetFeedMappingRequest, opts ...gax.CallOption) (*resourcespb.FeedMapping, error) {
	return c.internalClient.GetFeedMapping(ctx, req, opts...)
}

// MutateFeedMappings creates or removes feed mappings. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// DistinctError (at )
// FeedMappingError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NotEmptyError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *FeedMappingClient) MutateFeedMappings(ctx context.Context, req *servicespb.MutateFeedMappingsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedMappingsResponse, error) {
	return c.internalClient.MutateFeedMappings(ctx, req, opts...)
}

// feedMappingGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedMappingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedMappingClient
	CallOptions **FeedMappingCallOptions

	// The gRPC API client.
	feedMappingClient servicespb.FeedMappingServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedMappingClient creates a new feed mapping service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage feed mappings.
func NewFeedMappingClient(ctx context.Context, opts ...option.ClientOption) (*FeedMappingClient, error) {
	clientOpts := defaultFeedMappingGRPCClientOptions()
	if newFeedMappingClientHook != nil {
		hookOpts, err := newFeedMappingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FeedMappingClient{CallOptions: defaultFeedMappingCallOptions()}

	c := &feedMappingGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		feedMappingClient: servicespb.NewFeedMappingServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedMappingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedMappingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedMappingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedMappingGRPCClient) GetFeedMapping(ctx context.Context, req *servicespb.GetFeedMappingRequest, opts ...gax.CallOption) (*resourcespb.FeedMapping, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedMapping[0:len((*c.CallOptions).GetFeedMapping):len((*c.CallOptions).GetFeedMapping)], opts...)
	var resp *resourcespb.FeedMapping
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedMappingClient.GetFeedMapping(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *feedMappingGRPCClient) MutateFeedMappings(ctx context.Context, req *servicespb.MutateFeedMappingsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedMappingsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateFeedMappings[0:len((*c.CallOptions).MutateFeedMappings):len((*c.CallOptions).MutateFeedMappings)], opts...)
	var resp *servicespb.MutateFeedMappingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedMappingClient.MutateFeedMappings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
