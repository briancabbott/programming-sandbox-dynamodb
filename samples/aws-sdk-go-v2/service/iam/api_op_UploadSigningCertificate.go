// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uploads an X.509 signing certificate and associates it with the specified IAM
// user. Some AWS services use X.509 signing certificates to validate requests that
// are signed with a corresponding private key. When you upload the certificate,
// its default status is Active. If the UserName is not specified, the IAM user
// name is determined implicitly based on the AWS access key ID used to sign the
// request. This operation works for access keys under the AWS account.
// Consequently, you can use this operation to manage AWS account root user
// credentials even if the AWS account has no associated users. Because the body of
// an X.509 certificate can be large, you should use POST rather than GET when
// calling UploadSigningCertificate. For information about setting up signatures
// and authorization through the API, go to Signing AWS API Requests
// (https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html) in
// the AWS General Reference. For general information about using the Query API
// with IAM, go to Making Query Requests
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html) in the
// IAM User Guide.
func (c *Client) UploadSigningCertificate(ctx context.Context, params *UploadSigningCertificateInput, optFns ...func(*Options)) (*UploadSigningCertificateOutput, error) {
	if params == nil {
		params = &UploadSigningCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadSigningCertificate", params, optFns, addOperationUploadSigningCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadSigningCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadSigningCertificateInput struct {

	// The contents of the signing certificate. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	// * Any printable ASCII character ranging
	// from the space character (\u0020) through the end of the ASCII character
	// range
	//
	// * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	// * The special characters tab (\u0009), line feed
	// (\u000A), and carriage return (\u000D)
	//
	// This member is required.
	CertificateBody *string

	// The name of the user the signing certificate is for. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
}

// Contains the response to a successful UploadSigningCertificate request.
type UploadSigningCertificateOutput struct {

	// Information about the certificate.
	//
	// This member is required.
	Certificate *types.SigningCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadSigningCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUploadSigningCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUploadSigningCertificate{}, middleware.After)
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
	if err = addOpUploadSigningCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadSigningCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUploadSigningCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UploadSigningCertificate",
	}
}
