// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the discoverers.
func (c *Client) ListDiscoverers(ctx context.Context, params *ListDiscoverersInput, optFns ...func(*Options)) (*ListDiscoverersOutput, error) {
	if params == nil {
		params = &ListDiscoverersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDiscoverers", params, optFns, addOperationListDiscoverersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDiscoverersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDiscoverersInput struct {

	// Specifying this limits the results to only those discoverer IDs that start with
	// the specified prefix.
	DiscovererIdPrefix *string

	Limit int32

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// Specifying this limits the results to only those ARNs that start with the
	// specified prefix.
	SourceArnPrefix *string
}

type ListDiscoverersOutput struct {

	// An array of DiscovererSummary information.
	Discoverers []types.DiscovererSummary

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDiscoverersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDiscoverers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDiscoverers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDiscoverers(options.Region), middleware.Before); err != nil {
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

// ListDiscoverersAPIClient is a client that implements the ListDiscoverers
// operation.
type ListDiscoverersAPIClient interface {
	ListDiscoverers(context.Context, *ListDiscoverersInput, ...func(*Options)) (*ListDiscoverersOutput, error)
}

var _ ListDiscoverersAPIClient = (*Client)(nil)

// ListDiscoverersPaginatorOptions is the paginator options for ListDiscoverers
type ListDiscoverersPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDiscoverersPaginator is a paginator for ListDiscoverers
type ListDiscoverersPaginator struct {
	options   ListDiscoverersPaginatorOptions
	client    ListDiscoverersAPIClient
	params    *ListDiscoverersInput
	nextToken *string
	firstPage bool
}

// NewListDiscoverersPaginator returns a new ListDiscoverersPaginator
func NewListDiscoverersPaginator(client ListDiscoverersAPIClient, params *ListDiscoverersInput, optFns ...func(*ListDiscoverersPaginatorOptions)) *ListDiscoverersPaginator {
	options := ListDiscoverersPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDiscoverersInput{}
	}

	return &ListDiscoverersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDiscoverersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDiscoverers page.
func (p *ListDiscoverersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDiscoverersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.ListDiscoverers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDiscoverers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "ListDiscoverers",
	}
}
