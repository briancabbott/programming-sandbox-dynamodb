// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a permission set within a specified SSO instance. To grant users and
// groups access to AWS account resources, use CreateAccountAssignment.
func (c *Client) CreatePermissionSet(ctx context.Context, params *CreatePermissionSetInput, optFns ...func(*Options)) (*CreatePermissionSetOutput, error) {
	if params == nil {
		params = &CreatePermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePermissionSet", params, optFns, addOperationCreatePermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The name of the PermissionSet.
	//
	// This member is required.
	Name *string

	// The description of the PermissionSet.
	Description *string

	// Used to redirect users within the application during the federation
	// authentication process.
	RelayState *string

	// The length of time that the application user sessions are valid in the ISO-8601
	// standard.
	SessionDuration *string

	// The tags to attach to the new PermissionSet.
	Tags []types.Tag
}

type CreatePermissionSetOutput struct {

	// Defines the level of access on an AWS account.
	PermissionSet *types.PermissionSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePermissionSet{}, middleware.After)
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
	if err = addOpCreatePermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "CreatePermissionSet",
	}
}
