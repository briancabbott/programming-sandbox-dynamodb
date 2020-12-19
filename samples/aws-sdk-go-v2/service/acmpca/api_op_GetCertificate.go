// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a certificate from your private CA or one that has been shared with
// you. The ARN of the certificate is returned when you call the IssueCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_IssueCertificate.html)
// action. You must specify both the ARN of your private CA and the ARN of the
// issued certificate when calling the GetCertificate action. You can retrieve the
// certificate if it is in the ISSUED state. You can call the
// CreateCertificateAuthorityAuditReport
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html)
// action to create a report that contains information about all of the
// certificates issued and revoked by your private CA.
func (c *Client) GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) {
	if params == nil {
		params = &GetCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificate", params, optFns, addOperationGetCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateInput struct {

	// The ARN of the issued certificate. The ARN contains the certificate serial
	// number and must be in the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/286535153982981100925020015808220737245
	//
	// This member is required.
	CertificateArn *string

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html).
	// This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type GetCertificateOutput struct {

	// The base64 PEM-encoded certificate specified by the CertificateArn parameter.
	Certificate *string

	// The base64 PEM-encoded certificate chain that chains up to the on-premises root
	// CA certificate that you used to sign your private CA certificate.
	CertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificate{}, middleware.After)
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
	if err = addOpGetCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetCertificate",
	}
}
