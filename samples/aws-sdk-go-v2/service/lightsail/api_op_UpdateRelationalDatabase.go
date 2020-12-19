// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the update of one or more attributes of a database in Amazon Lightsail.
// Updates are applied immediately, or in cases where the updates could result in
// an outage, are applied during the database's predefined maintenance window. The
// update relational database operation supports tag-based access control via
// resource tags applied to the resource identified by relationalDatabaseName. For
// more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) UpdateRelationalDatabase(ctx context.Context, params *UpdateRelationalDatabaseInput, optFns ...func(*Options)) (*UpdateRelationalDatabaseOutput, error) {
	if params == nil {
		params = &UpdateRelationalDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRelationalDatabase", params, optFns, addOperationUpdateRelationalDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRelationalDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRelationalDatabaseInput struct {

	// The name of your database to update.
	//
	// This member is required.
	RelationalDatabaseName *string

	// When true, applies changes immediately. When false, applies changes during the
	// preferred maintenance window. Some changes may cause an outage. Default: false
	ApplyImmediately *bool

	// Indicates the certificate that needs to be associated with the database.
	CaCertificateIdentifier *string

	// When true, disables automated backup retention for your database. Disabling
	// backup retention deletes all automated database backups. Before disabling this,
	// you may want to create a snapshot of your database using the create relational
	// database snapshot operation. Updates are applied during the next maintenance
	// window because this can result in an outage.
	DisableBackupRetention *bool

	// When true, enables automated backup retention for your database. Updates are
	// applied during the next maintenance window because this can result in an outage.
	EnableBackupRetention *bool

	// The password for the master user of your database. The password can include any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain 8
	// to 41 characters.
	MasterUserPassword *string

	// The daily time range during which automated backups are created for your
	// database if automated backups are enabled. Constraints:
	//
	// * Must be in the
	// hh24:mi-hh24:mi format. Example: 16:00-16:30
	//
	// * Specified in Coordinated
	// Universal Time (UTC).
	//
	// * Must not conflict with the preferred maintenance
	// window.
	//
	// * Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur on your
	// database. The default is a 30-minute window selected at random from an 8-hour
	// block of time for each AWS Region, occurring on a random day of the week.
	// Constraints:
	//
	// * Must be in the ddd:hh24:mi-ddd:hh24:mi format.
	//
	// * Valid days:
	// Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// * Must be at least 30 minutes.
	//
	// * Specified
	// in Coordinated Universal Time (UTC).
	//
	// * Example: Tue:17:00-Tue:17:30
	PreferredMaintenanceWindow *string

	// Specifies the accessibility options for your database. A value of true specifies
	// a database that is available to resources outside of your Lightsail account. A
	// value of false specifies a database that is available only to your Lightsail
	// resources in the same region as your database.
	PubliclyAccessible *bool

	// When true, the master user password is changed to a new strong password
	// generated by Lightsail. Use the get relational database master user password
	// operation to get the new password.
	RotateMasterUserPassword *bool
}

type UpdateRelationalDatabaseOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRelationalDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRelationalDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRelationalDatabase{}, middleware.After)
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
	if err = addOpUpdateRelationalDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRelationalDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRelationalDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateRelationalDatabase",
	}
}
