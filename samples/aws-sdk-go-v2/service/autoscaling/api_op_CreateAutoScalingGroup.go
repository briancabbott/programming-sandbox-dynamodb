// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Auto Scaling group with the specified name and attributes. If you
// exceed your maximum limit of Auto Scaling groups, the call fails. To query this
// limit, call the DescribeAccountLimits API. For information about updating this
// limit, see Amazon EC2 Auto Scaling service quotas
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide. For introductory exercises for
// creating an Auto Scaling group, see Getting started with Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/GettingStartedTutorial.html)
// and Tutorial: Set up a scaled and load-balanced application
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-register-lbs-with-asg.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Auto
// Scaling groups
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html) in
// the Amazon EC2 Auto Scaling User Guide. Every Auto Scaling group has three size
// parameters (DesiredCapacity, MaxSize, and MinSize). Usually, you set these sizes
// based on a specific number of instances. However, if you configure a mixed
// instances policy that defines weights for the instance types, you must specify
// these sizes with the same units that you use for weighting instances.
func (c *Client) CreateAutoScalingGroup(ctx context.Context, params *CreateAutoScalingGroupInput, optFns ...func(*Options)) (*CreateAutoScalingGroupOutput, error) {
	if params == nil {
		params = &CreateAutoScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoScalingGroup", params, optFns, addOperationCreateAutoScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoScalingGroupInput struct {

	// The name of the Auto Scaling group. This name must be unique per Region per
	// account.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum size of the group. With a mixed instances policy that uses instance
	// weighting, Amazon EC2 Auto Scaling may need to go above MaxSize to meet your
	// capacity requirements. In this event, Amazon EC2 Auto Scaling will never go
	// above MaxSize by more than your largest instance weight (weights that define how
	// many units each instance contributes to the desired capacity of the group).
	//
	// This member is required.
	MaxSize *int32

	// The minimum size of the group.
	//
	// This member is required.
	MinSize *int32

	// A list of Availability Zones where instances in the Auto Scaling group can be
	// created. This parameter is optional if you specify one or more subnets for
	// VPCZoneIdentifier. Conditional: If your account supports EC2-Classic and VPC,
	// this parameter is required to launch instances into EC2-Classic.
	AvailabilityZones []string

	// Indicates whether Capacity Rebalancing is enabled. Otherwise, Capacity
	// Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto
	// Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a
	// Spot Instance is at an elevated risk of interruption. After launching a new
	// instance, it then terminates an old instance. For more information, see Amazon
	// EC2 Auto Scaling Capacity Rebalancing
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/capacity-rebalance.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	CapacityRebalance *bool

	// The amount of time, in seconds, after a scaling activity completes before
	// another scaling activity can start. The default value is 300. This setting
	// applies when using simple scaling policies, but not when using other scaling
	// policies or scheduled scaling. For more information, see Scaling cooldowns for
	// Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	DefaultCooldown *int32

	// The desired capacity is the initial capacity of the Auto Scaling group at the
	// time of its creation and the capacity it attempts to maintain. It can scale
	// beyond this capacity if you configure auto scaling. This number must be greater
	// than or equal to the minimum size of the group and less than or equal to the
	// maximum size of the group. If you do not specify a desired capacity, the default
	// is the minimum size of the group.
	DesiredCapacity *int32

	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before
	// checking the health status of an EC2 instance that has come into service. During
	// this time, any health check failures for the instance are ignored. The default
	// value is 0. For more information, see Health check grace period
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period)
	// in the Amazon EC2 Auto Scaling User Guide. Conditional: Required if you are
	// adding an ELB health check.
	HealthCheckGracePeriod *int32

	// The service to use for the health checks. The valid values are EC2 (default) and
	// ELB. If you configure an Auto Scaling group to use load balancer (ELB) health
	// checks, it considers the instance unhealthy if it fails either the EC2 status
	// checks or the load balancer health checks. For more information, see Health
	// checks for Auto Scaling instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	HealthCheckType *string

	// The ID of the instance used to base the launch configuration on. If specified,
	// Amazon EC2 Auto Scaling uses the configuration values from the specified
	// instance to create a new launch configuration. To get the instance ID, use the
	// Amazon EC2 DescribeInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html)
	// API operation. For more information, see Creating an Auto Scaling group using an
	// EC2 instance
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceId *string

	// The name of the launch configuration to use to launch instances. Conditional:
	// You must specify either a launch template (LaunchTemplate or
	// MixedInstancesPolicy) or a launch configuration (LaunchConfigurationName or
	// InstanceId).
	LaunchConfigurationName *string

	// Parameters used to specify the launch template
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html)
	// and version to use to launch instances. Conditional: You must specify either a
	// launch template (LaunchTemplate or MixedInstancesPolicy) or a launch
	// configuration (LaunchConfigurationName or InstanceId). The launch template that
	// is specified must be configured for use with an Auto Scaling group. For more
	// information, see Creating a launch template for an Auto Scaling group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplate *types.LaunchTemplateSpecification

	// One or more lifecycle hooks for the group, which specify actions to perform when
	// Amazon EC2 Auto Scaling launches or terminates instances.
	LifecycleHookSpecificationList []types.LifecycleHookSpecification

	// A list of Classic Load Balancers associated with this Auto Scaling group. For
	// Application Load Balancers and Network Load Balancers, specify TargetGroupARNs
	// instead.
	LoadBalancerNames []string

	// The maximum amount of time, in seconds, that an instance can be in service. The
	// default is null. If specified, the value must be either 0 or a number equal to
	// or greater than 86,400 seconds (1 day). For more information, see Replacing Auto
	// Scaling instances based on maximum instance lifetime
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MaxInstanceLifetime *int32

	// An embedded object that specifies a mixed instances policy. The required
	// parameters must be specified. If optional parameters are unspecified, their
	// default values are used. The policy includes parameters that not only define the
	// distribution of On-Demand Instances and Spot Instances, the maximum price to pay
	// for Spot Instances, and how the Auto Scaling group allocates instance types to
	// fulfill On-Demand and Spot capacities, but also the parameters that specify the
	// instance configuration information—the launch template and instance types. The
	// policy can also include a weight for each instance type and different launch
	// templates for individual instance types. For more information, see Auto Scaling
	// groups with multiple instance types and purchase options
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html)
	// in the Amazon EC2 Auto Scaling User Guide. Conditional: You must specify either
	// a launch template (LaunchTemplate or MixedInstancesPolicy) or a launch
	// configuration (LaunchConfigurationName or InstanceId).
	MixedInstancesPolicy *types.MixedInstancesPolicy

	// Indicates whether newly launched instances are protected from termination by
	// Amazon EC2 Auto Scaling when scaling in. For more information about preventing
	// instances from terminating on scale in, see Instance scale-in protection
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection)
	// in the Amazon EC2 Auto Scaling User Guide.
	NewInstancesProtectedFromScaleIn *bool

	// The name of an existing placement group into which to launch your instances, if
	// any. A placement group is a logical grouping of instances within a single
	// Availability Zone. You cannot specify multiple Availability Zones and a
	// placement group. For more information, see Placement Groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	PlacementGroup *string

	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling
	// group uses to call other AWS services on your behalf. By default, Amazon EC2
	// Auto Scaling uses a service-linked role named AWSServiceRoleForAutoScaling,
	// which it creates if it does not exist. For more information, see Service-linked
	// roles
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	ServiceLinkedRoleARN *string

	// One or more tags. You can tag your Auto Scaling group and propagate the tags to
	// the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS
	// volumes. To add tags to Amazon EBS volumes, specify the tags in a launch
	// template but use caution. If the launch template specifies an instance tag with
	// a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling
	// overrides the value of that instance tag with the value specified by the Auto
	// Scaling group. For more information, see Tagging Auto Scaling groups and
	// instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	Tags []types.Tag

	// The Amazon Resource Names (ARN) of the target groups to associate with the Auto
	// Scaling group. Instances are registered as targets in a target group, and
	// traffic is routed to the target group. For more information, see Elastic Load
	// Balancing and Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TargetGroupARNs []string

	// A policy or a list of policies that are used to select the instance to
	// terminate. These policies are executed in the order that you list them. For more
	// information, see Controlling which Auto Scaling instances terminate during scale
	// in
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TerminationPolicies []string

	// A comma-separated list of subnet IDs for a virtual private cloud (VPC) where
	// instances in the Auto Scaling group can be created. If you specify
	// VPCZoneIdentifier with AvailabilityZones, the subnets that you specify for this
	// parameter must reside in those Availability Zones. Conditional: If your account
	// supports EC2-Classic and VPC, this parameter is required to launch instances
	// into a VPC.
	VPCZoneIdentifier *string
}

type CreateAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAutoScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAutoScalingGroup{}, middleware.After)
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
	if err = addOpCreateAutoScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoScalingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateAutoScalingGroup",
	}
}
