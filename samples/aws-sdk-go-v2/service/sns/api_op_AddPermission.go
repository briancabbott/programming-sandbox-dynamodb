// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a statement to a topic's access control policy, granting access for the
// specified AWS accounts to the specified actions.
func (c *Client) AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) {
	if params == nil {
		params = &AddPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPermission", params, optFns, addOperationAddPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddPermissionInput struct {

	// The AWS account IDs of the users (principals) who will be given access to the
	// specified actions. The users must have AWS accounts, but do not need to be
	// signed up for this service.
	//
	// This member is required.
	AWSAccountId []string

	// The action you want to allow for the specified principal(s). Valid values: Any
	// Amazon SNS action name, for example Publish.
	//
	// This member is required.
	ActionName []string

	// A unique identifier for the new policy statement.
	//
	// This member is required.
	Label *string

	// The ARN of the topic whose access control policy you wish to modify.
	//
	// This member is required.
	TopicArn *string
}

type AddPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddPermission{}, middleware.After)
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
	if err = addOpAddPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "AddPermission",
	}
}
