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

// Returns all Systems Manager (SSM) documents in the current AWS account and
// Region. You can limit the results of this request by using a filter.
func (c *Client) ListDocuments(ctx context.Context, params *ListDocumentsInput, optFns ...func(*Options)) (*ListDocumentsOutput, error) {
	if params == nil {
		params = &ListDocumentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocuments", params, optFns, addOperationListDocumentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentsInput struct {

	// This data type is deprecated. Instead, use Filters.
	DocumentFilterList []types.DocumentFilter

	// One or more DocumentKeyValuesFilter objects. Use a filter to return a more
	// specific list of results. For keys, you can specify one or more key-value pair
	// tags that have been applied to a document. Other valid keys include Owner, Name,
	// PlatformTypes, DocumentType, and TargetType. For example, to return documents
	// you own use Key=Owner,Values=Self. To specify a custom key-value pair, use the
	// format Key=tag:tagName,Values=valueName.
	Filters []types.DocumentKeyValuesFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type ListDocumentsOutput struct {

	// The names of the Systems Manager documents.
	DocumentIdentifiers []types.DocumentIdentifier

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDocumentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocuments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocuments{}, middleware.After)
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
	if err = addOpListDocumentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocuments(options.Region), middleware.Before); err != nil {
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

// ListDocumentsAPIClient is a client that implements the ListDocuments operation.
type ListDocumentsAPIClient interface {
	ListDocuments(context.Context, *ListDocumentsInput, ...func(*Options)) (*ListDocumentsOutput, error)
}

var _ ListDocumentsAPIClient = (*Client)(nil)

// ListDocumentsPaginatorOptions is the paginator options for ListDocuments
type ListDocumentsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDocumentsPaginator is a paginator for ListDocuments
type ListDocumentsPaginator struct {
	options   ListDocumentsPaginatorOptions
	client    ListDocumentsAPIClient
	params    *ListDocumentsInput
	nextToken *string
	firstPage bool
}

// NewListDocumentsPaginator returns a new ListDocumentsPaginator
func NewListDocumentsPaginator(client ListDocumentsAPIClient, params *ListDocumentsInput, optFns ...func(*ListDocumentsPaginatorOptions)) *ListDocumentsPaginator {
	options := ListDocumentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDocumentsInput{}
	}

	return &ListDocumentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDocumentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDocuments page.
func (p *ListDocumentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDocumentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDocuments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDocuments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListDocuments",
	}
}
