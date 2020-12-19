// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tags that are attached to the specified user. The returned list of
// tags is sorted by tag key. For more information about tagging, see Tagging IAM
// Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in
// the IAM User Guide.
func (c *Client) ListUserTags(ctx context.Context, params *ListUserTagsInput, optFns ...func(*Options)) (*ListUserTagsOutput, error) {
	if params == nil {
		params = &ListUserTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserTags", params, optFns, addOperationListUserTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserTagsInput struct {

	// The name of the IAM user whose tags you want to see. This parameter accepts
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that consist of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: =,.@-
	//
	// This member is required.
	UserName *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// (Optional) Use this only when paginating results to indicate the maximum number
	// of items that you want in the response. If additional items exist beyond the
	// maximum that you specify, the IsTruncated response element is true. If you do
	// not include this parameter, it defaults to 100. Note that IAM might return fewer
	// results, even when more results are available. In that case, the IsTruncated
	// response element returns true, and Marker contains a value to include in the
	// subsequent call that tells the service where to continue from.
	MaxItems *int32
}

type ListUserTagsOutput struct {

	// The list of tags that are currently attached to the user. Each tag consists of a
	// key name and an associated value. If no tags are attached to the specified user,
	// the response contains an empty list.
	//
	// This member is required.
	Tags []types.Tag

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can use the Marker request parameter to make a subsequent
	// pagination request that retrieves more items. Note that IAM might return fewer
	// than the MaxItems number of results even when more results are available. Check
	// IsTruncated after every call to ensure that you receive all of your results.
	IsTruncated bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUserTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListUserTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListUserTags{}, middleware.After)
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
	if err = addOpListUserTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListUserTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListUserTags",
	}
}
