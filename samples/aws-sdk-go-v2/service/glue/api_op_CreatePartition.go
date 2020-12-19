// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new partition.
func (c *Client) CreatePartition(ctx context.Context, params *CreatePartitionInput, optFns ...func(*Options)) (*CreatePartitionOutput, error) {
	if params == nil {
		params = &CreatePartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePartition", params, optFns, addOperationCreatePartitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePartitionInput struct {

	// The name of the metadata database in which the partition is to be created.
	//
	// This member is required.
	DatabaseName *string

	// A PartitionInput structure defining the partition to be created.
	//
	// This member is required.
	PartitionInput *types.PartitionInput

	// The name of the metadata table in which the partition is to be created.
	//
	// This member is required.
	TableName *string

	// The AWS account ID of the catalog in which the partition is to be created.
	CatalogId *string
}

type CreatePartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePartitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePartition{}, middleware.After)
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
	if err = addOpCreatePartitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePartition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePartition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreatePartition",
	}
}
