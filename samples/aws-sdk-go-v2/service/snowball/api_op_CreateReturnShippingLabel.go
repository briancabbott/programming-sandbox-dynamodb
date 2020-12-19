// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a shipping label that will be used to return the Snow device to AWS.
func (c *Client) CreateReturnShippingLabel(ctx context.Context, params *CreateReturnShippingLabelInput, optFns ...func(*Options)) (*CreateReturnShippingLabelOutput, error) {
	if params == nil {
		params = &CreateReturnShippingLabelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReturnShippingLabel", params, optFns, addOperationCreateReturnShippingLabelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReturnShippingLabelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReturnShippingLabelInput struct {

	// The ID for a job that you want to create the return shipping label for. For
	// example JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string

	// The shipping speed for a particular job. This speed doesn't dictate how soon the
	// device is returned to AWS. This speed represents how quickly it moves to its
	// destination while in transit. Regional shipping speeds are as follows:
	ShippingOption types.ShippingOption
}

type CreateReturnShippingLabelOutput struct {

	// The status information of the task on a Snow device that is being returned to
	// AWS.
	Status types.ShippingLabelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateReturnShippingLabelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReturnShippingLabel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReturnShippingLabel{}, middleware.After)
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
	if err = addOpCreateReturnShippingLabelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReturnShippingLabel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReturnShippingLabel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateReturnShippingLabel",
	}
}
