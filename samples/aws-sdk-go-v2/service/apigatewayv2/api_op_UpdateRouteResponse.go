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

// Updates a RouteResponse.
func (c *Client) UpdateRouteResponse(ctx context.Context, params *UpdateRouteResponseInput, optFns ...func(*Options)) (*UpdateRouteResponseOutput, error) {
	if params == nil {
		params = &UpdateRouteResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRouteResponse", params, optFns, addOperationUpdateRouteResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRouteResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a RouteResponse.
type UpdateRouteResponseInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The route ID.
	//
	// This member is required.
	RouteId *string

	// The route response ID.
	//
	// This member is required.
	RouteResponseId *string

	// The model selection expression for the route response. Supported only for
	// WebSocket APIs.
	ModelSelectionExpression *string

	// The response models for the route response.
	ResponseModels map[string]string

	// The route response parameters.
	ResponseParameters map[string]types.ParameterConstraints

	// The route response key.
	RouteResponseKey *string
}

type UpdateRouteResponseOutput struct {

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

func addOperationUpdateRouteResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRouteResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRouteResponse{}, middleware.After)
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
	if err = addOpUpdateRouteResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRouteResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRouteResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateRouteResponse",
	}
}
