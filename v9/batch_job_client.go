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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/scotthenley/go-googleads/pb/v9/resources"
	servicespb "github.com/scotthenley/go-googleads/pb/v9/services"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newBatchJobClientHook clientHook

// BatchJobCallOptions contains the retry settings for each method of BatchJobClient.
type BatchJobCallOptions struct {
	MutateBatchJob        []gax.CallOption
	GetBatchJob           []gax.CallOption
	ListBatchJobResults   []gax.CallOption
	RunBatchJob           []gax.CallOption
	AddBatchJobOperations []gax.CallOption
}

func defaultBatchJobGRPCClientOptions() []option.ClientOption {
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

func defaultBatchJobCallOptions() *BatchJobCallOptions {
	return &BatchJobCallOptions{
		MutateBatchJob: []gax.CallOption{
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
		GetBatchJob: []gax.CallOption{
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
		ListBatchJobResults: []gax.CallOption{
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
		RunBatchJob: []gax.CallOption{
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
		AddBatchJobOperations: []gax.CallOption{
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

// internalBatchJobClient is an interface that defines the methods availaible from Google Ads API.
type internalBatchJobClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateBatchJob(context.Context, *servicespb.MutateBatchJobRequest, ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error)
	GetBatchJob(context.Context, *servicespb.GetBatchJobRequest, ...gax.CallOption) (*resourcespb.BatchJob, error)
	ListBatchJobResults(context.Context, *servicespb.ListBatchJobResultsRequest, ...gax.CallOption) *BatchJobResultIterator
	RunBatchJob(context.Context, *servicespb.RunBatchJobRequest, ...gax.CallOption) (*RunBatchJobOperation, error)
	RunBatchJobOperation(name string) *RunBatchJobOperation
	AddBatchJobOperations(context.Context, *servicespb.AddBatchJobOperationsRequest, ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error)
}

// BatchJobClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage batch jobs.
type BatchJobClient struct {
	// The internal transport-dependent client.
	internalClient internalBatchJobClient

	// The call options for this service.
	CallOptions *BatchJobCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BatchJobClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BatchJobClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BatchJobClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateBatchJob mutates a batch job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *BatchJobClient) MutateBatchJob(ctx context.Context, req *servicespb.MutateBatchJobRequest, opts ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error) {
	return c.internalClient.MutateBatchJob(ctx, req, opts...)
}

// GetBatchJob returns the batch job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BatchJobClient) GetBatchJob(ctx context.Context, req *servicespb.GetBatchJobRequest, opts ...gax.CallOption) (*resourcespb.BatchJob, error) {
	return c.internalClient.GetBatchJob(ctx, req, opts...)
}

// ListBatchJobResults returns the results of the batch job. The job must be done.
// Supports standard list paging.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BatchJobClient) ListBatchJobResults(ctx context.Context, req *servicespb.ListBatchJobResultsRequest, opts ...gax.CallOption) *BatchJobResultIterator {
	return c.internalClient.ListBatchJobResults(ctx, req, opts...)
}

// RunBatchJob runs the batch job.
//
// The Operation.metadata field type is BatchJobMetadata. When finished, the
// long running operation will not contain errors or a response. Instead, use
// ListBatchJobResults to get the results of the job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BatchJobClient) RunBatchJob(ctx context.Context, req *servicespb.RunBatchJobRequest, opts ...gax.CallOption) (*RunBatchJobOperation, error) {
	return c.internalClient.RunBatchJob(ctx, req, opts...)
}

// RunBatchJobOperation returns a new RunBatchJobOperation from a given name.
// The name must be that of a previously created RunBatchJobOperation, possibly from a different process.
func (c *BatchJobClient) RunBatchJobOperation(name string) *RunBatchJobOperation {
	return c.internalClient.RunBatchJobOperation(name)
}

// AddBatchJobOperations add operations to the batch job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *BatchJobClient) AddBatchJobOperations(ctx context.Context, req *servicespb.AddBatchJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error) {
	return c.internalClient.AddBatchJobOperations(ctx, req, opts...)
}

// batchJobGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type batchJobGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BatchJobClient
	CallOptions **BatchJobCallOptions

	// The gRPC API client.
	batchJobClient servicespb.BatchJobServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBatchJobClient creates a new batch job service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage batch jobs.
func NewBatchJobClient(ctx context.Context, opts ...option.ClientOption) (*BatchJobClient, error) {
	clientOpts := defaultBatchJobGRPCClientOptions()
	if newBatchJobClientHook != nil {
		hookOpts, err := newBatchJobClientHook(ctx, clientHookParams{})
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
	client := BatchJobClient{CallOptions: defaultBatchJobCallOptions()}

	c := &batchJobGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		batchJobClient:   servicespb.NewBatchJobServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *batchJobGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *batchJobGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *batchJobGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *batchJobGRPCClient) MutateBatchJob(ctx context.Context, req *servicespb.MutateBatchJobRequest, opts ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateBatchJob[0:len((*c.CallOptions).MutateBatchJob):len((*c.CallOptions).MutateBatchJob)], opts...)
	var resp *servicespb.MutateBatchJobResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchJobClient.MutateBatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *batchJobGRPCClient) GetBatchJob(ctx context.Context, req *servicespb.GetBatchJobRequest, opts ...gax.CallOption) (*resourcespb.BatchJob, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBatchJob[0:len((*c.CallOptions).GetBatchJob):len((*c.CallOptions).GetBatchJob)], opts...)
	var resp *resourcespb.BatchJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchJobClient.GetBatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *batchJobGRPCClient) ListBatchJobResults(ctx context.Context, req *servicespb.ListBatchJobResultsRequest, opts ...gax.CallOption) *BatchJobResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListBatchJobResults[0:len((*c.CallOptions).ListBatchJobResults):len((*c.CallOptions).ListBatchJobResults)], opts...)
	it := &BatchJobResultIterator{}
	req = proto.Clone(req).(*servicespb.ListBatchJobResultsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*servicespb.BatchJobResult, string, error) {
		resp := &servicespb.ListBatchJobResultsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.batchJobClient.ListBatchJobResults(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *batchJobGRPCClient) RunBatchJob(ctx context.Context, req *servicespb.RunBatchJobRequest, opts ...gax.CallOption) (*RunBatchJobOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunBatchJob[0:len((*c.CallOptions).RunBatchJob):len((*c.CallOptions).RunBatchJob)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchJobClient.RunBatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunBatchJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *batchJobGRPCClient) AddBatchJobOperations(ctx context.Context, req *servicespb.AddBatchJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AddBatchJobOperations[0:len((*c.CallOptions).AddBatchJobOperations):len((*c.CallOptions).AddBatchJobOperations)], opts...)
	var resp *servicespb.AddBatchJobOperationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchJobClient.AddBatchJobOperations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RunBatchJobOperation manages a long-running operation from RunBatchJob.
type RunBatchJobOperation struct {
	lro *longrunning.Operation
}

// RunBatchJobOperation returns a new RunBatchJobOperation from a given name.
// The name must be that of a previously created RunBatchJobOperation, possibly from a different process.
func (c *batchJobGRPCClient) RunBatchJobOperation(name string) *RunBatchJobOperation {
	return &RunBatchJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunBatchJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunBatchJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunBatchJobOperation) Metadata() (*resourcespb.BatchJob_BatchJobMetadata, error) {
	var meta resourcespb.BatchJob_BatchJobMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunBatchJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunBatchJobOperation) Name() string {
	return op.lro.Name()
}

// BatchJobResultIterator manages a stream of *servicespb.BatchJobResult.
type BatchJobResultIterator struct {
	items    []*servicespb.BatchJobResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*servicespb.BatchJobResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BatchJobResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BatchJobResultIterator) Next() (*servicespb.BatchJobResult, error) {
	var item *servicespb.BatchJobResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BatchJobResultIterator) bufLen() int {
	return len(it.items)
}

func (it *BatchJobResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
