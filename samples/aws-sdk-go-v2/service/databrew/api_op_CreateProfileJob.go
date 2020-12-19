// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new job to profile an AWS Glue DataBrew dataset that exists in the
// current AWS account.
func (c *Client) CreateProfileJob(ctx context.Context, params *CreateProfileJobInput, optFns ...func(*Options)) (*CreateProfileJobOutput, error) {
	if params == nil {
		params = &CreateProfileJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfileJob", params, optFns, addOperationCreateProfileJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileJobInput struct {

	// The name of the dataset that this job is to act upon.
	//
	// This member is required.
	DatasetName *string

	// The name of the job to be created.
	//
	// This member is required.
	Name *string

	// An Amazon S3 location (bucket name an object key) where DataBrew can read input
	// data, or write output from a job.
	//
	// This member is required.
	OutputLocation *types.S3Location

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role to be assumed for this request.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job.
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//
	// * SSE-KMS -
	// para>SSE-KMS - server-side encryption with AWS KMS-managed keys.
	//
	// * SSE-S3 -
	// Server-side encryption with keys managed by Amazon S3.
	EncryptionMode types.EncryptionMode

	// A value that enables or disables Amazon CloudWatch logging for the current AWS
	// account. If logging is enabled, CloudWatch writes one log stream for each job
	// run.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that DataBrew can use when the job processes data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// Metadata tags to apply to this job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32
}

type CreateProfileJobOutput struct {

	// The name of the job that was created.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProfileJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfileJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfileJob{}, middleware.After)
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
	if err = addOpCreateProfileJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfileJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProfileJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "CreateProfileJob",
	}
}
