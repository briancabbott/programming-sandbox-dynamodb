// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the resource-based policy attached to a private CA. Deletion will remove
// any access that the policy has granted. If there is no policy attached to the
// private CA, this action will return successful. If you delete a policy that was
// applied through AWS Resource Access Manager (RAM), the CA will be removed from
// all shares in which it was included. The AWS Certificate Manager Service Linked
// Role that the policy supports is not affected when you delete the policy. The
// current policy can be shown with GetPolicy
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_GetPolicy.html) and
// updated with PutPolicy
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_PutPolicy.html).
// About Policies
//
// * A policy grants access on a private CA to an AWS customer
// account, to AWS Organizations, or to an AWS Organizations unit. Policies are
// under the control of a CA administrator. For more information, see Using a
// Resource Based Policy with ACM Private CA.
//
// * A policy permits a user of AWS
// Certificate Manager (ACM) to issue ACM certificates signed by a CA in another
// account.
//
// * For ACM to manage automatic renewal of these certificates, the ACM
// user must configure a Service Linked Role (SLR). The SLR allows the ACM service
// to assume the identity of the user, subject to confirmation against the ACM
// Private CA policy. For more information, see Using a Service Linked Role with
// ACM (https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html).
//
// * Updates
// made in AWS Resource Manager (RAM) are reflected in policies. For more
// information, see Using AWS Resource Access Manager (RAM) with ACM Private CA.
func (c *Client) DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) {
	if params == nil {
		params = &DeletePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePolicy", params, optFns, addOperationDeletePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePolicyInput struct {

	// The Amazon Resource Number (ARN) of the private CA that will have its policy
	// deleted. You can find the CA's ARN by calling the ListCertificateAuthorities
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListCertificateAuthorities.html)
	// action. The ARN value must have the form
	// arn:aws:acm-pca:region:account:certificate-authority/01234567-89ab-cdef-0123-0123456789ab.
	//
	// This member is required.
	ResourceArn *string
}

type DeletePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePolicy{}, middleware.After)
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
	if err = addOpDeletePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DeletePolicy",
	}
}
