// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates the authentication flow.
func (c *Client) InitiateAuth(ctx context.Context, params *InitiateAuthInput, optFns ...func(*Options)) (*InitiateAuthOutput, error) {
	if params == nil {
		params = &InitiateAuthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitiateAuth", params, optFns, addOperationInitiateAuthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InitiateAuthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Initiates the authentication request.
type InitiateAuthInput struct {

	// The authentication flow for this call to execute. The API action will depend on
	// this value. For example:
	//
	// * REFRESH_TOKEN_AUTH will take in a valid refresh
	// token and return new tokens.
	//
	// * USER_SRP_AUTH will take in USERNAME and SRP_A
	// and return the SRP variables to be used for next challenge execution.
	//
	// *
	// USER_PASSWORD_AUTH will take in USERNAME and PASSWORD and return the next
	// challenge or tokens.
	//
	// Valid values include:
	//
	// * USER_SRP_AUTH: Authentication
	// flow for the Secure Remote Password (SRP) protocol.
	//
	// *
	// REFRESH_TOKEN_AUTH/REFRESH_TOKEN: Authentication flow for refreshing the access
	// token and ID token by supplying a valid refresh token.
	//
	// * CUSTOM_AUTH: Custom
	// authentication flow.
	//
	// * USER_PASSWORD_AUTH: Non-SRP authentication flow;
	// USERNAME and PASSWORD are passed directly. If a user migration Lambda trigger is
	// set, this flow will invoke the user migration Lambda if the USERNAME is not
	// found in the user pool.
	//
	// * ADMIN_USER_PASSWORD_AUTH: Admin-based user password
	// authentication. This replaces the ADMIN_NO_SRP_AUTH authentication flow. In this
	// flow, Cognito receives the password in the request instead of using the SRP
	// process to verify passwords.
	//
	// ADMIN_NO_SRP_AUTH is not a valid value.
	//
	// This member is required.
	AuthFlow types.AuthFlowType

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for InitiateAuth
	// calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The authentication parameters. These are inputs corresponding to the AuthFlow
	// that you are invoking. The required values depend on the value of AuthFlow:
	//
	// *
	// For USER_SRP_AUTH: USERNAME (required), SRP_A (required), SECRET_HASH (required
	// if the app client is configured with a client secret), DEVICE_KEY.
	//
	// * For
	// REFRESH_TOKEN_AUTH/REFRESH_TOKEN: REFRESH_TOKEN (required), SECRET_HASH
	// (required if the app client is configured with a client secret), DEVICE_KEY.
	//
	// *
	// For CUSTOM_AUTH: USERNAME (required), SECRET_HASH (if app client is configured
	// with client secret), DEVICE_KEY. To start the authentication flow with password
	// verification, include ChallengeName: SRP_A and SRP_A: (The SRP_A Value).
	AuthParameters map[string]string

	// A map of custom key-value pairs that you can provide as input for certain custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the InitiateAuth API
	// action, Amazon Cognito invokes the AWS Lambda functions that are specified for
	// various triggers. The ClientMetadata value is passed as input to the functions
	// for only the following triggers:
	//
	// * Pre signup
	//
	// * Pre authentication
	//
	// * User
	// migration
	//
	// When Amazon Cognito invokes the functions for these triggers, it
	// passes a JSON payload, which the function receives as input. This payload
	// contains a validationData attribute, which provides the data that you assigned
	// to the ClientMetadata parameter in your InitiateAuth request. In your function
	// code in AWS Lambda, you can process the validationData value to enhance your
	// workflow for your specific needs. When you use the InitiateAuth API action,
	// Amazon Cognito also invokes the functions for the following triggers, but it
	// does not provide the ClientMetadata value as input:
	//
	// * Post authentication
	//
	// *
	// Custom message
	//
	// * Pre token generation
	//
	// * Create auth challenge
	//
	// * Define auth
	// challenge
	//
	// * Verify auth challenge
	//
	// For more information, see Customizing User
	// Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	// * Amazon Cognito does
	// not store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.
	//
	// * Amazon Cognito does not validate the ClientMetadata
	// value.
	//
	// * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	// use it to provide sensitive information.
	ClientMetadata map[string]string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType
}

// Initiates the authentication response.
type InitiateAuthOutput struct {

	// The result of the authentication response. This is only returned if the caller
	// does not need to pass another challenge. If the caller does need to pass another
	// challenge before it gets tokens, ChallengeName, ChallengeParameters, and Session
	// are returned.
	AuthenticationResult *types.AuthenticationResultType

	// The name of the challenge which you are responding to with this call. This is
	// returned to you in the AdminInitiateAuth response if you need to pass another
	// challenge. Valid values include the following. Note that all of these challenges
	// require USERNAME and SECRET_HASH (if applicable) in the parameters.
	//
	// * SMS_MFA:
	// Next challenge is to supply an SMS_MFA_CODE, delivered via SMS.
	//
	// *
	// PASSWORD_VERIFIER: Next challenge is to supply PASSWORD_CLAIM_SIGNATURE,
	// PASSWORD_CLAIM_SECRET_BLOCK, and TIMESTAMP after the client-side SRP
	// calculations.
	//
	// * CUSTOM_CHALLENGE: This is returned if your custom
	// authentication flow determines that the user should pass another challenge
	// before tokens are issued.
	//
	// * DEVICE_SRP_AUTH: If device tracking was enabled on
	// your user pool and the previous challenges were passed, this challenge is
	// returned so that Amazon Cognito can start tracking this device.
	//
	// *
	// DEVICE_PASSWORD_VERIFIER: Similar to PASSWORD_VERIFIER, but for devices only.
	//
	// *
	// NEW_PASSWORD_REQUIRED: For users which are required to change their passwords
	// after successful first login. This challenge should be passed with NEW_PASSWORD
	// and any other required attributes.
	ChallengeName types.ChallengeNameType

	// The challenge parameters. These are returned to you in the InitiateAuth response
	// if you need to pass another challenge. The responses in this parameter should be
	// used to compute inputs to the next call (RespondToAuthChallenge). All challenges
	// require USERNAME and SECRET_HASH (if applicable).
	ChallengeParameters map[string]string

	// The session which should be passed both ways in challenge-response calls to the
	// service. If the caller needs to go through another challenge, they return a
	// session with other challenge parameters. This session should be passed as it is
	// to the next RespondToAuthChallenge API call.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInitiateAuthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInitiateAuth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInitiateAuth{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpInitiateAuthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateAuth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInitiateAuth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InitiateAuth",
	}
}