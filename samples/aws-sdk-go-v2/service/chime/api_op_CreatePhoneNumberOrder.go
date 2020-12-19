// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an order for phone numbers to be provisioned. Choose from Amazon Chime
// Business Calling and Amazon Chime Voice Connector product types. For toll-free
// numbers, you must use the Amazon Chime Voice Connector product type.
func (c *Client) CreatePhoneNumberOrder(ctx context.Context, params *CreatePhoneNumberOrderInput, optFns ...func(*Options)) (*CreatePhoneNumberOrderOutput, error) {
	if params == nil {
		params = &CreatePhoneNumberOrderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePhoneNumberOrder", params, optFns, addOperationCreatePhoneNumberOrderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePhoneNumberOrderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePhoneNumberOrderInput struct {

	// List of phone numbers, in E.164 format.
	//
	// This member is required.
	E164PhoneNumbers []string

	// The phone number product type.
	//
	// This member is required.
	ProductType types.PhoneNumberProductType
}

type CreatePhoneNumberOrderOutput struct {

	// The phone number order details.
	PhoneNumberOrder *types.PhoneNumberOrder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePhoneNumberOrderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePhoneNumberOrder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePhoneNumberOrder{}, middleware.After)
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
	if err = addOpCreatePhoneNumberOrderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePhoneNumberOrder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePhoneNumberOrder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreatePhoneNumberOrder",
	}
}