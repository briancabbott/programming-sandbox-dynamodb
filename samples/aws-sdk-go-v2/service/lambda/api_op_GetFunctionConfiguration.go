// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the version-specific settings of a Lambda function or version. The
// output includes only options that can vary between versions of a function. To
// modify these settings, use UpdateFunctionConfiguration. To get all of a
// function's details, including function-level settings, use GetFunction.
func (c *Client) GetFunctionConfiguration(ctx context.Context, params *GetFunctionConfigurationInput, optFns ...func(*Options)) (*GetFunctionConfigurationOutput, error) {
	if params == nil {
		params = &GetFunctionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFunctionConfiguration", params, optFns, addOperationGetFunctionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFunctionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionConfigurationInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//
	// * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	// * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// * Partial ARN -
	// 123456789012:function:my-function.
	//
	// You can append a version number or alias to
	// any of the formats. The length constraint applies only to the full ARN. If you
	// specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify a version or alias to get details about a published version of the
	// function.
	Qualifier *string
}

// Details about a function's configuration.
type GetFunctionConfigurationOutput struct {

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables.
	Environment *types.EnvironmentResponse

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []types.Layer

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string

	// The memory that's allocated to the function.
	MemorySize *int32

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime types.Runtime

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFunctionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionConfiguration{}, middleware.After)
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
	if err = addOpGetFunctionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFunctionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionConfiguration",
	}
}