// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a type's registration, including its current status
// and type and version identifiers. When you initiate a registration request using
// RegisterType, you can then use DescribeTypeRegistration to monitor the progress
// of that registration request. Once the registration request has completed, use
// DescribeType to return detailed informaiton about a type.
func (c *Client) DescribeTypeRegistration(ctx context.Context, params *DescribeTypeRegistrationInput, optFns ...func(*Options)) (*DescribeTypeRegistrationOutput, error) {
	if params == nil {
		params = &DescribeTypeRegistrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTypeRegistration", params, optFns, addOperationDescribeTypeRegistrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTypeRegistrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTypeRegistrationInput struct {

	// The identifier for this registration request. This registration token is
	// generated by CloudFormation when you initiate a registration request using
	// RegisterType.
	//
	// This member is required.
	RegistrationToken *string
}

type DescribeTypeRegistrationOutput struct {

	// The description of the type registration request.
	Description *string

	// The current status of the type registration request.
	ProgressStatus types.RegistrationStatus

	// The Amazon Resource Name (ARN) of the type being registered. For registration
	// requests with a ProgressStatus of other than COMPLETE, this will be null.
	TypeArn *string

	// The Amazon Resource Name (ARN) of this specific version of the type being
	// registered. For registration requests with a ProgressStatus of other than
	// COMPLETE, this will be null.
	TypeVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTypeRegistrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTypeRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTypeRegistration{}, middleware.After)
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
	if err = addOpDescribeTypeRegistrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTypeRegistration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTypeRegistration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeTypeRegistration",
	}
}
