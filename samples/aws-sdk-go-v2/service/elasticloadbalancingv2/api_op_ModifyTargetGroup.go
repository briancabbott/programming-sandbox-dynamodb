// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the health checks used when evaluating the health state of the targets
// in the specified target group.
func (c *Client) ModifyTargetGroup(ctx context.Context, params *ModifyTargetGroupInput, optFns ...func(*Options)) (*ModifyTargetGroupOutput, error) {
	if params == nil {
		params = &ModifyTargetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTargetGroup", params, optFns, addOperationModifyTargetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTargetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTargetGroupInput struct {

	// The Amazon Resource Name (ARN) of the target group.
	//
	// This member is required.
	TargetGroupArn *string

	// Indicates whether health checks are enabled.
	HealthCheckEnabled *bool

	// The approximate amount of time, in seconds, between health checks of an
	// individual target. For TCP health checks, the supported values are 10 or 30
	// seconds. With Network Load Balancers, you can't modify this setting.
	HealthCheckIntervalSeconds *int32

	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC
	// protocol version] The path of a custom health check method with the format
	// /package.service/method. The default is /AWS.ALB/healthcheck.
	HealthCheckPath *string

	// The port the load balancer uses when performing health checks on targets.
	HealthCheckPort *string

	// The protocol the load balancer uses when performing health checks on targets.
	// The TCP protocol is supported for health checks only if the protocol of the
	// target group is TCP, TLS, UDP, or TCP_UDP. The GENEVE, TLS, UDP, and TCP_UDP
	// protocols are not supported for health checks. With Network Load Balancers, you
	// can't modify this setting.
	HealthCheckProtocol types.ProtocolEnum

	// [HTTP/HTTPS health checks] The amount of time, in seconds, during which no
	// response means a failed health check. With Network Load Balancers, you can't
	// modify this setting.
	HealthCheckTimeoutSeconds *int32

	// The number of consecutive health checks successes required before considering an
	// unhealthy target healthy.
	HealthyThresholdCount *int32

	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a
	// successful response from a target. With Network Load Balancers, you can't modify
	// this setting.
	Matcher *types.Matcher

	// The number of consecutive health check failures required before considering the
	// target unhealthy. For target groups with a protocol of TCP or TLS, this value
	// must be the same as the healthy threshold count.
	UnhealthyThresholdCount *int32
}

type ModifyTargetGroupOutput struct {

	// Information about the modified target group.
	TargetGroups []types.TargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyTargetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyTargetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyTargetGroup{}, middleware.After)
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
	if err = addOpModifyTargetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTargetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTargetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyTargetGroup",
	}
}
