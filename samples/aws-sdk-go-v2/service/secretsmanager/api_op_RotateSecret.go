// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures and starts the asynchronous process of rotating this secret. If you
// include the configuration parameters, the operation sets those values for the
// secret and then immediately starts a rotation. If you do not include the
// configuration parameters, the operation starts a rotation with the values
// already stored in the secret. After the rotation completes, the protected
// service and its clients all use the new version of the secret. This required
// configuration information includes the ARN of an AWS Lambda function and the
// time between scheduled rotations. The Lambda rotation function creates a new
// version of the secret and creates or updates the credentials on the protected
// service to match. After testing the new credentials, the function marks the new
// secret with the staging label AWSCURRENT so that your clients all immediately
// begin to use the new version. For more information about rotating secrets and
// how to configure a Lambda function to rotate the secrets for your protected
// service, see Rotating Secrets in AWS Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html)
// in the AWS Secrets Manager User Guide. Secrets Manager schedules the next
// rotation when the previous one completes. Secrets Manager schedules the date by
// adding the rotation interval (number of days) to the actual date of the last
// rotation. The service chooses the hour within that 24-hour date window randomly.
// The minute is also chosen somewhat randomly, but weighted towards the top of the
// hour and influenced by a variety of factors that help distribute load. The
// rotation function must end with the versions of the secret in one of two
// states:
//
// * The AWSPENDING and AWSCURRENT staging labels are attached to the same
// version of the secret, or
//
// * The AWSPENDING staging label is not attached to any
// version of the secret.
//
// If the AWSPENDING staging label is present but not
// attached to the same version as AWSCURRENT then any later invocation of
// RotateSecret assumes that a previous rotation request is still in progress and
// returns an error. Minimum permissions To run this command, you must have the
// following permissions:
//
// * secretsmanager:RotateSecret
//
// * lambda:InvokeFunction
// (on the function specified in the secret's metadata)
//
// Related operations
//
// * To
// list the secrets in your account, use ListSecrets.
//
// * To get the details for a
// version of a secret, use DescribeSecret.
//
// * To create a new version of a secret,
// use CreateSecret.
//
// * To attach staging labels to or remove staging labels from a
// version of a secret, use UpdateSecretVersionStage.
func (c *Client) RotateSecret(ctx context.Context, params *RotateSecretInput, optFns ...func(*Options)) (*RotateSecretOutput, error) {
	if params == nil {
		params = &RotateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RotateSecret", params, optFns, addOperationRotateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RotateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateSecretInput struct {

	// Specifies the secret that you want to rotate. You can specify either the Amazon
	// Resource Name (ARN) or the friendly name of the secret. If you specify an ARN,
	// we generally recommend that you specify a complete ARN. You can specify a
	// partial ARN too—for example, if you don’t include the final hyphen and six
	// random characters that Secrets Manager adds at the end of the ARN when you
	// created the secret. A partial ARN match can work as long as it uniquely matches
	// only one secret. However, if your secret has a name that ends in a hyphen
	// followed by six characters (before Secrets Manager adds the hyphen and six
	// characters to the ARN) and you try to use that as a partial ARN, then those
	// characters cause Secrets Manager to assume that you’re specifying a complete
	// ARN. This confusion can cause unexpected results. To avoid this situation, we
	// recommend that you don’t create secret names ending with a hyphen followed by
	// six characters. If you specify an incomplete ARN without the random suffix, and
	// instead provide the 'friendly name', you must not include the random suffix. If
	// you do include the random suffix added by Secrets Manager, you receive either a
	// ResourceNotFoundException or an AccessDeniedException error, depending on your
	// permissions.
	//
	// This member is required.
	SecretId *string

	// (Optional) Specifies a unique identifier for the new version of the secret that
	// helps ensure idempotency. If you use the AWS CLI or one of the AWS SDK to call
	// this operation, then you can leave this parameter empty. The CLI or SDK
	// generates a random UUID for you and includes that in the request for this
	// parameter. If you don't use the SDK and instead generate a raw HTTP request to
	// the Secrets Manager service endpoint, then you must generate a
	// ClientRequestToken yourself for new versions and include that value in the
	// request. You only need to specify your own value if you implement your own retry
	// logic and want to ensure that a given secret is not created twice. We recommend
	// that you generate a UUID-type
	// (https://wikipedia.org/wiki/Universally_unique_identifier) value to ensure
	// uniqueness within the specified secret. Secrets Manager uses this value to
	// prevent the accidental creation of duplicate versions if there are failures and
	// retries during the function's processing. This value becomes the VersionId of
	// the new version.
	ClientRequestToken *string

	// (Optional) Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaARN *string

	// A structure that defines the rotation configuration for this secret.
	RotationRules *types.RotationRulesType
}

type RotateSecretOutput struct {

	// The ARN of the secret.
	ARN *string

	// The friendly name of the secret.
	Name *string

	// The ID of the new version of the secret created by the rotation started by this
	// request.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRotateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRotateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRotateSecret{}, middleware.After)
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
	if err = addIdempotencyToken_opRotateSecretMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRotateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRotateSecret(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRotateSecret struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRotateSecret) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRotateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RotateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RotateSecretInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRotateSecretMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRotateSecret{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRotateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "RotateSecret",
	}
}
