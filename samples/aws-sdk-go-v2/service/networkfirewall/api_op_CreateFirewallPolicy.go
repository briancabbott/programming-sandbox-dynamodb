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

// Creates the firewall policy for the firewall according to the specifications. An
// AWS Network Firewall firewall policy defines the behavior of a firewall, in a
// collection of stateless and stateful rule groups and other settings. You can use
// one firewall policy for multiple firewalls.
func (c *Client) CreateFirewallPolicy(ctx context.Context, params *CreateFirewallPolicyInput, optFns ...func(*Options)) (*CreateFirewallPolicyOutput, error) {
	if params == nil {
		params = &CreateFirewallPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFirewallPolicy", params, optFns, addOperationCreateFirewallPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFirewallPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFirewallPolicyInput struct {

	// The rule groups and policy actions to use in the firewall policy.
	//
	// This member is required.
	FirewallPolicy *types.FirewallPolicy

	// The descriptive name of the firewall policy. You can't change the name of a
	// firewall policy after you create it.
	//
	// This member is required.
	FirewallPolicyName *string

	// A description of the firewall policy.
	Description *string

	// Indicates whether you want Network Firewall to just check the validity of the
	// request, rather than run the request. If set to TRUE, Network Firewall checks
	// whether the request can run successfully, but doesn't actually make the
	// requested changes. The call returns the value that the request would return if
	// you ran it with dry run set to FALSE, but doesn't make additions or changes to
	// your resources. This option allows you to make sure that you have the required
	// permissions to run the request and that your request parameters are valid. If
	// set to FALSE, Network Firewall makes the requested changes to your resources.
	DryRun bool

	// The key:value pairs to associate with the resource.
	Tags []types.Tag
}

type CreateFirewallPolicyOutput struct {

	// The high-level properties of a firewall policy. This, along with the
	// FirewallPolicy, define the policy. You can retrieve all objects for a firewall
	// policy by calling DescribeFirewallPolicy.
	//
	// This member is required.
	FirewallPolicyResponse *types.FirewallPolicyResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the firewall policy. The token marks the state of the
	// policy resource at the time of the request. To make changes to the policy, you
	// provide the token in your request. Network Firewall uses the token to ensure
	// that the policy hasn't changed since you last retrieved it. If it has changed,
	// the operation fails with an InvalidTokenException. If this happens, retrieve the
	// firewall policy again to get a current copy of it with current token. Reapply
	// your changes as needed, then try the operation again using the new token.
	//
	// This member is required.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFirewallPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateFirewallPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateFirewallPolicy{}, middleware.After)
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
	if err = addOpCreateFirewallPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFirewallPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFirewallPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "CreateFirewallPolicy",
	}
}
