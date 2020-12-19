// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves entries from the specified fleet's event log. You can specify a time
// range to limit the result set. Use the pagination parameters to retrieve results
// as a set of sequential pages. If successful, a collection of event log entries
// matching the request are returned. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// * Describe
// fleets:
//
// * DescribeFleetAttributes
//
// * DescribeFleetCapacity
//
// *
// DescribeFleetPortSettings
//
// * DescribeFleetUtilization
//
// *
// DescribeRuntimeConfiguration
//
// * DescribeEC2InstanceLimits
//
// *
// DescribeFleetEvents
//
// * UpdateFleetAttributes
//
// * StartFleetActions or
// StopFleetActions
func (c *Client) DescribeFleetEvents(ctx context.Context, params *DescribeFleetEventsInput, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) {
	if params == nil {
		params = &DescribeFleetEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetEvents", params, optFns, addOperationDescribeFleetEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeFleetEventsInput struct {

	// A unique identifier for a fleet to get event logs for. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Most recent date to retrieve event logs for. If no end time is specified, this
	// call returns entries from the specified start time up to the present. Format is
	// a number expressed in Unix time as milliseconds (ex: "1469498468.057").
	EndTime *time.Time

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Earliest date to retrieve event logs for. If no start time is specified, this
	// call returns entries starting from when the fleet was created to the specified
	// end time. Format is a number expressed in Unix time as milliseconds (ex:
	// "1469498468.057").
	StartTime *time.Time
}

// Represents the returned data in response to a request operation.
type DescribeFleetEventsOutput struct {

	// A collection of objects containing event log entries for the specified fleet.
	Events []types.Event

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFleetEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetEvents{}, middleware.After)
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
	if err = addOpDescribeFleetEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetEvents(options.Region), middleware.Before); err != nil {
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

// DescribeFleetEventsAPIClient is a client that implements the DescribeFleetEvents
// operation.
type DescribeFleetEventsAPIClient interface {
	DescribeFleetEvents(context.Context, *DescribeFleetEventsInput, ...func(*Options)) (*DescribeFleetEventsOutput, error)
}

var _ DescribeFleetEventsAPIClient = (*Client)(nil)

// DescribeFleetEventsPaginatorOptions is the paginator options for
// DescribeFleetEvents
type DescribeFleetEventsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetEventsPaginator is a paginator for DescribeFleetEvents
type DescribeFleetEventsPaginator struct {
	options   DescribeFleetEventsPaginatorOptions
	client    DescribeFleetEventsAPIClient
	params    *DescribeFleetEventsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetEventsPaginator returns a new DescribeFleetEventsPaginator
func NewDescribeFleetEventsPaginator(client DescribeFleetEventsAPIClient, params *DescribeFleetEventsInput, optFns ...func(*DescribeFleetEventsPaginatorOptions)) *DescribeFleetEventsPaginator {
	options := DescribeFleetEventsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeFleetEventsInput{}
	}

	return &DescribeFleetEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetEventsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeFleetEvents page.
func (p *DescribeFleetEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeFleetEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeFleetEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetEvents",
	}
}