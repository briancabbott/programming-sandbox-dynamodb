// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the entity detection jobs that you have submitted.
func (c *Client) ListEntitiesDetectionJobs(ctx context.Context, params *ListEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionJobsOutput, error) {
	if params == nil {
		params = &ListEntitiesDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitiesDetectionJobs", params, optFns, addOperationListEntitiesDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.EntitiesDetectionJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListEntitiesDetectionJobsOutput struct {

	// A list containing the properties of each job that is returned.
	EntitiesDetectionJobPropertiesList []types.EntitiesDetectionJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntitiesDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntitiesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntitiesDetectionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesDetectionJobs(options.Region), middleware.Before); err != nil {
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

// ListEntitiesDetectionJobsAPIClient is a client that implements the
// ListEntitiesDetectionJobs operation.
type ListEntitiesDetectionJobsAPIClient interface {
	ListEntitiesDetectionJobs(context.Context, *ListEntitiesDetectionJobsInput, ...func(*Options)) (*ListEntitiesDetectionJobsOutput, error)
}

var _ ListEntitiesDetectionJobsAPIClient = (*Client)(nil)

// ListEntitiesDetectionJobsPaginatorOptions is the paginator options for
// ListEntitiesDetectionJobs
type ListEntitiesDetectionJobsPaginatorOptions struct {
	// The maximum number of results to return in each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntitiesDetectionJobsPaginator is a paginator for ListEntitiesDetectionJobs
type ListEntitiesDetectionJobsPaginator struct {
	options   ListEntitiesDetectionJobsPaginatorOptions
	client    ListEntitiesDetectionJobsAPIClient
	params    *ListEntitiesDetectionJobsInput
	nextToken *string
	firstPage bool
}

// NewListEntitiesDetectionJobsPaginator returns a new
// ListEntitiesDetectionJobsPaginator
func NewListEntitiesDetectionJobsPaginator(client ListEntitiesDetectionJobsAPIClient, params *ListEntitiesDetectionJobsInput, optFns ...func(*ListEntitiesDetectionJobsPaginatorOptions)) *ListEntitiesDetectionJobsPaginator {
	options := ListEntitiesDetectionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListEntitiesDetectionJobsInput{}
	}

	return &ListEntitiesDetectionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntitiesDetectionJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListEntitiesDetectionJobs page.
func (p *ListEntitiesDetectionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntitiesDetectionJobsOutput, error) {
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

	result, err := p.client.ListEntitiesDetectionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEntitiesDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListEntitiesDetectionJobs",
	}
}
