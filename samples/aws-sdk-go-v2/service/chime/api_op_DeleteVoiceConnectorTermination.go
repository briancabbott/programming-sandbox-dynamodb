// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the termination settings for the specified Amazon Chime Voice Connector.
// If emergency calling is configured for the Amazon Chime Voice Connector, it must
// be deleted prior to deleting the termination settings.
func (c *Client) DeleteVoiceConnectorTermination(ctx context.Context, params *DeleteVoiceConnectorTerminationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationOutput, error) {
	if params == nil {
		params = &DeleteVoiceConnectorTerminationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVoiceConnectorTermination", params, optFns, addOperationDeleteVoiceConnectorTerminationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVoiceConnectorTerminationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorTerminationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type DeleteVoiceConnectorTerminationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVoiceConnectorTerminationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnectorTermination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnectorTermination{}, middleware.After)
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
	if err = addOpDeleteVoiceConnectorTerminationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnectorTermination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVoiceConnectorTermination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteVoiceConnectorTermination",
	}
}
