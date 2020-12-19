// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Start (run) the multiplex. Starting the multiplex does not start the channels.
// You must explicitly start each channel.
func (c *Client) StartMultiplex(ctx context.Context, params *StartMultiplexInput, optFns ...func(*Options)) (*StartMultiplexOutput, error) {
	if params == nil {
		params = &StartMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMultiplex", params, optFns, addOperationStartMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StartMultiplexRequest
type StartMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string
}

// Placeholder documentation for StartMultiplexResponse
type StartMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []string

	// A list of the multiplex output destinations.
	Destinations []types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount int32

	// The number of programs in the multiplex.
	ProgramCount int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMultiplex{}, middleware.After)
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
	if err = addOpStartMultiplexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMultiplex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMultiplex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "StartMultiplex",
	}
}
