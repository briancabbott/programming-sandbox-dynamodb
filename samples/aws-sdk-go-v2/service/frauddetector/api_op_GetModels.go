// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets one or more models. Gets all models for the AWS account if no model type
// and no model id provided. Gets all models for the AWS account and model type, if
// the model type is specified but model id is not provided. Gets a specific model
// if (model type, model id) tuple is specified. This is a paginated API. If you
// provide a null maxResults, this action retrieves a maximum of 10 records per
// page. If you provide a maxResults, the value must be between 1 and 10. To get
// the next page results, provide the pagination token from the response as part of
// your request. A null pagination token fetches the records from the beginning.
func (c *Client) GetModels(ctx context.Context, params *GetModelsInput, optFns ...func(*Options)) (*GetModelsOutput, error) {
	if params == nil {
		params = &GetModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModels", params, optFns, addOperationGetModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelsInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The model ID.
	ModelId *string

	// The model type.
	ModelType types.ModelTypeEnum

	// The next token for the subsequent request.
	NextToken *string
}

type GetModelsOutput struct {

	// The array of models.
	Models []types.Model

	// The next page token to be used in subsequent requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModels(options.Region), middleware.Before); err != nil {
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

// GetModelsAPIClient is a client that implements the GetModels operation.
type GetModelsAPIClient interface {
	GetModels(context.Context, *GetModelsInput, ...func(*Options)) (*GetModelsOutput, error)
}

var _ GetModelsAPIClient = (*Client)(nil)

// GetModelsPaginatorOptions is the paginator options for GetModels
type GetModelsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetModelsPaginator is a paginator for GetModels
type GetModelsPaginator struct {
	options   GetModelsPaginatorOptions
	client    GetModelsAPIClient
	params    *GetModelsInput
	nextToken *string
	firstPage bool
}

// NewGetModelsPaginator returns a new GetModelsPaginator
func NewGetModelsPaginator(client GetModelsAPIClient, params *GetModelsInput, optFns ...func(*GetModelsPaginatorOptions)) *GetModelsPaginator {
	options := GetModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetModelsInput{}
	}

	return &GetModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetModelsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetModels page.
func (p *GetModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetModelsOutput, error) {
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

	result, err := p.client.GetModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetModels",
	}
}
