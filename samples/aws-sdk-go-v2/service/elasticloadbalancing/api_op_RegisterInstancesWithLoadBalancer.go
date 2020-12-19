// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified instances to the specified load balancer. The instance must
// be a running instance in the same network as the load balancer (EC2-Classic or
// the same VPC). If you have EC2-Classic instances and a load balancer in a VPC
// with ClassicLink enabled, you can link the EC2-Classic instances to that VPC and
// then register the linked EC2-Classic instances with the load balancer in the
// VPC. Note that RegisterInstanceWithLoadBalancer completes when the request has
// been registered. Instance registration takes a little time to complete. To check
// the state of the registered instances, use DescribeLoadBalancers or
// DescribeInstanceHealth. After the instance is registered, it starts receiving
// traffic and requests from the load balancer. Any instance that is not in one of
// the Availability Zones registered for the load balancer is moved to the
// OutOfService state. If an Availability Zone is added to the load balancer later,
// any instances registered with the load balancer move to the InService state. To
// deregister instances from a load balancer, use
// DeregisterInstancesFromLoadBalancer. For more information, see Register or
// De-Register EC2 Instances
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-deregister-register-instances.html)
// in the Classic Load Balancers Guide.
func (c *Client) RegisterInstancesWithLoadBalancer(ctx context.Context, params *RegisterInstancesWithLoadBalancerInput, optFns ...func(*Options)) (*RegisterInstancesWithLoadBalancerOutput, error) {
	if params == nil {
		params = &RegisterInstancesWithLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterInstancesWithLoadBalancer", params, optFns, addOperationRegisterInstancesWithLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterInstancesWithLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterInstancesWithLoadBalancer.
type RegisterInstancesWithLoadBalancerInput struct {

	// The IDs of the instances.
	//
	// This member is required.
	Instances []types.Instance

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the output of RegisterInstancesWithLoadBalancer.
type RegisterInstancesWithLoadBalancerOutput struct {

	// The updated list of instances for the load balancer.
	Instances []types.Instance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterInstancesWithLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRegisterInstancesWithLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterInstancesWithLoadBalancer{}, middleware.After)
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
	if err = addOpRegisterInstancesWithLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterInstancesWithLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterInstancesWithLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "RegisterInstancesWithLoadBalancer",
	}
}
