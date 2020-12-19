// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Obtains the manual snapshot limits for a directory.
func (c *Client) GetSnapshotLimits(ctx context.Context, params *GetSnapshotLimitsInput, optFns ...func(*Options)) (*GetSnapshotLimitsOutput, error) {
	if params == nil {
		params = &GetSnapshotLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshotLimits", params, optFns, addOperationGetSnapshotLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the GetSnapshotLimits operation.
type GetSnapshotLimitsInput struct {

	// Contains the identifier of the directory to obtain the limits for.
	//
	// This member is required.
	DirectoryId *string
}

// Contains the results of the GetSnapshotLimits operation.
type GetSnapshotLimitsOutput struct {

	// A SnapshotLimits object that contains the manual snapshot limits for the
	// specified directory.
	SnapshotLimits *types.SnapshotLimits

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSnapshotLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSnapshotLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSnapshotLimits{}, middleware.After)
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
	if err = addOpGetSnapshotLimitsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshotLimits(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSnapshotLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "GetSnapshotLimits",
	}
}
