// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a broker. Note: This API is asynchronous.
func (c *Client) CreateBroker(ctx context.Context, params *CreateBrokerInput, optFns ...func(*Options)) (*CreateBrokerOutput, error) {
	if params == nil {
		params = &CreateBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBroker", params, optFns, addOperationCreateBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a broker using the specified properties.
type CreateBrokerInput struct {

	// The authentication strategy used to secure the broker.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. Enables automatic upgrades to new minor versions for brokers, as
	// Apache releases the versions. The automatic upgrades occur during the
	// maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade bool

	// Required. The name of the broker. This value must be unique in your AWS account,
	// 1-50 characters long, must contain only letters, numbers, dashes, and
	// underscores, and must not contain whitespaces, brackets, wildcard characters, or
	// special characters.
	BrokerName *string

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// The unique ID that the requester receives for the created broker. Amazon MQ
	// passes your ID with the API action. Note: We recommend using a Universally
	// Unique Identifier (UUID) for the creatorRequestId. You may omit the
	// creatorRequestId if your application doesn't require idempotency.
	CreatorRequestId *string

	// Required. The deployment mode of the broker.
	DeploymentMode types.DeploymentMode

	// Encryption options for the broker.
	EncryptionOptions *types.EncryptionOptions

	// Required. The type of broker engine. Note: Currently, Amazon MQ supports
	// ACTIVEMQ and RABBITMQ.
	EngineType types.EngineType

	// Required. The version of the broker engine. For a list of supported engine
	// versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// Required. The broker's instance type.
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// Required. Enables connections from applications outside of the VPC that hosts
	// the broker's subnets.
	PubliclyAccessible bool

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	// The broker's storage type.
	StorageType types.BrokerStorageType

	// The list of groups that define which subnets and IP ranges the broker can use
	// from different Availability Zones. A SINGLE_INSTANCE deployment requires one
	// subnet (for example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ deployment
	// (ACTIVEMQ) requires two subnets. A CLUSTER_MULTI_AZ deployment (RABBITMQ) has no
	// subnet requirements when deployed with public accessibility, deployment without
	// public accessibility requires at least one subnet.
	SubnetIds []string

	// Create tags when creating the broker.
	Tags map[string]string

	// Required. The list of broker users (persons or applications) who can access
	// queues and topics. For RabbitMQ brokers, one and only one administrative user is
	// accepted and created when a broker is first provisioned. All subsequent broker
	// users are created by making RabbitMQ API calls directly to brokers or via the
	// RabbitMQ Web Console. This value can contain only alphanumeric characters,
	// dashes, periods, underscores, and tildes (- . _ ~). This value must be 2-100
	// characters long.
	Users []types.User
}

type CreateBrokerOutput struct {

	// The Amazon Resource Name (ARN) of the broker.
	BrokerArn *string

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBroker{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateBrokerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBroker(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateBroker struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBroker) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBrokerInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBrokerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateBroker{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "CreateBroker",
	}
}