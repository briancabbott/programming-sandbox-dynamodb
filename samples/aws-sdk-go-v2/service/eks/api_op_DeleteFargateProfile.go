// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Fargate profile. When you delete a Fargate profile, any pods
// running on Fargate that were created with the profile are deleted. If those pods
// match another Fargate profile, then they are scheduled on Fargate with that
// profile. If they no longer match any Fargate profiles, then they are not
// scheduled on Fargate and they may remain in a pending state. Only one Fargate
// profile in a cluster can be in the DELETING status at a time. You must wait for
// a Fargate profile to finish deleting before you can delete any other profiles in
// that cluster.
func (c *Client) DeleteFargateProfile(ctx context.Context, params *DeleteFargateProfileInput, optFns ...func(*Options)) (*DeleteFargateProfileOutput, error) {
	if params == nil {
		params = &DeleteFargateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFargateProfile", params, optFns, addOperationDeleteFargateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFargateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFargateProfileInput struct {

	// The name of the Amazon EKS cluster associated with the Fargate profile to
	// delete.
	//
	// This member is required.
	ClusterName *string

	// The name of the Fargate profile to delete.
	//
	// This member is required.
	FargateProfileName *string
}

type DeleteFargateProfileOutput struct {

	// The deleted Fargate profile.
	FargateProfile *types.FargateProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFargateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFargateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFargateProfile{}, middleware.After)
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
	if err = addOpDeleteFargateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFargateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFargateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DeleteFargateProfile",
	}
}
