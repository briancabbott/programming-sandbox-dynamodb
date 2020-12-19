// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the encryption configuration for X-Ray data.
func (c *Client) PutEncryptionConfig(ctx context.Context, params *PutEncryptionConfigInput, optFns ...func(*Options)) (*PutEncryptionConfigOutput, error) {
	if params == nil {
		params = &PutEncryptionConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEncryptionConfig", params, optFns, addOperationPutEncryptionConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEncryptionConfigInput struct {

	// The type of encryption. Set to KMS to use your own key for encryption. Set to
	// NONE for default encryption.
	//
	// This member is required.
	Type types.EncryptionType

	// An AWS KMS customer master key (CMK) in one of the following formats:
	//
	// * Alias -
	// The name of the key. For example, alias/MyKey.
	//
	// * Key ID - The KMS key ID of the
	// key. For example, ae4aa6d49-a4d8-9df9-a475-4ff6d7898456. AWS X-Ray does not
	// support asymmetric CMKs.
	//
	// * ARN - The full Amazon Resource Name of the key ID or
	// alias. For example,
	// arn:aws:kms:us-east-2:123456789012:key/ae4aa6d49-a4d8-9df9-a475-4ff6d7898456.
	// Use this format to specify a key in a different account.
	//
	// Omit this key if you
	// set Type to NONE.
	KeyId *string
}

type PutEncryptionConfigOutput struct {

	// The new encryption configuration.
	EncryptionConfig *types.EncryptionConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutEncryptionConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEncryptionConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEncryptionConfig{}, middleware.After)
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
	if err = addOpPutEncryptionConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEncryptionConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEncryptionConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "PutEncryptionConfig",
	}
}
