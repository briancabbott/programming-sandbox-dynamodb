// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB cluster snapshot. To share a manual DB cluster snapshot with other AWS
// accounts, specify restore as the AttributeName and use the ValuesToAdd parameter
// to add a list of IDs of the AWS accounts that are authorized to restore the
// manual DB cluster snapshot. Use the value all to make the manual DB cluster
// snapshot public, which means that it can be copied or restored by all AWS
// accounts. Do not add the all value for any manual DB cluster snapshots that
// contain private information that you don't want available to all AWS accounts.
// If a manual DB cluster snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized AWS account IDs for the ValuesToAdd parameter.
// You can't use all as a value for that parameter in this case. To view which AWS
// accounts have access to copy or restore a manual DB cluster snapshot, or whether
// a manual DB cluster snapshot public or private, use the
// DescribeDBClusterSnapshotAttributes API action.
func (c *Client) ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBClusterSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterSnapshotAttribute", params, optFns, addOperationModifyDBClusterSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterSnapshotAttributeInput struct {

	// The name of the DB cluster snapshot attribute to modify. To manage authorization
	// for other AWS accounts to copy or restore a manual DB cluster snapshot, set this
	// value to restore.
	//
	// This member is required.
	AttributeName *string

	// The identifier for the DB cluster snapshot to modify the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	// A list of DB cluster snapshot attributes to add to the attribute specified by
	// AttributeName. To authorize other AWS accounts to copy or restore a manual DB
	// cluster snapshot, set this list to include one or more AWS account IDs, or all
	// to make the manual DB cluster snapshot restorable by any AWS account. Do not add
	// the all value for any manual DB cluster snapshots that contain private
	// information that you don't want available to all AWS accounts.
	ValuesToAdd []string

	// A list of DB cluster snapshot attributes to remove from the attribute specified
	// by AttributeName. To remove authorization for other AWS accounts to copy or
	// restore a manual DB cluster snapshot, set this list to include one or more AWS
	// account identifiers, or all to remove authorization for any AWS account to copy
	// or restore the DB cluster snapshot. If you specify all, an AWS account whose
	// account ID is explicitly added to the restore attribute can still copy or
	// restore a manual DB cluster snapshot.
	ValuesToRemove []string
}

type ModifyDBClusterSnapshotAttributeOutput struct {

	// Contains the results of a successful call to the
	// DescribeDBClusterSnapshotAttributes API action. Manual DB cluster snapshot
	// attributes are used to authorize other AWS accounts to copy or restore a manual
	// DB cluster snapshot. For more information, see the
	// ModifyDBClusterSnapshotAttribute API action.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
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
	if err = addOpModifyDBClusterSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterSnapshotAttribute",
	}
}
