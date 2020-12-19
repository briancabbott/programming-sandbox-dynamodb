// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the association proposal request between the specified Direct Connect
// gateway and virtual private gateway or transit gateway.
func (c *Client) DeleteDirectConnectGatewayAssociationProposal(ctx context.Context, params *DeleteDirectConnectGatewayAssociationProposalInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayAssociationProposalOutput, error) {
	if params == nil {
		params = &DeleteDirectConnectGatewayAssociationProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDirectConnectGatewayAssociationProposal", params, optFns, addOperationDeleteDirectConnectGatewayAssociationProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDirectConnectGatewayAssociationProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayAssociationProposalInput struct {

	// The ID of the proposal.
	//
	// This member is required.
	ProposalId *string
}

type DeleteDirectConnectGatewayAssociationProposalOutput struct {

	// The ID of the associated gateway.
	DirectConnectGatewayAssociationProposal *types.DirectConnectGatewayAssociationProposal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDirectConnectGatewayAssociationProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGatewayAssociationProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGatewayAssociationProposal{}, middleware.After)
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
	if err = addOpDeleteDirectConnectGatewayAssociationProposalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociationProposal(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociationProposal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteDirectConnectGatewayAssociationProposal",
	}
}
