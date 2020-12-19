// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a member account from its parent organization. This version of the
// operation is performed by the account that wants to leave. To remove a member
// account as a user in the management account, use RemoveAccountFromOrganization
// instead. This operation can be called only from a member account in the
// organization.
//
// * The management account in an organization with all features
// enabled can set service control policies (SCPs) that can restrict what
// administrators of member accounts can do. This includes preventing them from
// successfully calling LeaveOrganization and leaving the organization.
//
// * You can
// leave an organization as a member account only if the account is configured with
// the information required to operate as a standalone account. When you create an
// account in an organization using the AWS Organizations console, API, or CLI
// commands, the information required of standalone accounts is not automatically
// collected. For each account that you want to make standalone, you must perform
// the following steps. If any of the steps are already completed for this account,
// that step doesn't appear.
//
// * Choose a support plan
//
// * Provide and verify the
// required contact information
//
// * Provide a current payment method
//
// AWS uses the
// payment method to charge for any billable (not free tier) AWS activity that
// occurs while the account isn't attached to an organization. Follow the steps at
// To leave an organization when all required account information has not yet been
// provided
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
// in the AWS Organizations User Guide.
//
// * You can leave an organization only after
// you enable IAM user access to billing in your account. For more information, see
// Activating Access to the Billing and Cost Management Console
// (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
// in the AWS Billing and Cost Management User Guide.
//
// * After the account leaves
// the organization, all tags that were attached to the account object in the
// organization are deleted. AWS accounts outside of an organization do not support
// tags.
func (c *Client) LeaveOrganization(ctx context.Context, params *LeaveOrganizationInput, optFns ...func(*Options)) (*LeaveOrganizationOutput, error) {
	if params == nil {
		params = &LeaveOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LeaveOrganization", params, optFns, addOperationLeaveOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LeaveOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LeaveOrganizationInput struct {
}

type LeaveOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationLeaveOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpLeaveOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpLeaveOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLeaveOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opLeaveOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "LeaveOrganization",
	}
}
