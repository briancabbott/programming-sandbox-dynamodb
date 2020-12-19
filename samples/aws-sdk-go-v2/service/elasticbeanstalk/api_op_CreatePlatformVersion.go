// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new version of your custom platform.
func (c *Client) CreatePlatformVersion(ctx context.Context, params *CreatePlatformVersionInput, optFns ...func(*Options)) (*CreatePlatformVersionOutput, error) {
	if params == nil {
		params = &CreatePlatformVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlatformVersion", params, optFns, addOperationCreatePlatformVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlatformVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create a new platform version.
type CreatePlatformVersionInput struct {

	// The location of the platform definition archive in Amazon S3.
	//
	// This member is required.
	PlatformDefinitionBundle *types.S3Location

	// The name of your custom platform.
	//
	// This member is required.
	PlatformName *string

	// The number, such as 1.0.2, for the new platform version.
	//
	// This member is required.
	PlatformVersion *string

	// The name of the builder environment.
	EnvironmentName *string

	// The configuration option settings to apply to the builder environment.
	OptionSettings []types.ConfigurationOptionSetting

	// Specifies the tags applied to the new platform version. Elastic Beanstalk
	// applies these tags only to the platform version. Environments that you create
	// using the platform version don't inherit the tags.
	Tags []types.Tag
}

type CreatePlatformVersionOutput struct {

	// The builder used to create the custom platform.
	Builder *types.Builder

	// Detailed information about the new version of the custom platform.
	PlatformSummary *types.PlatformSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePlatformVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePlatformVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePlatformVersion{}, middleware.After)
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
	if err = addOpCreatePlatformVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlatformVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlatformVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreatePlatformVersion",
	}
}
