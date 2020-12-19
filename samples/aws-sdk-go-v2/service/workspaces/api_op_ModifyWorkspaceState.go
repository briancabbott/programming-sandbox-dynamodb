// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the state of the specified WorkSpace. To maintain a WorkSpace without being
// interrupted, set the WorkSpace state to ADMIN_MAINTENANCE. WorkSpaces in this
// state do not respond to requests to reboot, stop, start, rebuild, or restore. An
// AutoStop WorkSpace in this state is not stopped. Users cannot log into a
// WorkSpace in the ADMIN_MAINTENANCE state.
func (c *Client) ModifyWorkspaceState(ctx context.Context, params *ModifyWorkspaceStateInput, optFns ...func(*Options)) (*ModifyWorkspaceStateOutput, error) {
	if params == nil {
		params = &ModifyWorkspaceStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyWorkspaceState", params, optFns, addOperationModifyWorkspaceStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyWorkspaceStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspaceStateInput struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	// The WorkSpace state.
	//
	// This member is required.
	WorkspaceState types.TargetWorkspaceState
}

type ModifyWorkspaceStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyWorkspaceStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyWorkspaceState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyWorkspaceState{}, middleware.After)
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
	if err = addOpModifyWorkspaceStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyWorkspaceState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyWorkspaceState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceState",
	}
}
