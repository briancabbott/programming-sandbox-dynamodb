// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the specified flow definition.
func (c *Client) DescribeFlowDefinition(ctx context.Context, params *DescribeFlowDefinitionInput, optFns ...func(*Options)) (*DescribeFlowDefinitionOutput, error) {
	if params == nil {
		params = &DescribeFlowDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlowDefinition", params, optFns, addOperationDescribeFlowDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowDefinitionInput struct {

	// The name of the flow definition.
	//
	// This member is required.
	FlowDefinitionName *string
}

type DescribeFlowDefinitionOutput struct {

	// The timestamp when the flow definition was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the flow defintion.
	//
	// This member is required.
	FlowDefinitionArn *string

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// This member is required.
	FlowDefinitionName *string

	// The status of the flow definition. Valid values are listed below.
	//
	// This member is required.
	FlowDefinitionStatus types.FlowDefinitionStatus

	// An object containing information about who works on the task, the workforce task
	// price, and other task details.
	//
	// This member is required.
	HumanLoopConfig *types.HumanLoopConfig

	// An object containing information about the output file.
	//
	// This member is required.
	OutputConfig *types.FlowDefinitionOutputConfig

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// execution role for the flow definition.
	//
	// This member is required.
	RoleArn *string

	// The reason your flow definition failed.
	FailureReason *string

	// An object containing information about what triggers a human review workflow.
	HumanLoopActivationConfig *types.HumanLoopActivationConfig

	// Container for configuring the source of human task requests. Used to specify if
	// Amazon Rekognition or Amazon Textract is used as an integration source.
	HumanLoopRequestSource *types.HumanLoopRequestSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFlowDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFlowDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFlowDefinition{}, middleware.After)
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
	if err = addOpDescribeFlowDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlowDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFlowDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeFlowDefinition",
	}
}
