// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a RouteResponse for a Route.
func (c *Client) CreateRouteResponse(ctx context.Context, params *CreateRouteResponseInput, optFns ...func(*Options)) (*CreateRouteResponseOutput, error) {
	if params == nil {
		params = &CreateRouteResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRouteResponse", params, optFns, addOperationCreateRouteResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new RouteResponse resource to represent a route response.
type CreateRouteResponseInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The route ID.
	//
	// This member is required.
	RouteId *string

	// The route response key.
	//
	// This member is required.
	RouteResponseKey *string

	// The model selection expression for the route response. Supported only for
	// WebSocket APIs.
	ModelSelectionExpression *string

	// The response models for the route response.
	ResponseModels map[string]string

	// The route response parameters.
	ResponseParameters map[string]types.ParameterConstraints
}

type CreateRouteResponseOutput struct {

	// Represents the model selection expression of a route response. Supported only
	// for WebSocket APIs.
	ModelSelectionExpression *string

	// Represents the response models of a route response.
	ResponseModels map[string]string

	// Represents the response parameters of a route response.
	ResponseParameters map[string]types.ParameterConstraints

	// Represents the identifier of a route response.
	RouteResponseId *string

	// Represents the route response key of a route response.
	RouteResponseKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRouteResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRouteResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRouteResponse{}, middleware.After)
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
	if err = addOpCreateRouteResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRouteResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRouteResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRouteResponse",
	}
}
