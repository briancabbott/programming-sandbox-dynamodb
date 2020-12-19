// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows network ingress to a cache security group. Applications using ElastiCache
// must be running on Amazon EC2, and Amazon EC2 security groups are used as the
// authorization mechanism. You cannot authorize ingress from an Amazon EC2
// security group in one region to an ElastiCache cluster in another region.
func (c *Client) AuthorizeCacheSecurityGroupIngress(ctx context.Context, params *AuthorizeCacheSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeCacheSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &AuthorizeCacheSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeCacheSecurityGroupIngress", params, optFns, addOperationAuthorizeCacheSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeCacheSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AuthorizeCacheSecurityGroupIngress operation.
type AuthorizeCacheSecurityGroupIngressInput struct {

	// The cache security group that allows network ingress.
	//
	// This member is required.
	CacheSecurityGroupName *string

	// The Amazon EC2 security group to be authorized for ingress to the cache security
	// group.
	//
	// This member is required.
	EC2SecurityGroupName *string

	// The AWS account number of the Amazon EC2 security group owner. Note that this is
	// not the same thing as an AWS access key ID - you must provide a valid AWS
	// account number for this parameter.
	//
	// This member is required.
	EC2SecurityGroupOwnerId *string
}

type AuthorizeCacheSecurityGroupIngressOutput struct {

	// Represents the output of one of the following operations:
	//
	// *
	// AuthorizeCacheSecurityGroupIngress
	//
	// * CreateCacheSecurityGroup
	//
	// *
	// RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *types.CacheSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAuthorizeCacheSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeCacheSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeCacheSecurityGroupIngress{}, middleware.After)
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
	if err = addOpAuthorizeCacheSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeCacheSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeCacheSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "AuthorizeCacheSecurityGroupIngress",
	}
}
