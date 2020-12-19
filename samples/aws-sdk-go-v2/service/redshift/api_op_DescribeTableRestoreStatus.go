// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the status of one or more table restore requests made using the
// RestoreTableFromClusterSnapshot API action. If you don't specify a value for the
// TableRestoreRequestId parameter, then DescribeTableRestoreStatus returns the
// status of all table restore requests ordered by the date and time of the request
// in ascending order. Otherwise DescribeTableRestoreStatus returns the status of
// the table specified by TableRestoreRequestId.
func (c *Client) DescribeTableRestoreStatus(ctx context.Context, params *DescribeTableRestoreStatusInput, optFns ...func(*Options)) (*DescribeTableRestoreStatusOutput, error) {
	if params == nil {
		params = &DescribeTableRestoreStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTableRestoreStatus", params, optFns, addOperationDescribeTableRestoreStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableRestoreStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeTableRestoreStatusInput struct {

	// The Amazon Redshift cluster that the table is being restored to.
	ClusterIdentifier *string

	// An optional pagination token provided by a previous DescribeTableRestoreStatus
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	// The identifier of the table restore request to return status for. If you don't
	// specify a TableRestoreRequestId value, then DescribeTableRestoreStatus returns
	// the status of all in-progress table restore requests.
	TableRestoreRequestId *string
}

//
type DescribeTableRestoreStatusOutput struct {

	// A pagination token that can be used in a subsequent DescribeTableRestoreStatus
	// request.
	Marker *string

	// A list of status details for one or more table restore requests.
	TableRestoreStatusDetails []types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTableRestoreStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTableRestoreStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTableRestoreStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTableRestoreStatus(options.Region), middleware.Before); err != nil {
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

// DescribeTableRestoreStatusAPIClient is a client that implements the
// DescribeTableRestoreStatus operation.
type DescribeTableRestoreStatusAPIClient interface {
	DescribeTableRestoreStatus(context.Context, *DescribeTableRestoreStatusInput, ...func(*Options)) (*DescribeTableRestoreStatusOutput, error)
}

var _ DescribeTableRestoreStatusAPIClient = (*Client)(nil)

// DescribeTableRestoreStatusPaginatorOptions is the paginator options for
// DescribeTableRestoreStatus
type DescribeTableRestoreStatusPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTableRestoreStatusPaginator is a paginator for
// DescribeTableRestoreStatus
type DescribeTableRestoreStatusPaginator struct {
	options   DescribeTableRestoreStatusPaginatorOptions
	client    DescribeTableRestoreStatusAPIClient
	params    *DescribeTableRestoreStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeTableRestoreStatusPaginator returns a new
// DescribeTableRestoreStatusPaginator
func NewDescribeTableRestoreStatusPaginator(client DescribeTableRestoreStatusAPIClient, params *DescribeTableRestoreStatusInput, optFns ...func(*DescribeTableRestoreStatusPaginatorOptions)) *DescribeTableRestoreStatusPaginator {
	options := DescribeTableRestoreStatusPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeTableRestoreStatusInput{}
	}

	return &DescribeTableRestoreStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTableRestoreStatusPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeTableRestoreStatus page.
func (p *DescribeTableRestoreStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTableRestoreStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeTableRestoreStatus(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTableRestoreStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeTableRestoreStatus",
	}
}
