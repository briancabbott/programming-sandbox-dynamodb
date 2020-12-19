// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information on VPC peering connections. Use this operation to get
// peering information for all fleets or for one specific fleet ID. To retrieve
// connection information, call this operation from the AWS account that is used to
// manage the Amazon GameLift fleets. Specify a fleet ID or leave the parameter
// empty to retrieve all connection records. If successful, the retrieved
// information includes both active and pending connections. Active connections
// identify the IpV4 CIDR block that the VPC uses to connect.
//
// *
// CreateVpcPeeringAuthorization
//
// * DescribeVpcPeeringAuthorizations
//
// *
// DeleteVpcPeeringAuthorization
//
// * CreateVpcPeeringConnection
//
// *
// DescribeVpcPeeringConnections
//
// * DeleteVpcPeeringConnection
func (c *Client) DescribeVpcPeeringConnections(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) {
	if params == nil {
		params = &DescribeVpcPeeringConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcPeeringConnections", params, optFns, addOperationDescribeVpcPeeringConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcPeeringConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeVpcPeeringConnectionsInput struct {

	// A unique identifier for a fleet. You can use either the fleet ID or ARN value.
	FleetId *string
}

// Represents the returned data in response to a request operation.
type DescribeVpcPeeringConnectionsOutput struct {

	// A collection of VPC peering connection records that match the request.
	VpcPeeringConnections []types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcPeeringConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringConnections",
	}
}