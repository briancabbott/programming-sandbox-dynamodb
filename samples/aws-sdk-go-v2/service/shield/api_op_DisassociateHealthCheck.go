// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes health-based detection from the Shield Advanced protection for a
// resource. Shield Advanced health-based detection uses the health of your AWS
// resource to improve responsiveness and accuracy in attack detection and
// mitigation. You define the health check in Route 53 and then associate or
// disassociate it with your Shield Advanced protection. For more information, see
// Shield Advanced Health-Based Detection
// (https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option)
// in the AWS WAF and AWS Shield Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) DisassociateHealthCheck(ctx context.Context, params *DisassociateHealthCheckInput, optFns ...func(*Options)) (*DisassociateHealthCheckOutput, error) {
	if params == nil {
		params = &DisassociateHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateHealthCheck", params, optFns, addOperationDisassociateHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateHealthCheckInput struct {

	// The Amazon Resource Name (ARN) of the health check that is associated with the
	// protection.
	//
	// This member is required.
	HealthCheckArn *string

	// The unique identifier (ID) for the Protection object to remove the health check
	// association from.
	//
	// This member is required.
	ProtectionId *string
}

type DisassociateHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateHealthCheck{}, middleware.After)
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
	if err = addOpDisassociateHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateHealthCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DisassociateHealthCheck",
	}
}
