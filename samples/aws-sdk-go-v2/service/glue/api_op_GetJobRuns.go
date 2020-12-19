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

// Retrieves metadata for all runs of a given job definition.
func (c *Client) GetJobRuns(ctx context.Context, params *GetJobRunsInput, optFns ...func(*Options)) (*GetJobRunsOutput, error) {
	if params == nil {
		params = &GetJobRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJobRuns", params, optFns, addOperationGetJobRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobRunsInput struct {

	// The name of the job definition for which to retrieve all job runs.
	//
	// This member is required.
	JobName *string

	// The maximum size of the response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetJobRunsOutput struct {

	// A list of job-run metadata objects.
	JobRuns []types.JobRun

	// A continuation token, if not all requested job runs have been returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetJobRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobRuns{}, middleware.After)
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
	if err = addOpGetJobRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobRuns(options.Region), middleware.Before); err != nil {
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

// GetJobRunsAPIClient is a client that implements the GetJobRuns operation.
type GetJobRunsAPIClient interface {
	GetJobRuns(context.Context, *GetJobRunsInput, ...func(*Options)) (*GetJobRunsOutput, error)
}

var _ GetJobRunsAPIClient = (*Client)(nil)

// GetJobRunsPaginatorOptions is the paginator options for GetJobRuns
type GetJobRunsPaginatorOptions struct {
	// The maximum size of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetJobRunsPaginator is a paginator for GetJobRuns
type GetJobRunsPaginator struct {
	options   GetJobRunsPaginatorOptions
	client    GetJobRunsAPIClient
	params    *GetJobRunsInput
	nextToken *string
	firstPage bool
}

// NewGetJobRunsPaginator returns a new GetJobRunsPaginator
func NewGetJobRunsPaginator(client GetJobRunsAPIClient, params *GetJobRunsInput, optFns ...func(*GetJobRunsPaginatorOptions)) *GetJobRunsPaginator {
	options := GetJobRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetJobRunsInput{}
	}

	return &GetJobRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetJobRunsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetJobRuns page.
func (p *GetJobRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetJobRunsOutput, error) {
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

	result, err := p.client.GetJobRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetJobRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetJobRuns",
	}
}
