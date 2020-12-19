// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Imports key material into an existing symmetric AWS KMS customer master key
// (CMK) that was created without key material. After you successfully import key
// material into a CMK, you can reimport the same key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material)
// into that CMK, but you cannot import different key material. You cannot perform
// this operation on an asymmetric CMK or on any CMK in a different AWS account.
// For more information about creating CMKs with no key material and then importing
// key material, see Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
// the AWS Key Management Service Developer Guide. Before using this operation,
// call GetParametersForImport. Its response includes a public key and an import
// token. Use the public key to encrypt the key material. Then, submit the import
// token from the same GetParametersForImport response. When calling this
// operation, you must specify the following values:
//
// * The key ID or key ARN of a
// CMK with no key material. Its Origin must be EXTERNAL. To create a CMK with no
// key material, call CreateKey and set the value of its Origin parameter to
// EXTERNAL. To get the Origin of a CMK, call DescribeKey.)
//
// * The encrypted key
// material. To get the public key to encrypt the key material, call
// GetParametersForImport.
//
// * The import token that GetParametersForImport
// returned. You must use a public key and token from the same
// GetParametersForImport response.
//
// * Whether the key material expires and if so,
// when. If you set an expiration date, AWS KMS deletes the key material from the
// CMK on the specified date, and the CMK becomes unusable. To use the CMK again,
// you must reimport the same key material. The only way to change an expiration
// date is by reimporting the same key material and specifying a new expiration
// date.
//
// When this operation is successful, the key state of the CMK changes from
// PendingImport to Enabled, and you can use the CMK. If this operation fails, use
// the exception to help determine the problem. If the error is related to the key
// material, the import token, or wrapping key, use GetParametersForImport to get a
// new public key and import token for the CMK and repeat the import procedure. For
// help, see How To Import Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#importing-keys-overview)
// in the AWS Key Management Service Developer Guide. The CMK that you use for this
// operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) ImportKeyMaterial(ctx context.Context, params *ImportKeyMaterialInput, optFns ...func(*Options)) (*ImportKeyMaterialOutput, error) {
	if params == nil {
		params = &ImportKeyMaterialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportKeyMaterial", params, optFns, addOperationImportKeyMaterialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportKeyMaterialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportKeyMaterialInput struct {

	// The encrypted key material to import. The key material must be encrypted with
	// the public wrapping key that GetParametersForImport returned, using the wrapping
	// algorithm that you specified in the same GetParametersForImport request.
	//
	// This member is required.
	EncryptedKeyMaterial []byte

	// The import token that you received in the response to a previous
	// GetParametersForImport request. It must be from the same response that contained
	// the public key that you used to encrypt the key material.
	//
	// This member is required.
	ImportToken []byte

	// The identifier of the symmetric CMK that receives the imported key material. The
	// CMK's Origin must be EXTERNAL. This must be the same CMK specified in the KeyID
	// parameter of the corresponding GetParametersForImport request. Specify the key
	// ID or the Amazon Resource Name (ARN) of the CMK. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// Specifies whether the key material expires. The default is KEY_MATERIAL_EXPIRES,
	// in which case you must include the ValidTo parameter. When this parameter is set
	// to KEY_MATERIAL_DOES_NOT_EXPIRE, you must omit the ValidTo parameter.
	ExpirationModel types.ExpirationModelType

	// The time at which the imported key material expires. When the key material
	// expires, AWS KMS deletes the key material and the CMK becomes unusable. You must
	// omit this parameter when the ExpirationModel parameter is set to
	// KEY_MATERIAL_DOES_NOT_EXPIRE. Otherwise it is required.
	ValidTo *time.Time
}

type ImportKeyMaterialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportKeyMaterialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportKeyMaterial{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportKeyMaterial{}, middleware.After)
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
	if err = addOpImportKeyMaterialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportKeyMaterial(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportKeyMaterial(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ImportKeyMaterial",
	}
}
