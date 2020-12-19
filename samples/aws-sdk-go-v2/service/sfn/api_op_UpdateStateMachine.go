// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates an existing state machine by modifying its definition, roleArn, or
// loggingConfiguration. Running executions will continue to use the previous
// definition and roleArn. You must include at least one of definition or roleArn
// or you will receive a MissingRequiredParameter error. All StartExecution calls
// within a few seconds will use the updated definition and roleArn. Executions
// started immediately after calling UpdateStateMachine may use the previous state
// machine definition and roleArn.
func (c *Client) UpdateStateMachine(ctx context.Context, params *UpdateStateMachineInput, optFns ...func(*Options)) (*UpdateStateMachineOutput, error) {
	if params == nil {
		params = &UpdateStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStateMachine", params, optFns, addOperationUpdateStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language
	// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	Definition *string

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// The Amazon Resource Name (ARN) of the IAM role of the state machine.
	RoleArn *string

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration
}

type UpdateStateMachineOutput struct {

	// The date and time the state machine was updated.
	//
	// This member is required.
	UpdateDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateStateMachine{}, middleware.After)
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
	if err = addOpUpdateStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "UpdateStateMachine",
	}
}