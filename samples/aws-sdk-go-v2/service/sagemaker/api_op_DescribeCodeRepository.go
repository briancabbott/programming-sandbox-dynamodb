// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets details about the specified Git repository.
func (c *Client) DescribeCodeRepository(ctx context.Context, params *DescribeCodeRepositoryInput, optFns ...func(*Options)) (*DescribeCodeRepositoryOutput, error) {
	if params == nil {
		params = &DescribeCodeRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCodeRepository", params, optFns, addOperationDescribeCodeRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCodeRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeRepositoryInput struct {

	// The name of the Git repository to describe.
	//
	// This member is required.
	CodeRepositoryName *string
}

type DescribeCodeRepositoryOutput struct {

	// The Amazon Resource Name (ARN) of the Git repository.
	//
	// This member is required.
	CodeRepositoryArn *string

	// The name of the Git repository.
	//
	// This member is required.
	CodeRepositoryName *string

	// The date and time that the repository was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The date and time that the repository was last changed.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Configuration details about the repository, including the URL where the
	// repository is located, the default branch, and the Amazon Resource Name (ARN) of
	// the AWS Secrets Manager secret that contains the credentials used to access the
	// repository.
	GitConfig *types.GitConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCodeRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCodeRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCodeRepository{}, middleware.After)
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
	if err = addOpDescribeCodeRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCodeRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeCodeRepository",
	}
}
