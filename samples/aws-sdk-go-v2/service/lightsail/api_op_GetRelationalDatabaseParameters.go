// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the runtime parameters offered by the underlying database
// software, or engine, for a specific database in Amazon Lightsail. In addition to
// the parameter names and values, this operation returns other information about
// each parameter. This information includes whether changes require a reboot,
// whether the parameter is modifiable, the allowed values, and the data types.
func (c *Client) GetRelationalDatabaseParameters(ctx context.Context, params *GetRelationalDatabaseParametersInput, optFns ...func(*Options)) (*GetRelationalDatabaseParametersOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseParameters", params, optFns, addOperationGetRelationalDatabaseParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseParametersInput struct {

	// The name of your database for which to get parameters.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseParameters request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabaseParametersOutput struct {

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseParameters request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// An object describing the result of your get relational database parameters
	// request.
	Parameters []types.RelationalDatabaseParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseParameters{}, middleware.After)
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
	if err = addOpGetRelationalDatabaseParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseParameters",
	}
}