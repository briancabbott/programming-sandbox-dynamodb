// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about one or more specified events for one or more
// accounts in your organization. Information includes standard event data (Region,
// service, and so on, as returned by DescribeEventsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventsForOrganization.html)),
// a detailed event description, and possible additional metadata that depends upon
// the nature of the event. Affected entities are not included; to retrieve those,
// use the DescribeAffectedEntitiesForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operation. Before you can call this operation, you must first enable AWS Health
// to work with AWS Organizations. To do this, call the
// EnableHealthServiceAccessForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's master account. When you call the
// DescribeEventDetailsForOrganization operation, you specify the
// organizationEventDetailFilters object in the request. Depending on the AWS
// Health event type, note the following differences:
//
// * If the event is public,
// the awsAccountId parameter must be empty. If you specify an account ID for a
// public event, then an error message is returned. That's because the event might
// apply to all AWS accounts and isn't specific to an account in your
// organization.
//
// * If the event is specific to an account, then you must specify
// the awsAccountId parameter in the request. If you don't specify an account ID,
// an error message returns because the event is specific to an AWS account in your
// organization.
//
// For more information, see Event
// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html).
func (c *Client) DescribeEventDetailsForOrganization(ctx context.Context, params *DescribeEventDetailsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventDetailsForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeEventDetailsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventDetailsForOrganization", params, optFns, addOperationDescribeEventDetailsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventDetailsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventDetailsForOrganizationInput struct {

	// A set of JSON elements that includes the awsAccountId and the eventArn.
	//
	// This member is required.
	OrganizationEventDetailFilters []types.EventAccountFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string
}

type DescribeEventDetailsForOrganizationOutput struct {

	// Error messages for any events that could not be retrieved.
	FailedSet []types.OrganizationEventDetailsErrorItem

	// Information about the events that could be retrieved.
	SuccessfulSet []types.OrganizationEventDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventDetailsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventDetailsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventDetailsForOrganization{}, middleware.After)
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
	if err = addOpDescribeEventDetailsForOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventDetailsForOrganization",
	}
}