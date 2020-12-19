// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns attributes related to AWS Elastic Beanstalk that are associated with the
// calling AWS account. The result currently has one set of attributes—resource
// quotas.
func (c *Client) DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) {
	if params == nil {
		params = &DescribeAccountAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAttributes", params, optFns, addOperationDescribeAccountAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAttributesInput struct {
}

type DescribeAccountAttributesOutput struct {

	// The Elastic Beanstalk resource quotas associated with the calling AWS account.
	ResourceQuotas *types.ResourceQuotas

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAccountAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAccountAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeAccountAttributes",
	}
}
