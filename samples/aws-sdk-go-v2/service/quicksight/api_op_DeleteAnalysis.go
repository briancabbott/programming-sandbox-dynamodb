// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes an analysis from Amazon QuickSight. You can optionally include a
// recovery window during which you can restore the analysis. If you don't specify
// a recovery window value, the operation defaults to 30 days. QuickSight attaches
// a DeletionTime stamp to the response that specifies the end of the recovery
// window. At the end of the recovery window, QuickSight deletes the analysis
// permanently. At any time before recovery window ends, you can use the
// RestoreAnalysis API operation to remove the DeletionTime stamp and cancel the
// deletion of the analysis. The analysis remains visible in the API until it's
// deleted, so you can describe it but you can't make a template from it. An
// analysis that's scheduled for deletion isn't accessible in the QuickSight
// console. To access it in the console, restore it. Deleting an analysis doesn't
// delete the dashboards that you publish from it.
func (c *Client) DeleteAnalysis(ctx context.Context, params *DeleteAnalysisInput, optFns ...func(*Options)) (*DeleteAnalysisOutput, error) {
	if params == nil {
		params = &DeleteAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnalysis", params, optFns, addOperationDeleteAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAnalysisInput struct {

	// The ID of the analysis that you're deleting.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the AWS account where you want to delete an analysis.
	//
	// This member is required.
	AwsAccountId *string

	// This option defaults to the value NoForceDeleteWithoutRecovery. To immediately
	// delete the analysis, add the ForceDeleteWithoutRecovery option. You can't
	// restore an analysis after it's deleted.
	ForceDeleteWithoutRecovery bool

	// A value that specifies the number of days that QuickSight waits before it
	// deletes the analysis. You can't use this parameter with the
	// ForceDeleteWithoutRecovery option in the same API call. The default value is 30.
	RecoveryWindowInDays *int64
}

type DeleteAnalysisOutput struct {

	// The ID of the deleted analysis.
	AnalysisId *string

	// The Amazon Resource Name (ARN) of the deleted analysis.
	Arn *string

	// The date and time that the analysis is scheduled to be deleted.
	DeletionTime *time.Time

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAnalysis{}, middleware.After)
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
	if err = addOpDeleteAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteAnalysis",
	}
}
