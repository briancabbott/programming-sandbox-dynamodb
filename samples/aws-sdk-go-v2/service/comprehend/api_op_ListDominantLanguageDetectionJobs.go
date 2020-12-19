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

// Gets a list of the dominant language detection jobs that you have submitted.
func (c *Client) ListDominantLanguageDetectionJobs(ctx context.Context, params *ListDominantLanguageDetectionJobsInput, optFns ...func(*Options)) (*ListDominantLanguageDetectionJobsOutput, error) {
	if params == nil {
		params = &ListDominantLanguageDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDominantLanguageDetectionJobs", params, optFns, addOperationListDominantLanguageDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDominantLanguageDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDominantLanguageDetectionJobsInput struct {

	// Filters that jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.DominantLanguageDetectionJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListDominantLanguageDetectionJobsOutput struct {

	// A list containing the properties of each job that is returned.
	DominantLanguageDetectionJobPropertiesList []types.DominantLanguageDetectionJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDominantLanguageDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDominantLanguageDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDominantLanguageDetectionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDominantLanguageDetectionJobs(options.Region), middleware.Before); err != nil {
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

// ListDominantLanguageDetectionJobsAPIClient is a client that implements the
// ListDominantLanguageDetectionJobs operation.
type ListDominantLanguageDetectionJobsAPIClient interface {
	ListDominantLanguageDetectionJobs(context.Context, *ListDominantLanguageDetectionJobsInput, ...func(*Options)) (*ListDominantLanguageDetectionJobsOutput, error)
}

var _ ListDominantLanguageDetectionJobsAPIClient = (*Client)(nil)

// ListDominantLanguageDetectionJobsPaginatorOptions is the paginator options for
// ListDominantLanguageDetectionJobs
type ListDominantLanguageDetectionJobsPaginatorOptions struct {
	// The maximum number of results to return in each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDominantLanguageDetectionJobsPaginator is a paginator for
// ListDominantLanguageDetectionJobs
type ListDominantLanguageDetectionJobsPaginator struct {
	options   ListDominantLanguageDetectionJobsPaginatorOptions
	client    ListDominantLanguageDetectionJobsAPIClient
	params    *ListDominantLanguageDetectionJobsInput
	nextToken *string
	firstPage bool
}

// NewListDominantLanguageDetectionJobsPaginator returns a new
// ListDominantLanguageDetectionJobsPaginator
func NewListDominantLanguageDetectionJobsPaginator(client ListDominantLanguageDetectionJobsAPIClient, params *ListDominantLanguageDetectionJobsInput, optFns ...func(*ListDominantLanguageDetectionJobsPaginatorOptions)) *ListDominantLanguageDetectionJobsPaginator {
	options := ListDominantLanguageDetectionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDominantLanguageDetectionJobsInput{}
	}

	return &ListDominantLanguageDetectionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDominantLanguageDetectionJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDominantLanguageDetectionJobs page.
func (p *ListDominantLanguageDetectionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDominantLanguageDetectionJobsOutput, error) {
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

	result, err := p.client.ListDominantLanguageDetectionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDominantLanguageDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListDominantLanguageDetectionJobs",
	}
}
