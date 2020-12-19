// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the contents of the specified Systems Manager document.
func (c *Client) GetDocument(ctx context.Context, params *GetDocumentInput, optFns ...func(*Options)) (*GetDocumentOutput, error) {
	if params == nil {
		params = &GetDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocument", params, optFns, addOperationGetDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentInput struct {

	// The name of the Systems Manager document.
	//
	// This member is required.
	Name *string

	// Returns the document in the specified format. The document format can be either
	// JSON or YAML. JSON is the default format.
	DocumentFormat types.DocumentFormat

	// The document version for which you want information.
	DocumentVersion *string

	// An optional field specifying the version of the artifact associated with the
	// document. For example, "Release 12, Update 6". This value is unique across all
	// versions of a document and can't be changed.
	VersionName *string
}

type GetDocumentOutput struct {

	// A description of the document attachments, including names, locations, sizes,
	// and so on.
	AttachmentsContent []types.AttachmentContent

	// The contents of the Systems Manager document.
	Content *string

	// The document format, either JSON or YAML.
	DocumentFormat types.DocumentFormat

	// The document type.
	DocumentType types.DocumentType

	// The document version.
	DocumentVersion *string

	// The name of the Systems Manager document.
	Name *string

	// A list of SSM documents required by a document. For example, an
	// ApplicationConfiguration document requires an ApplicationConfigurationSchema
	// document.
	Requires []types.DocumentRequires

	// The status of the Systems Manager document, such as Creating, Active, Updating,
	// Failed, and Deleting.
	Status types.DocumentStatus

	// A message returned by AWS Systems Manager that explains the Status value. For
	// example, a Failed status might be explained by the StatusInformation message,
	// "The specified S3 bucket does not exist. Verify that the URL of the S3 bucket is
	// correct."
	StatusInformation *string

	// The version of the artifact associated with the document. For example, "Release
	// 12, Update 6". This value is unique across all versions of a document, and
	// cannot be changed.
	VersionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDocument{}, middleware.After)
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
	if err = addOpGetDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetDocument",
	}
}
