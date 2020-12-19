// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the properties associated with an InferICD10CM job. Use this operation to
// get the status of an inference job.
func (c *Client) DescribeICD10CMInferenceJob(ctx context.Context, params *DescribeICD10CMInferenceJobInput, optFns ...func(*Options)) (*DescribeICD10CMInferenceJobOutput, error) {
	if params == nil {
		params = &DescribeICD10CMInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeICD10CMInferenceJob", params, optFns, addOperationDescribeICD10CMInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeICD10CMInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeICD10CMInferenceJobInput struct {

	// The identifier that Amazon Comprehend Medical generated for the job. The
	// StartICD10CMInferenceJob operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribeICD10CMInferenceJobOutput struct {

	// An object that contains the properties associated with a detection job.
	ComprehendMedicalAsyncJobProperties *types.ComprehendMedicalAsyncJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeICD10CMInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeICD10CMInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeICD10CMInferenceJob{}, middleware.After)
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
	if err = addOpDescribeICD10CMInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeICD10CMInferenceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeICD10CMInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DescribeICD10CMInferenceJob",
	}
}
