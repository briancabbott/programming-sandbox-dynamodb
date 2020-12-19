// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a flow definition.
func (c *Client) CreateFlowDefinition(ctx context.Context, params *CreateFlowDefinitionInput, optFns ...func(*Options)) (*CreateFlowDefinitionOutput, error) {
	if params == nil {
		params = &CreateFlowDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlowDefinition", params, optFns, addOperationCreateFlowDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowDefinitionInput struct {

	// The name of your flow definition.
	//
	// This member is required.
	FlowDefinitionName *string

	// An object containing information about the tasks the human reviewers will
	// perform.
	//
	// This member is required.
	HumanLoopConfig *types.HumanLoopConfig

	// An object containing information about where the human review results will be
	// uploaded.
	//
	// This member is required.
	OutputConfig *types.FlowDefinitionOutputConfig

	// The Amazon Resource Name (ARN) of the role needed to call other services on your
	// behalf. For example,
	// arn:aws:iam::1234567890:role/service-role/AmazonSageMaker-ExecutionRole-20180111T151298.
	//
	// This member is required.
	RoleArn *string

	// An object containing information about the events that trigger a human workflow.
	HumanLoopActivationConfig *types.HumanLoopActivationConfig

	// Container for configuring the source of human task requests. Use to specify if
	// Amazon Rekognition or Amazon Textract is used as an integration source.
	HumanLoopRequestSource *types.HumanLoopRequestSource

	// An array of key-value pairs that contain metadata to help you categorize and
	// organize a flow definition. Each tag consists of a key and a value, both of
	// which you define.
	Tags []types.Tag
}

type CreateFlowDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the flow definition you create.
	//
	// This member is required.
	FlowDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFlowDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFlowDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFlowDefinition{}, middleware.After)
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
	if err = addOpCreateFlowDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFlowDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateFlowDefinition",
	}
}
