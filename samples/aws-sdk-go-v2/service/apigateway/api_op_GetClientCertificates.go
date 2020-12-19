// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a collection of ClientCertificate resources.
func (c *Client) GetClientCertificates(ctx context.Context, params *GetClientCertificatesInput, optFns ...func(*Options)) (*GetClientCertificatesOutput, error) {
	if params == nil {
		params = &GetClientCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClientCertificates", params, optFns, addOperationGetClientCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClientCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a collection of ClientCertificate resources.
type GetClientCertificatesInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// Represents a collection of ClientCertificate resources. Use Client-Side
// Certificate
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GetClientCertificatesOutput struct {

	// The current page of elements from this collection.
	Items []types.ClientCertificate

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetClientCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetClientCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetClientCertificates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClientCertificates(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetClientCertificatesAPIClient is a client that implements the
// GetClientCertificates operation.
type GetClientCertificatesAPIClient interface {
	GetClientCertificates(context.Context, *GetClientCertificatesInput, ...func(*Options)) (*GetClientCertificatesOutput, error)
}

var _ GetClientCertificatesAPIClient = (*Client)(nil)

// GetClientCertificatesPaginatorOptions is the paginator options for
// GetClientCertificates
type GetClientCertificatesPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetClientCertificatesPaginator is a paginator for GetClientCertificates
type GetClientCertificatesPaginator struct {
	options   GetClientCertificatesPaginatorOptions
	client    GetClientCertificatesAPIClient
	params    *GetClientCertificatesInput
	nextToken *string
	firstPage bool
}

// NewGetClientCertificatesPaginator returns a new GetClientCertificatesPaginator
func NewGetClientCertificatesPaginator(client GetClientCertificatesAPIClient, params *GetClientCertificatesInput, optFns ...func(*GetClientCertificatesPaginatorOptions)) *GetClientCertificatesPaginator {
	options := GetClientCertificatesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetClientCertificatesInput{}
	}

	return &GetClientCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetClientCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetClientCertificates page.
func (p *GetClientCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetClientCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetClientCertificates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetClientCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetClientCertificates",
	}
}
