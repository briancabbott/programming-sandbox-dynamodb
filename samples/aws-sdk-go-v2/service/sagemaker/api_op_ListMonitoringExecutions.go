// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns list of all monitoring job executions.
func (c *Client) ListMonitoringExecutions(ctx context.Context, params *ListMonitoringExecutionsInput, optFns ...func(*Options)) (*ListMonitoringExecutionsOutput, error) {
	if params == nil {
		params = &ListMonitoringExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitoringExecutions", params, optFns, addOperationListMonitoringExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitoringExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitoringExecutionsInput struct {

	// A filter that returns only jobs created after a specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only jobs created before a specified time.
	CreationTimeBefore *time.Time

	// Name of a specific endpoint to fetch jobs for.
	EndpointName *string

	// A filter that returns only jobs modified before a specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only jobs modified after a specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of jobs to return in the response. The default value is 10.
	MaxResults *int32

	// Name of a specific schedule to fetch jobs for.
	MonitoringScheduleName *string

	// The token returned if the response is truncated. To retrieve the next set of job
	// executions, use it in the next request.
	NextToken *string

	// Filter for jobs scheduled after a specified time.
	ScheduledTimeAfter *time.Time

	// Filter for jobs scheduled before a specified time.
	ScheduledTimeBefore *time.Time

	// Whether to sort results by Status, CreationTime, ScheduledTime field. The
	// default is CreationTime.
	SortBy types.MonitoringExecutionSortKey

	// Whether to sort the results in Ascending or Descending order. The default is
	// Descending.
	SortOrder types.SortOrder

	// A filter that retrieves only jobs with a specific status.
	StatusEquals types.ExecutionStatus
}

type ListMonitoringExecutionsOutput struct {

	// A JSON array in which each element is a summary for a monitoring execution.
	//
	// This member is required.
	MonitoringExecutionSummaries []types.MonitoringExecutionSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of jobs, use it in the subsequent reques
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMonitoringExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitoringExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitoringExecutions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitoringExecutions(options.Region), middleware.Before); err != nil {
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

// ListMonitoringExecutionsAPIClient is a client that implements the
// ListMonitoringExecutions operation.
type ListMonitoringExecutionsAPIClient interface {
	ListMonitoringExecutions(context.Context, *ListMonitoringExecutionsInput, ...func(*Options)) (*ListMonitoringExecutionsOutput, error)
}

var _ ListMonitoringExecutionsAPIClient = (*Client)(nil)

// ListMonitoringExecutionsPaginatorOptions is the paginator options for
// ListMonitoringExecutions
type ListMonitoringExecutionsPaginatorOptions struct {
	// The maximum number of jobs to return in the response. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitoringExecutionsPaginator is a paginator for ListMonitoringExecutions
type ListMonitoringExecutionsPaginator struct {
	options   ListMonitoringExecutionsPaginatorOptions
	client    ListMonitoringExecutionsAPIClient
	params    *ListMonitoringExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListMonitoringExecutionsPaginator returns a new
// ListMonitoringExecutionsPaginator
func NewListMonitoringExecutionsPaginator(client ListMonitoringExecutionsAPIClient, params *ListMonitoringExecutionsInput, optFns ...func(*ListMonitoringExecutionsPaginatorOptions)) *ListMonitoringExecutionsPaginator {
	options := ListMonitoringExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListMonitoringExecutionsInput{}
	}

	return &ListMonitoringExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitoringExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListMonitoringExecutions page.
func (p *ListMonitoringExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitoringExecutionsOutput, error) {
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

	result, err := p.client.ListMonitoringExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitoringExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListMonitoringExecutions",
	}
}
