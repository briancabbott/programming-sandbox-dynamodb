// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new RTMP distribution. An RTMP distribution is similar to a web
// distribution, but an RTMP distribution streams media files using the Adobe
// Real-Time Messaging Protocol (RTMP) instead of serving files using HTTP. To
// create a new distribution, submit a POST request to the CloudFront API
// version/distribution resource. The request body must include a document with a
// StreamingDistributionConfig element. The response echoes the
// StreamingDistributionConfig element and returns other information about the RTMP
// distribution. To get the status of your request, use the GET
// StreamingDistribution API action. When the value of Enabled is true and the
// value of Status is Deployed, your distribution is ready. A distribution usually
// deploys in less than 15 minutes. For more information about web distributions,
// see Working with RTMP Distributions
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-rtmp.html)
// in the Amazon CloudFront Developer Guide. Beginning with the 2012-05-05 version
// of the CloudFront API, we made substantial changes to the format of the XML
// document that you include in the request body when you create or update a web
// distribution or an RTMP distribution, and when you invalidate objects. With
// previous versions of the API, we discovered that it was too easy to accidentally
// delete one or more values for an element that accepts multiple values, for
// example, CNAMEs and trusted signers. Our changes for the 2012-05-05 release are
// intended to prevent these accidental deletions and to notify you when there's a
// mismatch between the number of values you say you're specifying in the Quantity
// element and the number of values specified.
func (c *Client) CreateStreamingDistribution(ctx context.Context, params *CreateStreamingDistributionInput, optFns ...func(*Options)) (*CreateStreamingDistributionOutput, error) {
	if params == nil {
		params = &CreateStreamingDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamingDistribution", params, optFns, addOperationCreateStreamingDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new streaming distribution.
type CreateStreamingDistributionInput struct {

	// The streaming distribution's configuration information.
	//
	// This member is required.
	StreamingDistributionConfig *types.StreamingDistributionConfig
}

// The returned result of the corresponding request.
type CreateStreamingDistributionOutput struct {

	// The current version of the streaming distribution created.
	ETag *string

	// The fully qualified URI of the new streaming distribution resource just created.
	Location *string

	// The streaming distribution's information.
	StreamingDistribution *types.StreamingDistribution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStreamingDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateStreamingDistribution{}, middleware.After)
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
	if err = addOpCreateStreamingDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamingDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamingDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateStreamingDistribution",
	}
}
