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

// Disassociates a batch of self-service actions from the specified provisioning
// artifact.
func (c *Client) BatchDisassociateServiceActionFromProvisioningArtifact(ctx context.Context, params *BatchDisassociateServiceActionFromProvisioningArtifactInput, optFns ...func(*Options)) (*BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	if params == nil {
		params = &BatchDisassociateServiceActionFromProvisioningArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateServiceActionFromProvisioningArtifact", params, optFns, addOperationBatchDisassociateServiceActionFromProvisioningArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateServiceActionFromProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateServiceActionFromProvisioningArtifactInput struct {

	// One or more associations, each consisting of the Action ID, the Product ID, and
	// the Provisioning Artifact ID.
	//
	// This member is required.
	ServiceActionAssociations []types.ServiceActionAssociation

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string
}

type BatchDisassociateServiceActionFromProvisioningArtifactOutput struct {

	// An object that contains a list of errors, along with information to help you
	// identify the self-service action.
	FailedServiceActionAssociations []types.FailedServiceActionAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDisassociateServiceActionFromProvisioningArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDisassociateServiceActionFromProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDisassociateServiceActionFromProvisioningArtifact{}, middleware.After)
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
	if err = addOpBatchDisassociateServiceActionFromProvisioningArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateServiceActionFromProvisioningArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDisassociateServiceActionFromProvisioningArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "BatchDisassociateServiceActionFromProvisioningArtifact",
	}
}
