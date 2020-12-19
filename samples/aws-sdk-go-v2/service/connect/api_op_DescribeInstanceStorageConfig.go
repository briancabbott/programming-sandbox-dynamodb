// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current storage configurations for the specified resource type,
// association ID, and instance ID.
func (c *Client) DescribeInstanceStorageConfig(ctx context.Context, params *DescribeInstanceStorageConfigInput, optFns ...func(*Options)) (*DescribeInstanceStorageConfigOutput, error) {
	if params == nil {
		params = &DescribeInstanceStorageConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceStorageConfig", params, optFns, addOperationDescribeInstanceStorageConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceStorageConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceStorageConfigInput struct {

	// The existing association identifier that uniquely identifies the resource type
	// and storage config for the given instance ID.
	//
	// This member is required.
	AssociationId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// A valid resource type.
	//
	// This member is required.
	ResourceType types.InstanceStorageResourceType
}

type DescribeInstanceStorageConfigOutput struct {

	// A valid storage type.
	StorageConfig *types.InstanceStorageConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstanceStorageConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInstanceStorageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInstanceStorageConfig{}, middleware.After)
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
	if err = addOpDescribeInstanceStorageConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceStorageConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceStorageConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "DescribeInstanceStorageConfig",
	}
}
