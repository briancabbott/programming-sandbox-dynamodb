// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how timestamp request and response headers are serialized.
func (c *Client) TimestampFormatHeaders(ctx context.Context, params *TimestampFormatHeadersInput, optFns ...func(*Options)) (*TimestampFormatHeadersOutput, error) {
	if params == nil {
		params = &TimestampFormatHeadersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TimestampFormatHeaders", params, optFns, addOperationTimestampFormatHeadersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TimestampFormatHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TimestampFormatHeadersInput struct {
	DefaultFormat *time.Time

	MemberDateTime *time.Time

	MemberEpochSeconds *time.Time

	MemberHttpDate *time.Time

	TargetDateTime *time.Time

	TargetEpochSeconds *time.Time

	TargetHttpDate *time.Time
}

type TimestampFormatHeadersOutput struct {
	DefaultFormat *time.Time

	MemberDateTime *time.Time

	MemberEpochSeconds *time.Time

	MemberHttpDate *time.Time

	TargetDateTime *time.Time

	TargetEpochSeconds *time.Time

	TargetHttpDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTimestampFormatHeadersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpTimestampFormatHeaders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpTimestampFormatHeaders{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTimestampFormatHeaders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTimestampFormatHeaders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TimestampFormatHeaders",
	}
}
