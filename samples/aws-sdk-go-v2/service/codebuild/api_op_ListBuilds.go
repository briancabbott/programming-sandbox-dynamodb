// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of build IDs, with each build ID representing a single build.
func (c *Client) ListBuilds(ctx context.Context, params *ListBuildsInput, optFns ...func(*Options)) (*ListBuildsOutput, error) {
	if params == nil {
		params = &ListBuildsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuilds", params, optFns, addOperationListBuildsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuildsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuildsInput struct {

	// During a previous call, if there are more than 100 items in the list, only the
	// first 100 items are returned, along with a unique string called a nextToken. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The order to list build IDs. Valid values include:
	//
	// * ASCENDING: List the build
	// IDs in ascending order by build ID.
	//
	// * DESCENDING: List the build IDs in
	// descending order by build ID.
	SortOrder types.SortOrderType
}

type ListBuildsOutput struct {

	// A list of build IDs, with each build ID representing a single build.
	Ids []string

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a nextToken. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBuildsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBuilds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuilds{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuilds(options.Region), middleware.Before); err != nil {
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

// ListBuildsAPIClient is a client that implements the ListBuilds operation.
type ListBuildsAPIClient interface {
	ListBuilds(context.Context, *ListBuildsInput, ...func(*Options)) (*ListBuildsOutput, error)
}

var _ ListBuildsAPIClient = (*Client)(nil)

// ListBuildsPaginatorOptions is the paginator options for ListBuilds
type ListBuildsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuildsPaginator is a paginator for ListBuilds
type ListBuildsPaginator struct {
	options   ListBuildsPaginatorOptions
	client    ListBuildsAPIClient
	params    *ListBuildsInput
	nextToken *string
	firstPage bool
}

// NewListBuildsPaginator returns a new ListBuildsPaginator
func NewListBuildsPaginator(client ListBuildsAPIClient, params *ListBuildsInput, optFns ...func(*ListBuildsPaginatorOptions)) *ListBuildsPaginator {
	options := ListBuildsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListBuildsInput{}
	}

	return &ListBuildsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuildsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListBuilds page.
func (p *ListBuildsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuildsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListBuilds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBuilds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListBuilds",
	}
}
