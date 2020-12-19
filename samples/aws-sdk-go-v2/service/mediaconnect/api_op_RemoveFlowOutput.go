// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an output from an existing flow. This request can be made only on an
// output that does not have an entitlement associated with it. If the output has
// an entitlement, you must revoke the entitlement instead. When an entitlement is
// revoked from a flow, the service automatically removes the associated output.
func (c *Client) RemoveFlowOutput(ctx context.Context, params *RemoveFlowOutputInput, optFns ...func(*Options)) (*RemoveFlowOutputOutput, error) {
	if params == nil {
		params = &RemoveFlowOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveFlowOutput", params, optFns, addOperationRemoveFlowOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveFlowOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveFlowOutputInput struct {

	// The flow that you want to remove an output from.
	//
	// This member is required.
	FlowArn *string

	// The ARN of the output that you want to remove.
	//
	// This member is required.
	OutputArn *string
}

type RemoveFlowOutputOutput struct {

	// The ARN of the flow that is associated with the output you removed.
	FlowArn *string

	// The ARN of the output that was removed.
	OutputArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveFlowOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveFlowOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveFlowOutput{}, middleware.After)
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
	if err = addOpRemoveFlowOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveFlowOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveFlowOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "RemoveFlowOutput",
	}
}
