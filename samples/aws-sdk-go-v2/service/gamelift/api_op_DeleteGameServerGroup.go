// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Terminates a game server group and permanently deletes the game server
// group record. You have several options for how these resources are impacted when
// deleting the game server group. Depending on the type of delete operation
// selected, this operation might affect these resources:
//
// * The game server
// group
//
// * The corresponding Auto Scaling group
//
// * All game servers that are
// currently running in the group
//
// To delete a game server group, identify the game
// server group to delete and specify the type of delete operation to initiate.
// Game server groups can only be deleted if they are in ACTIVE or ERROR status. If
// the delete request is successful, a series of operations are kicked off. The
// game server group status is changed to DELETE_SCHEDULED, which prevents new game
// servers from being registered and stops automatic scaling activity. Once all
// game servers in the game server group are deregistered, GameLift FleetIQ can
// begin deleting resources. If any of the delete operations fail, the game server
// group is placed in ERROR status. GameLift FleetIQ emits delete events to Amazon
// CloudWatch. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
// * CreateGameServerGroup
//
// * ListGameServerGroups
//
// *
// DescribeGameServerGroup
//
// * UpdateGameServerGroup
//
// * DeleteGameServerGroup
//
// *
// ResumeGameServerGroup
//
// * SuspendGameServerGroup
//
// * DescribeGameServerInstances
func (c *Client) DeleteGameServerGroup(ctx context.Context, params *DeleteGameServerGroupInput, optFns ...func(*Options)) (*DeleteGameServerGroupOutput, error) {
	if params == nil {
		params = &DeleteGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGameServerGroup", params, optFns, addOperationDeleteGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGameServerGroupInput struct {

	// A unique identifier for the game server group. Use either the GameServerGroup
	// name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The type of delete to perform. Options include the following:
	//
	// * SAFE_DELETE –
	// Terminates the game server group and EC2 Auto Scaling group only when it has no
	// game servers that are in UTILIZED status.
	//
	// * FORCE_DELETE – Terminates the game
	// server group, including all active game servers regardless of their utilization
	// status, and the EC2 Auto Scaling group.
	//
	// * RETAIN – Does a safe delete of the
	// game server group but retains the EC2 Auto Scaling group as is.
	DeleteOption types.GameServerGroupDeleteOption
}

type DeleteGameServerGroupOutput struct {

	// An object that describes the deleted game server group resource, with status
	// updated to DELETE_SCHEDULED.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGameServerGroup{}, middleware.After)
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
	if err = addOpDeleteGameServerGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGameServerGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGameServerGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteGameServerGroup",
	}
}
