// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpBatchGetNamedQuery struct {
}

func (*validateOpBatchGetNamedQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetNamedQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchGetQueryExecution struct {
}

func (*validateOpBatchGetQueryExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetQueryExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataCatalog struct {
}

func (*validateOpCreateDataCatalog) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataCatalog) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataCatalogInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataCatalogInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNamedQuery struct {
}

func (*validateOpCreateNamedQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNamedQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWorkGroup struct {
}

func (*validateOpCreateWorkGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWorkGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWorkGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWorkGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataCatalog struct {
}

func (*validateOpDeleteDataCatalog) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataCatalog) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataCatalogInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataCatalogInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNamedQuery struct {
}

func (*validateOpDeleteNamedQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNamedQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWorkGroup struct {
}

func (*validateOpDeleteWorkGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWorkGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWorkGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWorkGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDatabase struct {
}

func (*validateOpGetDatabase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDatabase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDatabaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDatabaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDataCatalog struct {
}

func (*validateOpGetDataCatalog) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDataCatalog) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDataCatalogInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDataCatalogInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNamedQuery struct {
}

func (*validateOpGetNamedQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNamedQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueryExecution struct {
}

func (*validateOpGetQueryExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueryExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueryResults struct {
}

func (*validateOpGetQueryResults) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueryResults) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueryResultsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueryResultsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTableMetadata struct {
}

func (*validateOpGetTableMetadata) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTableMetadata) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTableMetadataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTableMetadataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorkGroup struct {
}

func (*validateOpGetWorkGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorkGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDatabases struct {
}

func (*validateOpListDatabases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDatabases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDatabasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDatabasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTableMetadata struct {
}

func (*validateOpListTableMetadata) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTableMetadata) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTableMetadataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTableMetadataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartQueryExecution struct {
}

func (*validateOpStartQueryExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartQueryExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopQueryExecution struct {
}

func (*validateOpStopQueryExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopQueryExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDataCatalog struct {
}

func (*validateOpUpdateDataCatalog) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataCatalog) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataCatalogInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataCatalogInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkGroup struct {
}

func (*validateOpUpdateWorkGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetNamedQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetNamedQuery{}, middleware.After)
}

func addOpBatchGetQueryExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetQueryExecution{}, middleware.After)
}

func addOpCreateDataCatalogValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataCatalog{}, middleware.After)
}

func addOpCreateNamedQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNamedQuery{}, middleware.After)
}

func addOpCreateWorkGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWorkGroup{}, middleware.After)
}

func addOpDeleteDataCatalogValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataCatalog{}, middleware.After)
}

func addOpDeleteNamedQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNamedQuery{}, middleware.After)
}

func addOpDeleteWorkGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorkGroup{}, middleware.After)
}

func addOpGetDatabaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDatabase{}, middleware.After)
}

func addOpGetDataCatalogValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDataCatalog{}, middleware.After)
}

func addOpGetNamedQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNamedQuery{}, middleware.After)
}

func addOpGetQueryExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueryExecution{}, middleware.After)
}

func addOpGetQueryResultsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueryResults{}, middleware.After)
}

func addOpGetTableMetadataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTableMetadata{}, middleware.After)
}

func addOpGetWorkGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorkGroup{}, middleware.After)
}

func addOpListDatabasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDatabases{}, middleware.After)
}

func addOpListTableMetadataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTableMetadata{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartQueryExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartQueryExecution{}, middleware.After)
}

func addOpStopQueryExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopQueryExecution{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDataCatalogValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataCatalog{}, middleware.After)
}

func addOpUpdateWorkGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkGroup{}, middleware.After)
}

func validateEncryptionConfiguration(v *types.EncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionConfiguration"}
	if len(v.EncryptionOption) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionOption"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResultConfiguration(v *types.ResultConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResultConfiguration"}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResultConfigurationUpdates(v *types.ResultConfigurationUpdates) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResultConfigurationUpdates"}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateWorkGroupConfiguration(v *types.WorkGroupConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "WorkGroupConfiguration"}
	if v.ResultConfiguration != nil {
		if err := validateResultConfiguration(v.ResultConfiguration); err != nil {
			invalidParams.AddNested("ResultConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateWorkGroupConfigurationUpdates(v *types.WorkGroupConfigurationUpdates) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "WorkGroupConfigurationUpdates"}
	if v.ResultConfigurationUpdates != nil {
		if err := validateResultConfigurationUpdates(v.ResultConfigurationUpdates); err != nil {
			invalidParams.AddNested("ResultConfigurationUpdates", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetNamedQueryInput(v *BatchGetNamedQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetNamedQueryInput"}
	if v.NamedQueryIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NamedQueryIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetQueryExecutionInput(v *BatchGetQueryExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetQueryExecutionInput"}
	if v.QueryExecutionIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryExecutionIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataCatalogInput(v *CreateDataCatalogInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataCatalogInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNamedQueryInput(v *CreateNamedQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNamedQueryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateWorkGroupInput(v *CreateWorkGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWorkGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Configuration != nil {
		if err := validateWorkGroupConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataCatalogInput(v *DeleteDataCatalogInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataCatalogInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNamedQueryInput(v *DeleteNamedQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNamedQueryInput"}
	if v.NamedQueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NamedQueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteWorkGroupInput(v *DeleteWorkGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWorkGroupInput"}
	if v.WorkGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkGroup"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDatabaseInput(v *GetDatabaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDatabaseInput"}
	if v.CatalogName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogName"))
	}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDataCatalogInput(v *GetDataCatalogInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDataCatalogInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNamedQueryInput(v *GetNamedQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNamedQueryInput"}
	if v.NamedQueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NamedQueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueryExecutionInput(v *GetQueryExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueryExecutionInput"}
	if v.QueryExecutionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryExecutionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueryResultsInput(v *GetQueryResultsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueryResultsInput"}
	if v.QueryExecutionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryExecutionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTableMetadataInput(v *GetTableMetadataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTableMetadataInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.CatalogName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetWorkGroupInput(v *GetWorkGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkGroupInput"}
	if v.WorkGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkGroup"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDatabasesInput(v *ListDatabasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDatabasesInput"}
	if v.CatalogName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTableMetadataInput(v *ListTableMetadataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTableMetadataInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.CatalogName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartQueryExecutionInput(v *StartQueryExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartQueryExecutionInput"}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if v.ResultConfiguration != nil {
		if err := validateResultConfiguration(v.ResultConfiguration); err != nil {
			invalidParams.AddNested("ResultConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopQueryExecutionInput(v *StopQueryExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopQueryExecutionInput"}
	if v.QueryExecutionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryExecutionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDataCatalogInput(v *UpdateDataCatalogInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataCatalogInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWorkGroupInput(v *UpdateWorkGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkGroupInput"}
	if v.WorkGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkGroup"))
	}
	if v.ConfigurationUpdates != nil {
		if err := validateWorkGroupConfigurationUpdates(v.ConfigurationUpdates); err != nil {
			invalidParams.AddNested("ConfigurationUpdates", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
