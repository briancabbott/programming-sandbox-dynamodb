// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by activity workers and task states using the callback
// (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern to report to Step Functions that the task represented by the specified
// taskToken is still making progress. This action resets the Heartbeat clock. The
// Heartbeat threshold is specified in the state machine's Amazon States Language
// definition (HeartbeatSeconds). This action does not in itself create an event in
// the execution history. However, if the task times out, the execution history
// contains an ActivityTimedOut entry for activities, or a TaskTimedOut entry for
// for tasks using the job run
// (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync)
// or callback
// (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern. The Timeout of a task, defined in the state machine's Amazon States
// Language definition, is its maximum allowed duration, regardless of the number
// of SendTaskHeartbeat requests received. Use HeartbeatSeconds to configure the
// timeout interval for heartbeats.
func (c *Client) SendTaskHeartbeat(ctx context.Context, params *SendTaskHeartbeatInput, optFns ...func(*Options)) (*SendTaskHeartbeatOutput, error) {
	if params == nil {
		params = &SendTaskHeartbeatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendTaskHeartbeat", params, optFns, addOperationSendTaskHeartbeatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendTaskHeartbeatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTaskHeartbeatInput struct {

	// The token that represents this task. Task tokens are generated by Step Functions
	// when tasks are assigned to a worker, or in the context object
	// (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html)
	// when a workflow enters a task state. See GetActivityTaskOutput$taskToken.
	//
	// This member is required.
	TaskToken *string
}

type SendTaskHeartbeatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendTaskHeartbeatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendTaskHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendTaskHeartbeat{}, middleware.After)
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
	if err = addOpSendTaskHeartbeatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendTaskHeartbeat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendTaskHeartbeat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "SendTaskHeartbeat",
	}
}
