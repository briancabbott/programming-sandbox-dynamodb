// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates and defines the settings for a classification job.
func (c *Client) CreateClassificationJob(ctx context.Context, params *CreateClassificationJobInput, optFns ...func(*Options)) (*CreateClassificationJobOutput, error) {
	if params == nil {
		params = &CreateClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClassificationJob", params, optFns, addOperationCreateClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClassificationJobInput struct {

	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	//
	// This member is required.
	ClientToken *string

	// The schedule for running the job. Valid values are:
	//
	// * ONE_TIME - Run the job
	// only once. If you specify this value, don't specify a value for the
	// scheduleFrequency property.
	//
	// * SCHEDULED - Run the job on a daily, weekly, or
	// monthly basis. If you specify this value, use the scheduleFrequency property to
	// define the recurrence pattern for the job.
	//
	// This member is required.
	JobType types.JobType

	// A custom name for the job. The name can contain as many as 500 characters.
	//
	// This member is required.
	Name *string

	// The S3 buckets that contain the objects to analyze, and the scope of that
	// analysis.
	//
	// This member is required.
	S3JobDefinition *types.S3JobDefinition

	// The custom data identifiers to use for data analysis and classification.
	CustomDataIdentifierIds []string

	// A custom description of the job. The description can contain as many as 200
	// characters.
	Description *string

	// Specifies whether to analyze all existing, eligible objects immediately after
	// the job is created.
	InitialRun bool

	// The sampling depth, as a percentage, to apply when processing objects. This
	// value determines the percentage of eligible objects that the job analyzes. If
	// this value is less than 100, Amazon Macie selects the objects to analyze at
	// random, up to the specified percentage, and analyzes all the data in those
	// objects.
	SamplingPercentage int32

	// The recurrence pattern for running the job. To run the job only once, don't
	// specify a value for this property and set the value for the jobType property to
	// ONE_TIME.
	ScheduleFrequency *types.JobScheduleFrequency

	// A map of key-value pairs that specifies the tags to associate with the job. A
	// job can have a maximum of 50 tags. Each tag consists of a tag key and an
	// associated tag value. The maximum length of a tag key is 128 characters. The
	// maximum length of a tag value is 256 characters.
	Tags map[string]string
}

type CreateClassificationJobOutput struct {

	// The Amazon Resource Name (ARN) of the job.
	JobArn *string

	// The unique identifier for the job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateClassificationJob{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateClassificationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClassificationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateClassificationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateClassificationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateClassificationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClassificationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClassificationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateClassificationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateClassificationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateClassificationJob",
	}
}
