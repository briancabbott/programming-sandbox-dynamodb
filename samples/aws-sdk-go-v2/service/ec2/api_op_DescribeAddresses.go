// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Elastic IP addresses or all of your Elastic IP
// addresses. An Elastic IP address is for use in either the EC2-Classic platform
// or in a VPC. For more information, see Elastic IP Addresses
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeAddresses(ctx context.Context, params *DescribeAddressesInput, optFns ...func(*Options)) (*DescribeAddressesOutput, error) {
	if params == nil {
		params = &DescribeAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddresses", params, optFns, addOperationDescribeAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressesInput struct {

	// [EC2-VPC] Information about the allocation IDs.
	AllocationIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	// *
	// allocation-id - [EC2-VPC] The allocation ID for the address.
	//
	// * association-id -
	// [EC2-VPC] The association ID for the address.
	//
	// * domain - Indicates whether the
	// address is for use in EC2-Classic (standard) or in a VPC (vpc).
	//
	// * instance-id -
	// The ID of the instance the address is associated with, if any.
	//
	// *
	// network-border-group - A unique set of Availability Zones, Local Zones, or
	// Wavelength Zones from where AWS advertises IP addresses.
	//
	// * network-interface-id
	// - [EC2-VPC] The ID of the network interface that the address is associated with,
	// if any.
	//
	// * network-interface-owner-id - The AWS account ID of the owner.
	//
	// *
	// private-ip-address - [EC2-VPC] The private IP address associated with the
	// Elastic IP address.
	//
	// * public-ip - The Elastic IP address, or the carrier IP
	// address.
	//
	// * tag: - The key/value combination of a tag assigned to the resource.
	// Use the tag key in the filter name and the tag value as the filter value. For
	// example, to find all resources that have a tag with the key Owner and the value
	// TeamA, specify tag:Owner for the filter name and TeamA for the filter value.
	//
	// *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	Filters []types.Filter

	// One or more Elastic IP addresses. Default: Describes all your Elastic IP
	// addresses.
	PublicIps []string
}

type DescribeAddressesOutput struct {

	// Information about the Elastic IP addresses.
	Addresses []types.Address

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAddresses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddresses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAddresses",
	}
}
