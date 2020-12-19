// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your clusters.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusters", params, optFns, addOperationDescribeClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {

	// A list of up to 100 cluster names or full cluster Amazon Resource Name (ARN)
	// entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []string

	// Whether to include additional information about your clusters in the response.
	// If this field is omitted, the attachments, statistics, and tags are not
	// included. If ATTACHMENTS is specified, the attachments for the container
	// instances or tasks within the cluster are included. If SETTINGS is specified,
	// the settings for the cluster are included. If STATISTICS is specified, the
	// following additional information, separated by launch type, is included:
	//
	// *
	// runningEC2TasksCount
	//
	// * runningFargateTasksCount
	//
	// * pendingEC2TasksCount
	//
	// *
	// pendingFargateTasksCount
	//
	// * activeEC2ServiceCount
	//
	// *
	// activeFargateServiceCount
	//
	// * drainingEC2ServiceCount
	//
	// *
	// drainingFargateServiceCount
	//
	// If TAGS is specified, the metadata tags associated
	// with the cluster are included.
	Include []types.ClusterField
}

type DescribeClustersOutput struct {

	// The list of clusters.
	Clusters []types.Cluster

	// Any failures associated with the call.
	Failures []types.Failure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeClusters",
	}
}