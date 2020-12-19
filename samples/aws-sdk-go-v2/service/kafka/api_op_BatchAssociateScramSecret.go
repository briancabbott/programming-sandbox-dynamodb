// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates one or more Scram Secrets with an Amazon MSK cluster.
func (c *Client) BatchAssociateScramSecret(ctx context.Context, params *BatchAssociateScramSecretInput, optFns ...func(*Options)) (*BatchAssociateScramSecretOutput, error) {
	if params == nil {
		params = &BatchAssociateScramSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateScramSecret", params, optFns, addOperationBatchAssociateScramSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateScramSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Associates sasl scram secrets to cluster.
type BatchAssociateScramSecretInput struct {

	// The Amazon Resource Name (ARN) of the cluster to be updated.
	//
	// This member is required.
	ClusterArn *string

	// List of AWS Secrets Manager secret ARNs.
	//
	// This member is required.
	SecretArnList []string
}

type BatchAssociateScramSecretOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// List of errors when associating secrets to cluster.
	UnprocessedScramSecrets []types.UnprocessedScramSecret

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchAssociateScramSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAssociateScramSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAssociateScramSecret{}, middleware.After)
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
	if err = addOpBatchAssociateScramSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateScramSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateScramSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "BatchAssociateScramSecret",
	}
}
