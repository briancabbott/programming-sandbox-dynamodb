// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of registries that you have created, with minimal registry
// information. Registries in the Deleting status will not be included in the
// results. Empty results will be returned if there are no registries available.
func (c *Client) ListRegistries(ctx context.Context, params *ListRegistriesInput, optFns ...func(*Options)) (*ListRegistriesOutput, error) {
	if params == nil {
		params = &ListRegistriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegistries", params, optFns, addOperationListRegistriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegistriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegistriesInput struct {

	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type ListRegistriesOutput struct {

	// A continuation token for paginating the returned list of tokens, returned if the
	// current segment of the list is not the last.
	NextToken *string

	// An array of RegistryDetailedListItem objects containing minimal details of each
	// registry.
	Registries []types.RegistryListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRegistriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRegistries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRegistries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegistries(options.Region), middleware.Before); err != nil {
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

// ListRegistriesAPIClient is a client that implements the ListRegistries
// operation.
type ListRegistriesAPIClient interface {
	ListRegistries(context.Context, *ListRegistriesInput, ...func(*Options)) (*ListRegistriesOutput, error)
}

var _ ListRegistriesAPIClient = (*Client)(nil)

// ListRegistriesPaginatorOptions is the paginator options for ListRegistries
type ListRegistriesPaginatorOptions struct {
	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegistriesPaginator is a paginator for ListRegistries
type ListRegistriesPaginator struct {
	options   ListRegistriesPaginatorOptions
	client    ListRegistriesAPIClient
	params    *ListRegistriesInput
	nextToken *string
	firstPage bool
}

// NewListRegistriesPaginator returns a new ListRegistriesPaginator
func NewListRegistriesPaginator(client ListRegistriesAPIClient, params *ListRegistriesInput, optFns ...func(*ListRegistriesPaginatorOptions)) *ListRegistriesPaginator {
	options := ListRegistriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListRegistriesInput{}
	}

	return &ListRegistriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegistriesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRegistries page.
func (p *ListRegistriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegistriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRegistries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRegistries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListRegistries",
	}
}
