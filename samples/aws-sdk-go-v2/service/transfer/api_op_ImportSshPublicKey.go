// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a Secure Shell (SSH) public key to a user account identified by a UserName
// value assigned to the specific file transfer protocol-enabled server, identified
// by ServerId. The response returns the UserName value, the ServerId value, and
// the name of the SshPublicKeyId.
func (c *Client) ImportSshPublicKey(ctx context.Context, params *ImportSshPublicKeyInput, optFns ...func(*Options)) (*ImportSshPublicKeyOutput, error) {
	if params == nil {
		params = &ImportSshPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportSshPublicKey", params, optFns, addOperationImportSshPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportSshPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSshPublicKeyInput struct {

	// A system-assigned unique identifier for a server.
	//
	// This member is required.
	ServerId *string

	// The public key portion of an SSH key pair.
	//
	// This member is required.
	SshPublicKeyBody *string

	// The name of the user account that is assigned to one or more servers.
	//
	// This member is required.
	UserName *string
}

// Identifies the user, the server they belong to, and the identifier of the SSH
// public key associated with that user. A user can have more than one key on each
// server that they are associated with.
type ImportSshPublicKeyOutput struct {

	// A system-assigned unique identifier for a server.
	//
	// This member is required.
	ServerId *string

	// The name given to a public key by the system that was imported.
	//
	// This member is required.
	SshPublicKeyId *string

	// A user name assigned to the ServerID value that you specified.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportSshPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportSshPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportSshPublicKey{}, middleware.After)
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
	if err = addOpImportSshPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportSshPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportSshPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ImportSshPublicKey",
	}
}