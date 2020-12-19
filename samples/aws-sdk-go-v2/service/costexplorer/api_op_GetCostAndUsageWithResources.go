// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves cost and usage metrics with resources for your account. You can
// specify which cost and usage-related metric, such as BlendedCosts or
// UsageQuantity, that you want the request to return. You can also filter and
// group your data by various dimensions, such as SERVICE or AZ, in a specific time
// range. For a complete list of valid dimensions, see the GetDimensionValues
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html)
// operation. Management account in an organization in AWS Organizations have
// access to all member accounts. This API is currently available for the Amazon
// Elastic Compute Cloud – Compute service only. This is an opt-in only feature.
// You can enable this feature from the Cost Explorer Settings page. For
// information on how to access the Settings page, see Controlling Access for Cost
// Explorer
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-access.html) in
// the AWS Billing and Cost Management User Guide.
func (c *Client) GetCostAndUsageWithResources(ctx context.Context, params *GetCostAndUsageWithResourcesInput, optFns ...func(*Options)) (*GetCostAndUsageWithResourcesOutput, error) {
	if params == nil {
		params = &GetCostAndUsageWithResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostAndUsageWithResources", params, optFns, addOperationGetCostAndUsageWithResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostAndUsageWithResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostAndUsageWithResourcesInput struct {

	// Filters Amazon Web Services costs by different dimensions. For example, you can
	// specify SERVICE and LINKED_ACCOUNT and get the costs that are associated with
	// that account's usage of that service. You can nest Expression objects to define
	// any combination of dimension filters. For more information, see Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html).
	// The GetCostAndUsageWithResources operation requires that you either group by or
	// filter by a ResourceId. It requires the Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)"SERVICE
	// = Amazon Elastic Compute Cloud - Compute" in the filter.
	//
	// This member is required.
	Filter *types.Expression

	// Sets the start and end dates for retrieving Amazon Web Services costs. The range
	// must be within the last 14 days (the start date cannot be earlier than 14 days
	// ago). The start date is inclusive, but the end date is exclusive. For example,
	// if start is 2017-01-01 and end is 2017-05-01, then the cost and usage data is
	// retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Sets the AWS cost granularity to MONTHLY, DAILY, or HOURLY. If Granularity isn't
	// set, the response object doesn't include the Granularity, MONTHLY, DAILY, or
	// HOURLY.
	Granularity types.Granularity

	// You can group Amazon Web Services costs using up to two different groups:
	// DIMENSION, TAG, COST_CATEGORY.
	GroupBy []types.GroupDefinition

	// Which metrics are returned in the query. For more information about blended and
	// unblended rates, see Why does the "blended" annotation appear on some line items
	// in my bill?
	// (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	// Valid values are AmortizedCost, BlendedCost, NetAmortizedCost, NetUnblendedCost,
	// NormalizedUsageAmount, UnblendedCost, and UsageQuantity. If you return the
	// UsageQuantity metric, the service aggregates all usage numbers without taking
	// the units into account. For example, if you aggregate usageQuantity across all
	// of Amazon EC2, the results aren't meaningful because Amazon EC2 compute hours
	// and data transfer are measured in different units (for example, hours vs. GB).
	// To get more meaningful UsageQuantity metrics, filter by UsageType or
	// UsageTypeGroups. Metrics is required for GetCostAndUsageWithResources requests.
	Metrics []string

	// The token to retrieve the next set of results. AWS provides the token when the
	// response from a previous call has more results than the maximum page size.
	NextPageToken *string
}

type GetCostAndUsageWithResourcesOutput struct {

	// The groups that are specified by the Filter or GroupBy parameters in the
	// request.
	GroupDefinitions []types.GroupDefinition

	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string

	// The time period that is covered by the results in the response.
	ResultsByTime []types.ResultByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCostAndUsageWithResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostAndUsageWithResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostAndUsageWithResources{}, middleware.After)
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
	if err = addOpGetCostAndUsageWithResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostAndUsageWithResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostAndUsageWithResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetCostAndUsageWithResources",
	}
}
