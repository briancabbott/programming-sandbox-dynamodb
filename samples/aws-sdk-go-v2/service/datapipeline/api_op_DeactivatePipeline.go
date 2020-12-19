// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deactivates the specified running pipeline. The pipeline is set to the
// DEACTIVATING state until the deactivation process completes. To resume a
// deactivated pipeline, use ActivatePipeline. By default, the pipeline resumes
// from the last completed execution. Optionally, you can specify the date and time
// to resume the pipeline.
func (c *Client) DeactivatePipeline(ctx context.Context, params *DeactivatePipelineInput, optFns ...func(*Options)) (*DeactivatePipelineOutput, error) {
	if params == nil {
		params = &DeactivatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivatePipeline", params, optFns, addOperationDeactivatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeactivatePipeline.
type DeactivatePipelineInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether to cancel any running objects. The default is true, which sets
	// the state of any running objects to CANCELED. If this value is false, the
	// pipeline is deactivated after all running objects finish.
	CancelActive *bool
}

// Contains the output of DeactivatePipeline.
type DeactivatePipelineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeactivatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeactivatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeactivatePipeline{}, middleware.After)
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
	if err = addOpDeactivatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeactivatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DeactivatePipeline",
	}
}
