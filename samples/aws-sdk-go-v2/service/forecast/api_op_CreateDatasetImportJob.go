// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports your training data to an Amazon Forecast dataset. You provide the
// location of your training data in an Amazon Simple Storage Service (Amazon S3)
// bucket and the Amazon Resource Name (ARN) of the dataset that you want to import
// the data to. You must specify a DataSource object that includes an AWS Identity
// and Access Management (IAM) role that Amazon Forecast can assume to access the
// data, as Amazon Forecast makes a copy of your data and processes it in an
// internal AWS system. For more information, see aws-forecast-iam-roles. The
// training data must be in CSV format. The delimiter must be a comma (,). You can
// specify the path to a specific CSV file, the S3 bucket, or to a folder in the S3
// bucket. For the latter two cases, Amazon Forecast imports all files up to the
// limit of 10,000 files. Because dataset imports are not aggregated, your most
// recent dataset import is the one that is used when training a predictor or
// generating a forecast. Make sure that your most recent dataset import contains
// all of the data you want to model off of, and not just the new data collected
// since the previous import. To get a list of all your dataset import jobs,
// filtered by specified criteria, use the ListDatasetImportJobs operation.
func (c *Client) CreateDatasetImportJob(ctx context.Context, params *CreateDatasetImportJobInput, optFns ...func(*Options)) (*CreateDatasetImportJobOutput, error) {
	if params == nil {
		params = &CreateDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatasetImportJob", params, optFns, addOperationCreateDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetImportJobInput struct {

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data. The
	// training data must be stored in an Amazon S3 bucket. If encryption is used,
	// DataSource must include an AWS Key Management Service (KMS) key and the IAM role
	// must allow Amazon Forecast permission to access the key. The KMS key and IAM
	// role must match those specified in the EncryptionConfig parameter of the
	// CreateDataset operation.
	//
	// This member is required.
	DataSource *types.DataSource

	// The Amazon Resource Name (ARN) of the Amazon Forecast dataset that you want to
	// import data to.
	//
	// This member is required.
	DatasetArn *string

	// The name for the dataset import job. We recommend including the current
	// timestamp in the name, for example, 20190721DatasetImport. This can help you
	// avoid getting a ResourceAlreadyExistsException exception.
	//
	// This member is required.
	DatasetImportJobName *string

	// The optional metadata that you apply to the dataset import job to help you
	// categorize and organize them. Each tag consists of a key and an optional value,
	// both of which you define. The following basic restrictions apply to tags:
	//
	// *
	// Maximum number of tags per resource - 50.
	//
	// * For each resource, each tag key
	// must be unique, and each tag key can have only one value.
	//
	// * Maximum key length
	// - 128 Unicode characters in UTF-8.
	//
	// * Maximum value length - 256 Unicode
	// characters in UTF-8.
	//
	// * If your tagging schema is used across multiple services
	// and resources, remember that other services may have restrictions on allowed
	// characters. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	// * Tag
	// keys and values are case sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or
	// lowercase combination of such as a prefix for keys as it is reserved for AWS
	// use. You cannot edit or delete tag keys with this prefix. Values can have this
	// prefix. If a tag value has aws as its prefix but the key does not, then Forecast
	// considers it to be a user tag and will count against the limit of 50 tags. Tags
	// with only the key prefix of aws do not count against your tags per resource
	// limit.
	Tags []types.Tag

	// The format of timestamps in the dataset. The format that you specify depends on
	// the DataFrequency specified when the dataset was created. The following formats
	// are supported
	//
	// * "yyyy-MM-dd" For the following data frequencies: Y, M, W, and
	// D
	//
	// * "yyyy-MM-dd HH:mm:ss" For the following data frequencies: H, 30min, 15min,
	// and 1min; and optionally, for: Y, M, W, and D
	//
	// If the format isn't specified,
	// Amazon Forecast expects the format to be "yyyy-MM-dd HH:mm:ss".
	TimestampFormat *string
}

type CreateDatasetImportJobOutput struct {

	// The Amazon Resource Name (ARN) of the dataset import job.
	DatasetImportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetImportJob{}, middleware.After)
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
	if err = addOpCreateDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateDatasetImportJob",
	}
}
