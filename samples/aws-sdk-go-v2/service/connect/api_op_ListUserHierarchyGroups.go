// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides summary information about the hierarchy groups for the specified Amazon
// Connect instance. For more information about agent hierarchies, see Set Up Agent
// Hierarchies
// (https://docs.aws.amazon.com/connect/latest/adminguide/agent-hierarchy.html) in
// the Amazon Connect Administrator Guide.
func (c *Client) ListUserHierarchyGroups(ctx context.Context, params *ListUserHierarchyGroupsInput, optFns ...func(*Options)) (*ListUserHierarchyGroupsOutput, error) {
	if params == nil {
		params = &ListUserHierarchyGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserHierarchyGroups", params, optFns, addOperationListUserHierarchyGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserHierarchyGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserHierarchyGroupsInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type ListUserHierarchyGroupsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the hierarchy groups.
	UserHierarchyGroupSummaryList []types.HierarchyGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUserHierarchyGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUserHierarchyGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserHierarchyGroups{}, middleware.After)
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
	if err = addOpListUserHierarchyGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserHierarchyGroups(options.Region), middleware.Before); err != nil {
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

// ListUserHierarchyGroupsAPIClient is a client that implements the
// ListUserHierarchyGroups operation.
type ListUserHierarchyGroupsAPIClient interface {
	ListUserHierarchyGroups(context.Context, *ListUserHierarchyGroupsInput, ...func(*Options)) (*ListUserHierarchyGroupsOutput, error)
}

var _ ListUserHierarchyGroupsAPIClient = (*Client)(nil)

// ListUserHierarchyGroupsPaginatorOptions is the paginator options for
// ListUserHierarchyGroups
type ListUserHierarchyGroupsPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserHierarchyGroupsPaginator is a paginator for ListUserHierarchyGroups
type ListUserHierarchyGroupsPaginator struct {
	options   ListUserHierarchyGroupsPaginatorOptions
	client    ListUserHierarchyGroupsAPIClient
	params    *ListUserHierarchyGroupsInput
	nextToken *string
	firstPage bool
}

// NewListUserHierarchyGroupsPaginator returns a new
// ListUserHierarchyGroupsPaginator
func NewListUserHierarchyGroupsPaginator(client ListUserHierarchyGroupsAPIClient, params *ListUserHierarchyGroupsInput, optFns ...func(*ListUserHierarchyGroupsPaginatorOptions)) *ListUserHierarchyGroupsPaginator {
	options := ListUserHierarchyGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListUserHierarchyGroupsInput{}
	}

	return &ListUserHierarchyGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserHierarchyGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUserHierarchyGroups page.
func (p *ListUserHierarchyGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserHierarchyGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListUserHierarchyGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserHierarchyGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListUserHierarchyGroups",
	}
}
