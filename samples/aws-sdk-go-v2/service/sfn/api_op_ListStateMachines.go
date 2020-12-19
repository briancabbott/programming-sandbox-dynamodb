// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the existing state machines. If nextToken is returned, there are more
// results available. The value of nextToken is a unique pagination token for each
// page. Make the call again using the returned token to retrieve the next page.
// Keep all other arguments unchanged. Each pagination token expires after 24
// hours. Using an expired pagination token will return an HTTP 400 InvalidToken
// error. This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
func (c *Client) ListStateMachines(ctx context.Context, params *ListStateMachinesInput, optFns ...func(*Options)) (*ListStateMachinesOutput, error) {
	if params == nil {
		params = &ListStateMachinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStateMachines", params, optFns, addOperationListStateMachinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStateMachinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStateMachinesInput struct {

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string
}

type ListStateMachinesOutput struct {
	StateMachines []types.StateMachineListItem

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStateMachinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListStateMachines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListStateMachines{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStateMachines(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListStateMachinesAPIClient is a client that implements the ListStateMachines
// operation.
type ListStateMachinesAPIClient interface {
	ListStateMachines(context.Context, *ListStateMachinesInput, ...func(*Options)) (*ListStateMachinesOutput, error)
}

var _ ListStateMachinesAPIClient = (*Client)(nil)

// ListStateMachinesPaginatorOptions is the paginator options for ListStateMachines
type ListStateMachinesPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStateMachinesPaginator is a paginator for ListStateMachines
type ListStateMachinesPaginator struct {
	options   ListStateMachinesPaginatorOptions
	client    ListStateMachinesAPIClient
	params    *ListStateMachinesInput
	nextToken *string
	firstPage bool
}

// NewListStateMachinesPaginator returns a new ListStateMachinesPaginator
func NewListStateMachinesPaginator(client ListStateMachinesAPIClient, params *ListStateMachinesInput, optFns ...func(*ListStateMachinesPaginatorOptions)) *ListStateMachinesPaginator {
	options := ListStateMachinesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListStateMachinesInput{}
	}

	return &ListStateMachinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStateMachinesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListStateMachines page.
func (p *ListStateMachinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStateMachinesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListStateMachines(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListStateMachines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "ListStateMachines",
	}
}
