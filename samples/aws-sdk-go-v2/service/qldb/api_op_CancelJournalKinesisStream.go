// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Ends a given Amazon QLDB journal stream. Before a stream can be canceled, its
// current status must be ACTIVE. You can't restart a stream after you cancel it.
// Canceled QLDB stream resources are subject to a 7-day retention period, so they
// are automatically deleted after this limit expires.
func (c *Client) CancelJournalKinesisStream(ctx context.Context, params *CancelJournalKinesisStreamInput, optFns ...func(*Options)) (*CancelJournalKinesisStreamOutput, error) {
	if params == nil {
		params = &CancelJournalKinesisStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJournalKinesisStream", params, optFns, addOperationCancelJournalKinesisStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJournalKinesisStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJournalKinesisStreamInput struct {

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The unique ID that QLDB assigns to each QLDB journal stream.
	//
	// This member is required.
	StreamId *string
}

type CancelJournalKinesisStreamOutput struct {

	// The unique ID that QLDB assigns to each QLDB journal stream.
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelJournalKinesisStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJournalKinesisStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJournalKinesisStream{}, middleware.After)
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
	if err = addOpCancelJournalKinesisStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJournalKinesisStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelJournalKinesisStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "CancelJournalKinesisStream",
	}
}
