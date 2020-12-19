// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the Service Quotas template. Once the template is disabled, it does not
// request quota increases for new accounts in your organization. Disabling the
// quota template does not apply the quota increase requests from the template.
// Related operations
//
// * To enable the quota template, call
// AssociateServiceQuotaTemplate.
//
// * To delete a specific service quota from the
// template, use DeleteServiceQuotaIncreaseRequestFromTemplate.
func (c *Client) DisassociateServiceQuotaTemplate(ctx context.Context, params *DisassociateServiceQuotaTemplateInput, optFns ...func(*Options)) (*DisassociateServiceQuotaTemplateOutput, error) {
	if params == nil {
		params = &DisassociateServiceQuotaTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateServiceQuotaTemplate", params, optFns, addOperationDisassociateServiceQuotaTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateServiceQuotaTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateServiceQuotaTemplateInput struct {
}

type DisassociateServiceQuotaTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateServiceQuotaTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateServiceQuotaTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateServiceQuotaTemplate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateServiceQuotaTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateServiceQuotaTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "DisassociateServiceQuotaTemplate",
	}
}
