// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the domain's endpoint options, specifically whether all requests to the
// domain must arrive over HTTPS. For more information, see Configuring Domain
// Endpoint Options
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-domain-endpoint-options.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) UpdateDomainEndpointOptions(ctx context.Context, params *UpdateDomainEndpointOptionsInput, optFns ...func(*Options)) (*UpdateDomainEndpointOptionsOutput, error) {
	if params == nil {
		params = &UpdateDomainEndpointOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainEndpointOptions", params, optFns, addOperationUpdateDomainEndpointOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainEndpointOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateDomainEndpointOptions operation.
// Specifies the name of the domain you want to update and the domain endpoint
// options.
type UpdateDomainEndpointOptionsInput struct {

	// Whether to require that all requests to the domain arrive over HTTPS. We
	// recommend Policy-Min-TLS-1-2-2019-07 for TLSSecurityPolicy. For compatibility
	// with older clients, the default is Policy-Min-TLS-1-0-2019-07.
	//
	// This member is required.
	DomainEndpointOptions *types.DomainEndpointOptions

	// A string that represents the name of a domain.
	//
	// This member is required.
	DomainName *string
}

// The result of a UpdateDomainEndpointOptions request. Contains the configuration
// and status of the domain's endpoint options.
type UpdateDomainEndpointOptionsOutput struct {

	// The newly-configured domain endpoint options.
	DomainEndpointOptions *types.DomainEndpointOptionsStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDomainEndpointOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateDomainEndpointOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateDomainEndpointOptions{}, middleware.After)
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
	if err = addOpUpdateDomainEndpointOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainEndpointOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainEndpointOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "UpdateDomainEndpointOptions",
	}
}
