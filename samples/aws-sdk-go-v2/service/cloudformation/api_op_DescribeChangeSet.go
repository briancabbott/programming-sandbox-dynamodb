// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the inputs for the change set and a list of changes that AWS
// CloudFormation will make if you execute the change set. For more information,
// see Updating Stacks Using Change Sets
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html)
// in the AWS CloudFormation User Guide.
func (c *Client) DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) {
	if params == nil {
		params = &DescribeChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChangeSet", params, optFns, addOperationDescribeChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeChangeSet action.
type DescribeChangeSetInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// describe.
	//
	// This member is required.
	ChangeSetName *string

	// A string (provided by the DescribeChangeSet response output) that identifies the
	// next page of information that you want to retrieve.
	NextToken *string

	// If you specified the name of a change set, specify the stack name or ID (ARN) of
	// the change set you want to describe.
	StackName *string
}

// The output for the DescribeChangeSet action.
type DescribeChangeSetOutput struct {

	// If you execute the change set, the list of capabilities that were explicitly
	// acknowledged when the change set was created.
	Capabilities []types.Capability

	// The ARN of the change set.
	ChangeSetId *string

	// The name of the change set.
	ChangeSetName *string

	// A list of Change structures that describes the resources AWS CloudFormation
	// changes if you execute the change set.
	Changes []types.Change

	// The start time when the change set was created, in UTC.
	CreationTime *time.Time

	// Information about the change set.
	Description *string

	// If the change set execution status is AVAILABLE, you can execute the change set.
	// If you can’t execute the change set, the status indicates why. For example, a
	// change set might be in an UNAVAILABLE state because AWS CloudFormation is still
	// creating it or in an OBSOLETE state because the stack was already updated.
	ExecutionStatus types.ExecutionStatus

	// Verifies if IncludeNestedStacks is set to True.
	IncludeNestedStacks *bool

	// If the output exceeds 1 MB, a string that identifies the next page of changes.
	// If there is no additional page, this value is null.
	NextToken *string

	// The ARNs of the Amazon Simple Notification Service (Amazon SNS) topics that will
	// be associated with the stack if you execute the change set.
	NotificationARNs []string

	// A list of Parameter structures that describes the input parameters and their
	// values used to create the change set. For more information, see the Parameter
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []types.Parameter

	// Specifies the change set ID of the parent change set in the current nested
	// change set hierarchy.
	ParentChangeSetId *string

	// The rollback triggers for AWS CloudFormation to monitor during stack creation
	// and updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Specifies the change set ID of the root change set in the current nested change
	// set hierarchy.
	RootChangeSetId *string

	// The ARN of the stack that is associated with the change set.
	StackId *string

	// The name of the stack that is associated with the change set.
	StackName *string

	// The current status of the change set, such as CREATE_IN_PROGRESS,
	// CREATE_COMPLETE, or FAILED.
	Status types.ChangeSetStatus

	// A description of the change set's status. For example, if your attempt to create
	// a change set failed, AWS CloudFormation shows the error message.
	StatusReason *string

	// If you execute the change set, the tags that will be associated with the stack.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeChangeSet{}, middleware.After)
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
	if err = addOpDescribeChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeChangeSet",
	}
}
