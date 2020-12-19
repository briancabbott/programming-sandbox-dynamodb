// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the settings of a data set.
func (c *Client) UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) {
	if params == nil {
		params = &UpdateDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataset", params, optFns, addOperationUpdateDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetInput struct {

	// A list of DatasetAction objects.
	//
	// This member is required.
	Actions []types.DatasetAction

	// The name of the data set to update.
	//
	// This member is required.
	DatasetName *string

	// When dataset contents are created, they are delivered to destinations specified
	// here.
	ContentDeliveryRules []types.DatasetContentDeliveryRule

	// A list of data rules that send notifications to Amazon CloudWatch, when data
	// arrives late. To specify lateDataRules, the dataset must use a DeltaTimer
	// (https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html)
	// filter.
	LateDataRules []types.LateDataRule

	// How long, in days, dataset contents are kept for the dataset.
	RetentionPeriod *types.RetentionPeriod

	// A list of DatasetTrigger objects. The list can be empty or can contain up to
	// five DatasetTrigger objects.
	Triggers []types.DatasetTrigger

	// Optional. How many versions of dataset contents are kept. If not specified or
	// set to null, only the latest version plus the latest succeeded version (if they
	// are different) are kept for the time period specified by the retentionPeriod
	// parameter. For more information, see Keeping Multiple Versions of AWS IoT
	// Analytics Data Sets
	// (https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	// in the AWS IoT Analytics User Guide.
	VersioningConfiguration *types.VersioningConfiguration
}

type UpdateDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataset{}, middleware.After)
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
	if err = addOpUpdateDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "UpdateDataset",
	}
}
