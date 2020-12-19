// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified FirewallPolicy.
func (c *Client) DeleteFirewallPolicy(ctx context.Context, params *DeleteFirewallPolicyInput, optFns ...func(*Options)) (*DeleteFirewallPolicyOutput, error) {
	if params == nil {
		params = &DeleteFirewallPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFirewallPolicy", params, optFns, addOperationDeleteFirewallPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFirewallPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFirewallPolicyInput struct {

	// The Amazon Resource Name (ARN) of the firewall policy. You must specify the ARN
	// or the name, and you can specify both.
	FirewallPolicyArn *string

	// The descriptive name of the firewall policy. You can't change the name of a
	// firewall policy after you create it. You must specify the ARN or the name, and
	// you can specify both.
	FirewallPolicyName *string
}

type DeleteFirewallPolicyOutput struct {

	// The object containing the definition of the FirewallPolicyResponse that you
	// asked to delete.
	//
	// This member is required.
	FirewallPolicyResponse *types.FirewallPolicyResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFirewallPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFirewallPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFirewallPolicy{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFirewallPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFirewallPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "DeleteFirewallPolicy",
	}
}
