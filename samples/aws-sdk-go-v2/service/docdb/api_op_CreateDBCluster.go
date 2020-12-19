// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon DocumentDB cluster.
func (c *Client) CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBCluster", params, optFns, addOperationCreateDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBCluster.
type CreateDBClusterInput struct {

	// The cluster identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	// * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	// * The
	// first character must be a letter.
	//
	// * Cannot end with a hyphen or contain two
	// consecutive hyphens.
	//
	// Example: my-cluster
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the database engine to be used for this cluster. Valid values: docdb
	//
	// This member is required.
	Engine *string

	// The password for the master database user. This password can contain any
	// printable ASCII character except forward slash (/), double quote ("), or the
	// "at" symbol (@). Constraints: Must contain from 8 to 100 characters.
	//
	// This member is required.
	MasterUserPassword *string

	// The name of the master user for the cluster. Constraints:
	//
	// * Must be from 1 to
	// 63 letters or numbers.
	//
	// * The first character must be a letter.
	//
	// * Cannot be a
	// reserved word for the chosen database engine.
	//
	// This member is required.
	MasterUsername *string

	// A list of Amazon EC2 Availability Zones that instances in the cluster can be
	// created in.
	AvailabilityZones []string

	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//
	// * Must be a value from 1 to 35.
	BackupRetentionPeriod *int32

	// The name of the cluster parameter group to associate with this cluster.
	DBClusterParameterGroupName *string

	// A subnet group to associate with this cluster. Constraints: Must match the name
	// of an existing DBSubnetGroup. Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string

	// Specifies whether this cluster can be deleted. If DeletionProtection is enabled,
	// the cluster cannot be deleted unless it is modified and DeletionProtection is
	// disabled. DeletionProtection protects clusters from being accidentally deleted.
	DeletionProtection *bool

	// A list of log types that need to be enabled for exporting to Amazon CloudWatch
	// Logs. You can enable audit logs or profiler logs. For more information, see
	// Auditing Amazon DocumentDB Events
	// (https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html)
	// and  Profiling Amazon DocumentDB Operations
	// (https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html).
	EnableCloudwatchLogsExports []string

	// The version number of the database engine to use. The --engine-version will
	// default to the latest major engine version. For production workloads, we
	// recommend explicitly declaring this parameter with the intended major engine
	// version.
	EngineVersion *string

	// The AWS KMS key identifier for an encrypted cluster. The AWS KMS key identifier
	// is the Amazon Resource Name (ARN) for the AWS KMS encryption key. If you are
	// creating a cluster using the same AWS account that owns the AWS KMS encryption
	// key that is used to encrypt the new cluster, you can use the AWS KMS key alias
	// instead of the ARN for the AWS KMS encryption key. If an encryption key is not
	// specified in KmsKeyId:
	//
	// * If the StorageEncrypted parameter is true, Amazon
	// DocumentDB uses your default encryption key.
	//
	// AWS KMS creates the default
	// encryption key for your AWS account. Your AWS account has a different default
	// encryption key for each AWS Region.
	KmsKeyId *string

	// The port number on which the instances in the cluster accept connections.
	Port *int32

	// Not currently supported.
	PreSignedUrl *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. Constraints:
	//
	// * Must be in the format hh24:mi-hh24:mi.
	//
	// * Must be in
	// Universal Coordinated Time (UTC).
	//
	// * Must not conflict with the preferred
	// maintenance window.
	//
	// * Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. Valid days: Mon, Tue, Wed, Thu,
	// Fri, Sat, Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// Specifies whether the cluster is encrypted.
	StorageEncrypted *bool

	// The tags to be assigned to the cluster.
	Tags []types.Tag

	// A list of EC2 VPC security groups to associate with this cluster.
	VpcSecurityGroupIds []string
}

type CreateDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBCluster{}, middleware.After)
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
	if err = addOpCreateDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBCluster",
	}
}
