// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// For asynchronous function invocation, use Invoke. Invokes a function
// asynchronously.
func (c *Client) InvokeAsync(ctx context.Context, params *InvokeAsyncInput, optFns ...func(*Options)) (*InvokeAsyncOutput, error) {
	if params == nil {
		params = &InvokeAsyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeAsync", params, optFns, addOperationInvokeAsyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeAsyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeAsyncInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - my-function.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The JSON that you want to provide to your Lambda function as input.
	//
	// This member is required.
	InvokeArgs io.Reader
}

// A success response (202 Accepted) indicates that the request is queued for
// invocation.
type InvokeAsyncOutput struct {

	// The status code.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInvokeAsyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeAsync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeAsync{}, middleware.After)
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
	if err = addOpInvokeAsyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeAsync(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeAsync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "InvokeAsync",
	}
}
