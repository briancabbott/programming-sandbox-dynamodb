// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an IAM resource that describes an identity provider (IdP) that supports
// SAML 2.0. The SAML provider resource that you create with this operation can be
// used as a principal in an IAM role's trust policy. Such a policy can enable
// federated users who sign in using the SAML IdP to assume the role. You can
// create an IAM role that supports Web-based single sign-on (SSO) to the AWS
// Management Console or one that supports API access to AWS. When you create the
// SAML provider resource, you upload a SAML metadata document that you get from
// your IdP. That document includes the issuer's name, expiration information, and
// keys that can be used to validate the SAML authentication response (assertions)
// that the IdP sends. You must generate the metadata document using the identity
// management software that is used as your organization's IdP. This operation
// requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). For
// more information, see Enabling SAML 2.0 Federated Users to Access the AWS
// Management Console
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-saml.html)
// and About SAML 2.0-based Federation
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html)
// in the IAM User Guide.
func (c *Client) CreateSAMLProvider(ctx context.Context, params *CreateSAMLProviderInput, optFns ...func(*Options)) (*CreateSAMLProviderOutput, error) {
	if params == nil {
		params = &CreateSAMLProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSAMLProvider", params, optFns, addOperationCreateSAMLProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSAMLProviderInput struct {

	// The name of the provider to create. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	Name *string

	// An XML document generated by an identity provider (IdP) that supports SAML 2.0.
	// The document includes the issuer's name, expiration information, and keys that
	// can be used to validate the SAML authentication response (assertions) that are
	// received from the IdP. You must generate the metadata document using the
	// identity management software that is used as your organization's IdP. For more
	// information, see About SAML 2.0-based Federation
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html)
	// in the IAM User Guide
	//
	// This member is required.
	SAMLMetadataDocument *string
}

// Contains the response to a successful CreateSAMLProvider request.
type CreateSAMLProviderOutput struct {

	// The Amazon Resource Name (ARN) of the new SAML provider resource in IAM.
	SAMLProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSAMLProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSAMLProvider{}, middleware.After)
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
	if err = addOpCreateSAMLProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSAMLProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSAMLProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateSAMLProvider",
	}
}
