// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Modifies the specified replication task. You can't modify the task endpoints.
// The task must be stopped before you can modify it. For more information about
// AWS DMS tasks, see Working with Migration Tasks
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html) in the AWS
// Database Migration Service User Guide.
func (c *Client) ModifyReplicationTask(ctx context.Context, params *ModifyReplicationTaskInput, optFns ...func(*Options)) (*ModifyReplicationTaskOutput, error) {
	if params == nil {
		params = &ModifyReplicationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationTask", params, optFns, addOperationModifyReplicationTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyReplicationTaskInput struct {

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// This member is required.
	ReplicationTaskArn *string

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error. The value can be in date,
	// checkpoint, or LSN/SCN format. Date Example: --cdc-start-position
	// “2018-03-08T12:12:12” Checkpoint Example: --cdc-start-position
	// "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373” When you use
	// this task setting with a source PostgreSQL database, a logical replication slot
	// should already be created and associated with the source endpoint. You can
	// verify this by setting the slotName extra connection attribute to the name of
	// this logical replication slot. For more information, see Extra Connection
	// Attributes When Using PostgreSQL as a Source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib).
	CdcStartPosition *string

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation to
	// start. Specifying both values results in an error. Timestamp Example:
	// --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time

	// Indicates when you want a change data capture (CDC) operation to stop. The value
	// can be either server time or commit time. Server time example:
	// --cdc-stop-position “server_time:2018-02-09T12:12:12” Commit time example:
	// --cdc-stop-position “commit_time: 2018-02-09T12:12:12 “
	CdcStopPosition *string

	// The migration type. Valid values: full-load | cdc | full-load-and-cdc
	MigrationType types.MigrationTypeValue

	// The replication task identifier. Constraints:
	//
	// * Must contain 1-255 alphanumeric
	// characters or hyphens.
	//
	// * First character must be a letter.
	//
	// * Cannot end with a
	// hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier *string

	// JSON file that contains settings for the task, such as task metadata settings.
	ReplicationTaskSettings *string

	// When using the AWS CLI or boto3, provide the path of the JSON file that contains
	// the table mappings. Precede the path with file://. When working with the DMS
	// API, provide the JSON as the parameter value, for example: --table-mappings
	// file://mappingfile.json
	TableMappings *string

	// Supplemental information that the task requires to migrate the data for certain
	// source and target endpoints. For more information, see Specifying Supplemental
	// Data for Task Settings
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html) in
	// the AWS Database Migration Service User Guide.
	TaskData *string
}

//
type ModifyReplicationTaskOutput struct {

	// The replication task that was modified.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyReplicationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReplicationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReplicationTask{}, middleware.After)
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
	if err = addOpModifyReplicationTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReplicationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyReplicationTask",
	}
}
