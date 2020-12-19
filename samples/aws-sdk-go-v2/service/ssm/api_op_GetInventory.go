// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Query inventory information.
func (c *Client) GetInventory(ctx context.Context, params *GetInventoryInput, optFns ...func(*Options)) (*GetInventoryOutput, error) {
	if params == nil {
		params = &GetInventoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInventory", params, optFns, addOperationGetInventoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInventoryInput struct {

	// Returns counts of inventory types based on one or more expressions. For example,
	// if you aggregate by using an expression that uses the
	// AWS:InstanceInformation.PlatformType type, you can see a count of how many
	// Windows and Linux instances exist in your inventoried fleet.
	Aggregators []types.InventoryAggregator

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []types.InventoryFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The list of inventory item types to return.
	ResultAttributes []types.ResultAttribute
}

type GetInventoryOutput struct {

	// Collection of inventory entities such as a collection of instance inventory.
	Entities []types.InventoryResultEntity

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInventoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInventory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInventory{}, middleware.After)
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
	if err = addOpGetInventoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInventory(options.Region), middleware.Before); err != nil {
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

// GetInventoryAPIClient is a client that implements the GetInventory operation.
type GetInventoryAPIClient interface {
	GetInventory(context.Context, *GetInventoryInput, ...func(*Options)) (*GetInventoryOutput, error)
}

var _ GetInventoryAPIClient = (*Client)(nil)

// GetInventoryPaginatorOptions is the paginator options for GetInventory
type GetInventoryPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInventoryPaginator is a paginator for GetInventory
type GetInventoryPaginator struct {
	options   GetInventoryPaginatorOptions
	client    GetInventoryAPIClient
	params    *GetInventoryInput
	nextToken *string
	firstPage bool
}

// NewGetInventoryPaginator returns a new GetInventoryPaginator
func NewGetInventoryPaginator(client GetInventoryAPIClient, params *GetInventoryInput, optFns ...func(*GetInventoryPaginatorOptions)) *GetInventoryPaginator {
	options := GetInventoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetInventoryInput{}
	}

	return &GetInventoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInventoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetInventory page.
func (p *GetInventoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInventoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetInventory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInventory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetInventory",
	}
}
