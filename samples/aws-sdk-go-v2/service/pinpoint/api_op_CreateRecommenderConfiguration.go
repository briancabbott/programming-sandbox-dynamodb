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

// Creates an Amazon Pinpoint configuration for a recommender model.
func (c *Client) CreateRecommenderConfiguration(ctx context.Context, params *CreateRecommenderConfigurationInput, optFns ...func(*Options)) (*CreateRecommenderConfigurationOutput, error) {
	if params == nil {
		params = &CreateRecommenderConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRecommenderConfiguration", params, optFns, addOperationCreateRecommenderConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRecommenderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRecommenderConfigurationInput struct {

	// Specifies Amazon Pinpoint configuration settings for retrieving and processing
	// recommendation data from a recommender model.
	//
	// This member is required.
	CreateRecommenderConfiguration *types.CreateRecommenderConfigurationShape
}

type CreateRecommenderConfigurationOutput struct {

	// Provides information about Amazon Pinpoint configuration settings for retrieving
	// and processing data from a recommender model.
	//
	// This member is required.
	RecommenderConfigurationResponse *types.RecommenderConfigurationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRecommenderConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRecommenderConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRecommenderConfiguration{}, middleware.After)
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
	if err = addOpCreateRecommenderConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRecommenderConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRecommenderConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "CreateRecommenderConfiguration",
	}
}
