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

// Retrieves a list that describes the connection aliases used for cross-Region
// redirection. For more information, see  Cross-Region Redirection for Amazon
// WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
func (c *Client) DescribeConnectionAliases(ctx context.Context, params *DescribeConnectionAliasesInput, optFns ...func(*Options)) (*DescribeConnectionAliasesOutput, error) {
	if params == nil {
		params = &DescribeConnectionAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectionAliases", params, optFns, addOperationDescribeConnectionAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectionAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectionAliasesInput struct {

	// The identifiers of the connection aliases to describe.
	AliasIds []string

	// The maximum number of connection aliases to return.
	Limit *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// The identifier of the directory associated with the connection alias.
	ResourceId *string
}

type DescribeConnectionAliasesOutput struct {

	// Information about the specified connection aliases.
	ConnectionAliases []types.ConnectionAlias

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConnectionAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnectionAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnectionAliases{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectionAliases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectionAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeConnectionAliases",
	}
}
