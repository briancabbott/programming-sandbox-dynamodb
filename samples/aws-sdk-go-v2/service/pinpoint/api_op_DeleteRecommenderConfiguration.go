// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon Pinpoint configuration for a recommender model.
func (c *Client) DeleteRecommenderConfiguration(ctx context.Context, params *DeleteRecommenderConfigurationInput, optFns ...func(*Options)) (*DeleteRecommenderConfigurationOutput, error) {
	if params == nil {
		params = &DeleteRecommenderConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRecommenderConfiguration", params, optFns, addOperationDeleteRecommenderConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRecommenderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRecommenderConfigurationInput struct {

	// The unique identifier for the recommender model configuration. This identifier
	// is displayed as the Recommender ID on the Amazon Pinpoint console.
	//
	// This member is required.
	RecommenderId *string
}

type DeleteRecommenderConfigurationOutput struct {

	// Provides information about Amazon Pinpoint configuration settings for retrieving
	// and processing data from a recommender model.
	//
	// This member is required.
	RecommenderConfigurationResponse *types.RecommenderConfigurationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteRecommenderConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRecommenderConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRecommenderConfiguration{}, middleware.After)
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
	if err = addOpDeleteRecommenderConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRecommenderConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRecommenderConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteRecommenderConfiguration",
	}
}
