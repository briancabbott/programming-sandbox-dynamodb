// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Uploads a part by copying data from an existing object as data source. You
// specify the data source by adding the request header x-amz-copy-source in your
// request and a byte range by adding the request header x-amz-copy-source-range in
// your request. The minimum allowable part size for a multipart upload is 5 MB.
// For more information about multipart upload limits, go to Quick Facts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/qfacts.html) in the Amazon
// Simple Storage Service Developer Guide. Instead of using an existing object as
// part data, you might use the UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html) operation
// and provide data in your request. You must initiate a multipart upload before
// you can upload any part. In response to your initiate request. Amazon S3 returns
// a unique identifier, the upload ID, that you must include in your upload part
// request. For more information about using the UploadPartCopy operation, see the
// following:
//
// * For conceptual information about multipart uploads, see Uploading
// Objects Using Multipart Upload
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html) in the
// Amazon Simple Storage Service Developer Guide.
//
// * For information about
// permissions required to use the multipart upload API, see Multipart Upload API
// and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html) in the
// Amazon Simple Storage Service Developer Guide.
//
// * For information about copying
// objects using a single atomic operation vs. the multipart upload, see Operations
// on Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html) in the
// Amazon Simple Storage Service Developer Guide.
//
// * For information about using
// server-side encryption with customer-provided encryption keys with the
// UploadPartCopy operation, see CopyObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html) and
// UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html).
//
// Note the
// following additional considerations about the request headers
// x-amz-copy-source-if-match, x-amz-copy-source-if-none-match,
// x-amz-copy-source-if-unmodified-since, and
// x-amz-copy-source-if-modified-since:
//
// * Consideration 1 - If both of the
// x-amz-copy-source-if-match and x-amz-copy-source-if-unmodified-since headers are
// present in the request as follows: x-amz-copy-source-if-match condition
// evaluates to true, and; x-amz-copy-source-if-unmodified-since condition
// evaluates to false; Amazon S3 returns 200 OK and copies the data.
//
// *
// Consideration 2 - If both of the x-amz-copy-source-if-none-match and
// x-amz-copy-source-if-modified-since headers are present in the request as
// follows: x-amz-copy-source-if-none-match condition evaluates to false, and;
// x-amz-copy-source-if-modified-since condition evaluates to true; Amazon S3
// returns 412 Precondition Failed response code.
//
// Versioning If your bucket has
// versioning enabled, you could have multiple versions of the same object. By
// default, x-amz-copy-source identifies the current version of the object to copy.
// If the current version is a delete marker and you don't specify a versionId in
// the x-amz-copy-source, Amazon S3 returns a 404 error, because the object does
// not exist. If you specify versionId in the x-amz-copy-source and the versionId
// is a delete marker, Amazon S3 returns an HTTP 400 error, because you are not
// allowed to specify a delete marker as a version for the x-amz-copy-source. You
// can optionally specify a specific version of the source object to copy by adding
// the versionId subresource as shown in the following example: x-amz-copy-source:
// /bucket/object?versionId=version id Special Errors
//
// * Code: NoSuchUpload
//
// *
// Cause: The specified multipart upload does not exist. The upload ID might be
// invalid, or the multipart upload might have been aborted or completed.
//
// * HTTP
// Status Code: 404 Not Found
//
// * Code: InvalidRequest
//
// * Cause: The specified copy
// source is not supported as a byte-range copy source.
//
// * HTTP Status Code: 400
// Bad Request
//
// Related Resources
//
// * CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//
// *
// UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//
// *
// CompleteMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//
// *
// AbortMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
//
// *
// ListParts
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//
// *
// ListMultipartUploads
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html)
func (c *Client) UploadPartCopy(ctx context.Context, params *UploadPartCopyInput, optFns ...func(*Options)) (*UploadPartCopyOutput, error) {
	if params == nil {
		params = &UploadPartCopyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadPartCopy", params, optFns, addOperationUploadPartCopyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadPartCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadPartCopyInput struct {

	// The bucket name. When using this API with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation with an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide. When using this API with
	// Amazon S3 on Outposts, you must direct requests to the S3 on Outposts hostname.
	// The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this operation using S3 on Outposts through the AWS SDKs, you provide the
	// Outposts bucket ARN in place of the bucket name. For more information about S3
	// on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Specifies the source object for the copy operation. You specify the value in one
	// of two formats, depending on whether you want to access the source object
	// through an access point
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points.html):
	//
	// * For
	// objects not accessed through an access point, specify the name of the source
	// bucket and key of the source object, separated by a slash (/). For example, to
	// copy the object reports/january.pdf from the bucket awsexamplebucket, use
	// awsexamplebucket/reports/january.pdf. The value must be URL encoded.
	//
	// * For
	// objects accessed through access points, specify the Amazon Resource Name (ARN)
	// of the object as accessed through the access point, in the format
	// arn:aws:s3:::accesspoint//object/. For example, to copy the object
	// reports/january.pdf through access point my-access-point owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/reports/january.pdf.
	// The value must be URL encoded. Amazon S3 supports copy operations using access
	// points only when the source and destination buckets are in the same AWS Region.
	// Alternatively, for objects accessed through Amazon S3 on Outposts, specify the
	// ARN of the object as accessed in the format
	// arn:aws:s3-outposts:::outpost//object/. For example, to copy the object
	// reports/january.pdf through outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/object/reports/january.pdf.
	// The value must be URL encoded.
	//
	// To copy a specific version of an object, append
	// ?versionId= to the value (for example,
	// awsexamplebucket/reports/january.pdf?versionId=QUpfdndhfd8438MNFDN93jdnJFkdmqnh893).
	// If you don't specify a version ID, Amazon S3 copies the latest version of the
	// source object.
	//
	// This member is required.
	CopySource *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Part number of part being copied. This is a positive integer between 1 and
	// 10,000.
	//
	// This member is required.
	PartNumber int32

	// Upload ID identifying the multipart upload whose part is being copied.
	//
	// This member is required.
	UploadId *string

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time

	// The range of bytes to copy from the source object. The range value must use the
	// form bytes=first-last, where the first and last are the zero-based byte offsets
	// to copy. For example, bytes=0-9 indicates that you want to copy the first 10
	// bytes of the source. You can copy a range only if the source object is greater
	// than 5 MB.
	CopySourceRange *string

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256).
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one that
	// was used when the source object was created.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string

	// The account id of the expected destination bucket owner. If the destination
	// bucket is owned by a different account, the request will fail with an HTTP 403
	// (Access Denied) error.
	ExpectedBucketOwner *string

	// The account id of the expected source bucket owner. If the source bucket is
	// owned by a different account, the request will fail with an HTTP 403 (Access
	// Denied) error.
	ExpectedSourceBucketOwner *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header. This must be the same
	// encryption key specified in the initiate multipart upload request.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string
}

type UploadPartCopyOutput struct {

	// Container for all response elements.
	CopyPartResult *types.CopyPartResult

	// The version of the source object that was copied, if you have enabled versioning
	// on the source bucket.
	CopySourceVersionId *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadPartCopyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUploadPartCopy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUploadPartCopy{}, middleware.After)
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
	if err = addOpUploadPartCopyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadPartCopy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUploadPartCopyUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = s3cust.HandleResponseErrorWith200Status(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUploadPartCopy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UploadPartCopy",
	}
}

// getUploadPartCopyBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getUploadPartCopyBucketMember(input interface{}) (*string, bool) {
	in := input.(*UploadPartCopyInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addUploadPartCopyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getUploadPartCopyBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}