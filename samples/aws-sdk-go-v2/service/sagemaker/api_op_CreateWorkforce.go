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

// Use this operation to create a workforce. This operation will return an error if
// a workforce already exists in the AWS Region that you specify. You can only
// create one workforce in each AWS Region per AWS account. If you want to create a
// new workforce in an AWS Region where a workforce already exists, use the API
// operation to delete the existing workforce and then use CreateWorkforce to
// create a new workforce. To create a private workforce using Amazon Cognito, you
// must specify a Cognito user pool in CognitoConfig. You can also create an Amazon
// Cognito workforce using the Amazon SageMaker console. For more information, see
// Create a Private Workforce (Amazon Cognito)
// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html).
// To create a private workforce using your own OIDC Identity Provider (IdP),
// specify your IdP configuration in OidcConfig. Your OIDC IdP must support groups
// because groups are used by Ground Truth and Amazon A2I to create work teams. For
// more information, see  Create a Private Workforce (OIDC IdP)
// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private-oidc.html).
func (c *Client) CreateWorkforce(ctx context.Context, params *CreateWorkforceInput, optFns ...func(*Options)) (*CreateWorkforceOutput, error) {
	if params == nil {
		params = &CreateWorkforceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkforce", params, optFns, addOperationCreateWorkforceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkforceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkforceInput struct {

	// The name of the private workforce.
	//
	// This member is required.
	WorkforceName *string

	// Use this parameter to configure an Amazon Cognito private workforce. A single
	// Cognito workforce is created using and corresponds to a single  Amazon Cognito
	// user pool
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html).
	// Do not use OidcConfig if you specify values for CognitoConfig.
	CognitoConfig *types.CognitoConfig

	// Use this parameter to configure a private workforce using your own OIDC Identity
	// Provider. Do not use CognitoConfig if you specify values for OidcConfig.
	OidcConfig *types.OidcConfig

	// A list of IP address ranges (CIDRs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)). Used to
	// create an allow list of IP addresses for a private workforce. Workers will only
	// be able to login to their worker portal from an IP address within this range. By
	// default, a workforce isn't restricted to specific IP addresses.
	SourceIpConfig *types.SourceIpConfig

	// An array of key-value pairs that contain metadata to help you categorize and
	// organize our workforce. Each tag consists of a key and a value, both of which
	// you define.
	Tags []types.Tag
}

type CreateWorkforceOutput struct {

	// The Amazon Resource Name (ARN) of the workforce.
	//
	// This member is required.
	WorkforceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateWorkforceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkforce{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkforce{}, middleware.After)
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
	if err = addOpCreateWorkforceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkforce(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkforce(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateWorkforce",
	}
}
