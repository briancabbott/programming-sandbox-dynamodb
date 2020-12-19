// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use to associate an AWS CodeCommit repository or a repostory managed by AWS
// CodeStar Connections with Amazon CodeGuru Reviewer. When you associate a
// repository, CodeGuru Reviewer reviews source code changes in the repository's
// pull requests and provides automatic recommendations. You can view
// recommendations using the CodeGuru Reviewer console. For more information, see
// Recommendations in Amazon CodeGuru Reviewer
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/recommendations.html)
// in the Amazon CodeGuru Reviewer User Guide. If you associate a CodeCommit
// repository, it must be in the same AWS Region and AWS account where its CodeGuru
// Reviewer code reviews are configured. Bitbucket and GitHub Enterprise Server
// repositories are managed by AWS CodeStar Connections to connect to CodeGuru
// Reviewer. For more information, see Connect to a repository source provider
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/reviewer-ug/step-one.html#select-repository-source-provider)
// in the Amazon CodeGuru Reviewer User Guide. You cannot use the CodeGuru Reviewer
// SDK or the AWS CLI to associate a GitHub repository with Amazon CodeGuru
// Reviewer. To associate a GitHub repository, use the console. For more
// information, see Getting started with CodeGuru Reviewer
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/getting-started-with-guru.html)
// in the CodeGuru Reviewer User Guide.
func (c *Client) AssociateRepository(ctx context.Context, params *AssociateRepositoryInput, optFns ...func(*Options)) (*AssociateRepositoryOutput, error) {
	if params == nil {
		params = &AssociateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateRepository", params, optFns, addOperationAssociateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateRepositoryInput struct {

	// The repository to associate.
	//
	// This member is required.
	Repository *types.Repository

	// Amazon CodeGuru Reviewer uses this value to prevent the accidental creation of
	// duplicate repository associations if there are failures and retries.
	ClientRequestToken *string

	// An array of key-value pairs used to tag an associated repository. A tag is a
	// custom attribute label with two parts:
	//
	// * A tag key (for example, CostCenter,
	// Environment, Project, or Secret). Tag keys are case sensitive.
	//
	// * An optional
	// field known as a tag value (for example, 111122223333, Production, or a team
	// name). Omitting the tag value is the same as using an empty string. Like tag
	// keys, tag values are case sensitive.
	Tags map[string]string
}

type AssociateRepositoryOutput struct {

	// Information about the repository association.
	RepositoryAssociation *types.RepositoryAssociation

	// An array of key-value pairs used to tag an associated repository. A tag is a
	// custom attribute label with two parts:
	//
	// * A tag key (for example, CostCenter,
	// Environment, Project, or Secret). Tag keys are case sensitive.
	//
	// * An optional
	// field known as a tag value (for example, 111122223333, Production, or a team
	// name). Omitting the tag value is the same as using an empty string. Like tag
	// keys, tag values are case sensitive.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateRepository{}, middleware.After)
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
	if err = addIdempotencyToken_opAssociateRepositoryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateRepository(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpAssociateRepository struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateRepository) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateRepositoryInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAssociateRepositoryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateRepository{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "AssociateRepository",
	}
}
