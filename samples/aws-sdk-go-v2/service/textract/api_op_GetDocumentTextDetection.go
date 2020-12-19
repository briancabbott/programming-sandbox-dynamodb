// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the results for an Amazon Textract asynchronous operation that detects text
// in a document. Amazon Textract can detect lines of text and the words that make
// up a line of text. You start asynchronous text detection by calling
// StartDocumentTextDetection, which returns a job identifier (JobId). When the
// text detection operation finishes, Amazon Textract publishes a completion status
// to the Amazon Simple Notification Service (Amazon SNS) topic that's registered
// in the initial call to StartDocumentTextDetection. To get the results of the
// text-detection operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If so, call GetDocumentTextDetection, and pass
// the job identifier (JobId) from the initial call to StartDocumentTextDetection.
// GetDocumentTextDetection returns an array of Block objects. Each document page
// has as an associated Block of type PAGE. Each PAGE Block object is the parent of
// LINE Block objects that represent the lines of detected text on a page. A LINE
// Block object is a parent for each word that makes up the line. Words are
// represented by Block objects of type WORD. Use the MaxResults parameter to limit
// the number of blocks that are returned. If there are more results than specified
// in MaxResults, the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetDocumentTextDetection, and populate the NextToken request
// parameter with the token value that's returned from the previous call to
// GetDocumentTextDetection. For more information, see Document Text Detection
// (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html).
func (c *Client) GetDocumentTextDetection(ctx context.Context, params *GetDocumentTextDetectionInput, optFns ...func(*Options)) (*GetDocumentTextDetectionOutput, error) {
	if params == nil {
		params = &GetDocumentTextDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentTextDetection", params, optFns, addOperationGetDocumentTextDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentTextDetectionInput struct {

	// A unique identifier for the text detection job. The JobId is returned from
	// StartDocumentTextDetection. A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 1,000. If you specify a value greater than 1,000, a maximum
	// of 1,000 results is returned. The default value is 1,000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more blocks to
	// retrieve), Amazon Textract returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of blocks.
	NextToken *string
}

type GetDocumentTextDetectionOutput struct {

	// The results of the text-detection operation.
	Blocks []types.Block

	//
	DetectDocumentTextModelVersion *string

	// Information about a document that Amazon Textract processed. DocumentMetadata is
	// returned in every page of paginated responses from an Amazon Textract video
	// operation.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the text detection job.
	JobStatus types.JobStatus

	// If the response is truncated, Amazon Textract returns this token. You can use
	// this token in the subsequent request to retrieve the next set of text-detection
	// results.
	NextToken *string

	// Returns if the detection job could not be completed. Contains explanation for
	// what error occured.
	StatusMessage *string

	// A list of warnings that occurred during the text-detection operation for the
	// document.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDocumentTextDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDocumentTextDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDocumentTextDetection{}, middleware.After)
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
	if err = addOpGetDocumentTextDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentTextDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDocumentTextDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetDocumentTextDetection",
	}
}
