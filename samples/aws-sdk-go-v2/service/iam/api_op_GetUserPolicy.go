// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the specified inline policy document that is embedded in the specified
// IAM user. Policies returned by this API are URL-encoded compliant with RFC 3986
// (https://tools.ietf.org/html/rfc3986). You can use a URL decoding method to
// convert the policy back to plain JSON text. For example, if you use Java, you
// can use the decode method of the java.net.URLDecoder utility class in the Java
// SDK. Other languages and SDKs provide similar functionality. An IAM user can
// also have managed policies attached to it. To retrieve a managed policy document
// that is attached to a user, use GetPolicy to determine the policy's default
// version. Then use GetPolicyVersion to retrieve the policy document. For more
// information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) GetUserPolicy(ctx context.Context, params *GetUserPolicyInput, optFns ...func(*Options)) (*GetUserPolicyOutput, error) {
	if params == nil {
		params = &GetUserPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserPolicy", params, optFns, addOperationGetUserPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserPolicyInput struct {

	// The name of the policy document to get. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyName *string

	// The name of the user who the policy is associated with. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

// Contains the response to a successful GetUserPolicy request.
type GetUserPolicyOutput struct {

	// The policy document. IAM stores policies in JSON format. However, resources that
	// were created using AWS CloudFormation templates can be formatted in YAML. AWS
	// CloudFormation always converts a YAML policy to JSON format before submitting it
	// to IAM.
	//
	// This member is required.
	PolicyDocument *string

	// The name of the policy.
	//
	// This member is required.
	PolicyName *string

	// The user the policy is associated with.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUserPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetUserPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetUserPolicy{}, middleware.After)
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
	if err = addOpGetUserPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetUserPolicy",
	}
}
