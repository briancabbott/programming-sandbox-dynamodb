// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all ongoing DDoS attacks or all DDoS attacks during a specified time
// period.
func (c *Client) ListAttacks(ctx context.Context, params *ListAttacksInput, optFns ...func(*Options)) (*ListAttacksOutput, error) {
	if params == nil {
		params = &ListAttacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttacks", params, optFns, addOperationListAttacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttacksInput struct {

	// The end of the time period for the attacks. This is a timestamp type. The sample
	// request above indicates a number type because the default used by WAF is Unix
	// time in seconds. However any valid timestamp format
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	EndTime *types.TimeRange

	// The maximum number of AttackSummary objects to return. If you leave this blank,
	// Shield Advanced returns the first 20 results. This is a maximum value. Shield
	// Advanced might return the results in smaller batches. That is, the number of
	// objects returned could be less than MaxResults, even if there are still more
	// objects yet to return. If there are more objects to return, Shield Advanced
	// returns a value in NextToken that you can use in your next request, to get the
	// next batch of objects.
	MaxResults *int32

	// The ListAttacksRequest.NextMarker value from a previous call to
	// ListAttacksRequest. Pass null if this is the first call.
	NextToken *string

	// The ARN (Amazon Resource Name) of the resource that was attacked. If this is
	// left blank, all applicable resources for this account will be included.
	ResourceArns []string

	// The start of the time period for the attacks. This is a timestamp type. The
	// sample request above indicates a number type because the default used by WAF is
	// Unix time in seconds. However any valid timestamp format
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	StartTime *types.TimeRange
}

type ListAttacksOutput struct {

	// The attack information for the specified time range.
	AttackSummaries []types.AttackSummary

	// The token returned by a previous call to indicate that there is more data
	// available. If not null, more results are available. Pass this value for the
	// NextMarker parameter in a subsequent call to ListAttacks to retrieve the next
	// set of items. Shield Advanced might return the list of AttackSummary objects in
	// batches smaller than the number specified by MaxResults. If there are more
	// attack summary objects to return, Shield Advanced will always also return a
	// NextToken.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAttacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAttacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAttacks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttacks(options.Region), middleware.Before); err != nil {
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

// ListAttacksAPIClient is a client that implements the ListAttacks operation.
type ListAttacksAPIClient interface {
	ListAttacks(context.Context, *ListAttacksInput, ...func(*Options)) (*ListAttacksOutput, error)
}

var _ ListAttacksAPIClient = (*Client)(nil)

// ListAttacksPaginatorOptions is the paginator options for ListAttacks
type ListAttacksPaginatorOptions struct {
	// The maximum number of AttackSummary objects to return. If you leave this blank,
	// Shield Advanced returns the first 20 results. This is a maximum value. Shield
	// Advanced might return the results in smaller batches. That is, the number of
	// objects returned could be less than MaxResults, even if there are still more
	// objects yet to return. If there are more objects to return, Shield Advanced
	// returns a value in NextToken that you can use in your next request, to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttacksPaginator is a paginator for ListAttacks
type ListAttacksPaginator struct {
	options   ListAttacksPaginatorOptions
	client    ListAttacksAPIClient
	params    *ListAttacksInput
	nextToken *string
	firstPage bool
}

// NewListAttacksPaginator returns a new ListAttacksPaginator
func NewListAttacksPaginator(client ListAttacksAPIClient, params *ListAttacksInput, optFns ...func(*ListAttacksPaginatorOptions)) *ListAttacksPaginator {
	options := ListAttacksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListAttacksInput{}
	}

	return &ListAttacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttacksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAttacks page.
func (p *ListAttacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttacksOutput, error) {
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

	result, err := p.client.ListAttacks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListAttacks",
	}
}
