// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about projects.
func (c *Client) ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) {
	if params == nil {
		params = &ListProjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProjects", params, optFns, addOperationListProjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list projects operation.
type ListProjectsInput struct {

	// Optional. If no Amazon Resource Name (ARN) is specified, then AWS Device Farm
	// returns a list of all projects for the AWS account. You can also specify a
	// project ARN.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list projects request.
type ListProjectsOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the projects.
	Projects []types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProjects{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProjects(options.Region), middleware.Before); err != nil {
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

// ListProjectsAPIClient is a client that implements the ListProjects operation.
type ListProjectsAPIClient interface {
	ListProjects(context.Context, *ListProjectsInput, ...func(*Options)) (*ListProjectsOutput, error)
}

var _ ListProjectsAPIClient = (*Client)(nil)

// ListProjectsPaginatorOptions is the paginator options for ListProjects
type ListProjectsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProjectsPaginator is a paginator for ListProjects
type ListProjectsPaginator struct {
	options   ListProjectsPaginatorOptions
	client    ListProjectsAPIClient
	params    *ListProjectsInput
	nextToken *string
	firstPage bool
}

// NewListProjectsPaginator returns a new ListProjectsPaginator
func NewListProjectsPaginator(client ListProjectsAPIClient, params *ListProjectsInput, optFns ...func(*ListProjectsPaginatorOptions)) *ListProjectsPaginator {
	options := ListProjectsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListProjectsInput{}
	}

	return &ListProjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProjectsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListProjects page.
func (p *ListProjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProjectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListProjects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListProjects",
	}
}
