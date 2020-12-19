// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get information about an existing configuration set, including the dedicated IP
// pool that it's associated with, whether or not it's enabled for sending email,
// and more. Configuration sets are groups of rules that you can apply to the
// emails you send. You apply a configuration set to an email by including a
// reference to the configuration set in the headers of the email. When you apply a
// configuration set to an email, all of the rules in that configuration set are
// applied to the email.
func (c *Client) GetConfigurationSet(ctx context.Context, params *GetConfigurationSetInput, optFns ...func(*Options)) (*GetConfigurationSetOutput, error) {
	if params == nil {
		params = &GetConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfigurationSet", params, optFns, addOperationGetConfigurationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain information about a configuration set.
type GetConfigurationSetInput struct {

	// The name of the configuration set that you want to obtain more information
	// about.
	//
	// This member is required.
	ConfigurationSetName *string
}

// Information about a configuration set.
type GetConfigurationSetOutput struct {

	// The name of the configuration set.
	ConfigurationSetName *string

	// An object that defines the dedicated IP pool that is used to send emails that
	// you send using the configuration set.
	DeliveryOptions *types.DeliveryOptions

	// An object that defines whether or not Amazon SES collects reputation metrics for
	// the emails that you send that use the configuration set.
	ReputationOptions *types.ReputationOptions

	// An object that defines whether or not Amazon SES can send email that you send
	// using the configuration set.
	SendingOptions *types.SendingOptions

	// An object that contains information about the suppression list preferences for
	// your account.
	SuppressionOptions *types.SuppressionOptions

	// An array of objects that define the tags (keys and values) that are associated
	// with the configuration set.
	Tags []types.Tag

	// An object that defines the open and click tracking options for emails that you
	// send using the configuration set.
	TrackingOptions *types.TrackingOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConfigurationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationSet{}, middleware.After)
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
	if err = addOpGetConfigurationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfigurationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetConfigurationSet",
	}
}
