// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops sharing the specified portfolio with the specified account or organization
// node. Shares to an organization node can only be deleted by the management
// account of an organization or by a delegated administrator. Note that if a
// delegated admin is de-registered, portfolio shares created from that account are
// removed.
func (c *Client) DeletePortfolioShare(ctx context.Context, params *DeletePortfolioShareInput, optFns ...func(*Options)) (*DeletePortfolioShareOutput, error) {
	if params == nil {
		params = &DeletePortfolioShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePortfolioShare", params, optFns, addOperationDeletePortfolioShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePortfolioShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePortfolioShareInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The AWS account ID.
	AccountId *string

	// The organization node to whom you are going to stop sharing.
	OrganizationNode *types.OrganizationNode
}

type DeletePortfolioShareOutput struct {

	// The portfolio share unique identifier. This will only be returned if delete is
	// made to an organization node.
	PortfolioShareToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePortfolioShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePortfolioShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePortfolioShare{}, middleware.After)
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
	if err = addOpDeletePortfolioShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePortfolioShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePortfolioShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DeletePortfolioShare",
	}
}