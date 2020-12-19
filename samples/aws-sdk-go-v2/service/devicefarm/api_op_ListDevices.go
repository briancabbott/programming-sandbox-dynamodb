// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about unique device types.
func (c *Client) ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) {
	if params == nil {
		params = &ListDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevices", params, optFns, addOperationListDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the result of a list devices request.
type ListDevicesInput struct {

	// The Amazon Resource Name (ARN) of the project.
	Arn *string

	// Used to select a set of devices. A filter is made up of an attribute, an
	// operator, and one or more values.
	//
	// * Attribute: The aspect of a device such as
	// platform or model used as the selection criteria in a device filter. Allowed
	// values include:
	//
	// * ARN: The Amazon Resource Name (ARN) of the device (for
	// example, arn:aws:devicefarm:us-west-2::device:12345Example).
	//
	// * PLATFORM: The
	// device platform. Valid values are ANDROID or IOS.
	//
	// * OS_VERSION: The operating
	// system version (for example, 10.3.2).
	//
	// * MODEL: The device model (for example,
	// iPad 5th Gen).
	//
	// * AVAILABILITY: The current availability of the device. Valid
	// values are AVAILABLE, HIGHLY_AVAILABLE, BUSY, or TEMPORARY_NOT_AVAILABLE.
	//
	// *
	// FORM_FACTOR: The device form factor. Valid values are PHONE or TABLET.
	//
	// *
	// MANUFACTURER: The device manufacturer (for example, Apple).
	//
	// *
	// REMOTE_ACCESS_ENABLED: Whether the device is enabled for remote access. Valid
	// values are TRUE or FALSE.
	//
	// * REMOTE_DEBUG_ENABLED: Whether the device is enabled
	// for remote debugging. Valid values are TRUE or FALSE. Because remote debugging
	// is no longer supported
	// (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html),
	// this attribute is ignored.
	//
	// * INSTANCE_ARN: The Amazon Resource Name (ARN) of
	// the device instance.
	//
	// * INSTANCE_LABELS: The label of the device instance.
	//
	// *
	// FLEET_TYPE: The fleet type. Valid values are PUBLIC or PRIVATE.
	//
	// * Operator: The
	// filter operator.
	//
	// * The EQUALS operator is available for every attribute except
	// INSTANCE_LABELS.
	//
	// * The CONTAINS operator is available for the INSTANCE_LABELS
	// and MODEL attributes.
	//
	// * The IN and NOT_IN operators are available for the ARN,
	// OS_VERSION, MODEL, MANUFACTURER, and INSTANCE_ARN attributes.
	//
	// * The LESS_THAN,
	// GREATER_THAN, LESS_THAN_OR_EQUALS, and GREATER_THAN_OR_EQUALS operators are also
	// available for the OS_VERSION attribute.
	//
	// * Values: An array of one or more
	// filter values.
	//
	// * The IN and NOT_IN operators take a values array that has one
	// or more elements.
	//
	// * The other operators require an array with a single
	// element.
	//
	// * In a request, the AVAILABILITY attribute takes the following values:
	// AVAILABLE, HIGHLY_AVAILABLE, BUSY, or TEMPORARY_NOT_AVAILABLE.
	Filters []types.DeviceFilter

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list devices operation.
type ListDevicesOutput struct {

	// Information about the devices.
	Devices []types.Device

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDevices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevices(options.Region), middleware.Before); err != nil {
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

// ListDevicesAPIClient is a client that implements the ListDevices operation.
type ListDevicesAPIClient interface {
	ListDevices(context.Context, *ListDevicesInput, ...func(*Options)) (*ListDevicesOutput, error)
}

var _ ListDevicesAPIClient = (*Client)(nil)

// ListDevicesPaginatorOptions is the paginator options for ListDevices
type ListDevicesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevicesPaginator is a paginator for ListDevices
type ListDevicesPaginator struct {
	options   ListDevicesPaginatorOptions
	client    ListDevicesAPIClient
	params    *ListDevicesInput
	nextToken *string
	firstPage bool
}

// NewListDevicesPaginator returns a new ListDevicesPaginator
func NewListDevicesPaginator(client ListDevicesAPIClient, params *ListDevicesInput, optFns ...func(*ListDevicesPaginatorOptions)) *ListDevicesPaginator {
	options := ListDevicesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDevicesInput{}
	}

	return &ListDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDevices page.
func (p *ListDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListDevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListDevices",
	}
}
