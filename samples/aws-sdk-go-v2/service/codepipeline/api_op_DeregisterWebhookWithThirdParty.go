// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the connection between the webhook that was created by CodePipeline and
// the external tool with events to be detected. Currently supported only for
// webhooks that target an action type of GitHub.
func (c *Client) DeregisterWebhookWithThirdParty(ctx context.Context, params *DeregisterWebhookWithThirdPartyInput, optFns ...func(*Options)) (*DeregisterWebhookWithThirdPartyOutput, error) {
	if params == nil {
		params = &DeregisterWebhookWithThirdPartyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterWebhookWithThirdParty", params, optFns, addOperationDeregisterWebhookWithThirdPartyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterWebhookWithThirdPartyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterWebhookWithThirdPartyInput struct {

	// The name of the webhook you want to deregister.
	WebhookName *string
}

type DeregisterWebhookWithThirdPartyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterWebhookWithThirdPartyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterWebhookWithThirdParty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterWebhookWithThirdParty{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterWebhookWithThirdParty(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterWebhookWithThirdParty(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "DeregisterWebhookWithThirdParty",
	}
}
