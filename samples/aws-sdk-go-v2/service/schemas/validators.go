// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateDiscoverer struct {
}

func (*validateOpCreateDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDiscovererInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRegistry struct {
}

func (*validateOpCreateRegistry) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRegistry) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRegistryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRegistryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSchema struct {
}

func (*validateOpCreateSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDiscoverer struct {
}

func (*validateOpDeleteDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDiscovererInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRegistry struct {
}

func (*validateOpDeleteRegistry) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRegistry) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRegistryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRegistryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSchema struct {
}

func (*validateOpDeleteSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSchemaVersion struct {
}

func (*validateOpDeleteSchemaVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSchemaVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSchemaVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSchemaVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCodeBinding struct {
}

func (*validateOpDescribeCodeBinding) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCodeBinding) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCodeBindingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCodeBindingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDiscoverer struct {
}

func (*validateOpDescribeDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDiscovererInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRegistry struct {
}

func (*validateOpDescribeRegistry) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRegistry) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRegistryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRegistryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSchema struct {
}

func (*validateOpDescribeSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportSchema struct {
}

func (*validateOpExportSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCodeBindingSource struct {
}

func (*validateOpGetCodeBindingSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCodeBindingSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCodeBindingSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCodeBindingSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDiscoveredSchema struct {
}

func (*validateOpGetDiscoveredSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDiscoveredSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDiscoveredSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDiscoveredSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSchemas struct {
}

func (*validateOpListSchemas) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSchemas) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSchemasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSchemasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSchemaVersions struct {
}

func (*validateOpListSchemaVersions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSchemaVersions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSchemaVersionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSchemaVersionsInput(input); err != nil {
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

type validateOpPutCodeBinding struct {
}

func (*validateOpPutCodeBinding) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutCodeBinding) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutCodeBindingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutCodeBindingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutResourcePolicy struct {
}

func (*validateOpPutResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchSchemas struct {
}

func (*validateOpSearchSchemas) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchSchemas) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchSchemasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchSchemasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDiscoverer struct {
}

func (*validateOpStartDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDiscovererInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopDiscoverer struct {
}

func (*validateOpStopDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopDiscovererInput(input); err != nil {
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

type validateOpUpdateDiscoverer struct {
}

func (*validateOpUpdateDiscoverer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDiscoverer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDiscovererInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDiscovererInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateRegistry struct {
}

func (*validateOpUpdateRegistry) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateRegistry) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateRegistryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateRegistryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSchema struct {
}

func (*validateOpUpdateSchema) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSchemaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDiscoverer{}, middleware.After)
}

func addOpCreateRegistryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRegistry{}, middleware.After)
}

func addOpCreateSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSchema{}, middleware.After)
}

func addOpDeleteDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDiscoverer{}, middleware.After)
}

func addOpDeleteRegistryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRegistry{}, middleware.After)
}

func addOpDeleteSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSchema{}, middleware.After)
}

func addOpDeleteSchemaVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSchemaVersion{}, middleware.After)
}

func addOpDescribeCodeBindingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCodeBinding{}, middleware.After)
}

func addOpDescribeDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDiscoverer{}, middleware.After)
}

func addOpDescribeRegistryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRegistry{}, middleware.After)
}

func addOpDescribeSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSchema{}, middleware.After)
}

func addOpExportSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportSchema{}, middleware.After)
}

func addOpGetCodeBindingSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCodeBindingSource{}, middleware.After)
}

func addOpGetDiscoveredSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDiscoveredSchema{}, middleware.After)
}

func addOpListSchemasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSchemas{}, middleware.After)
}

func addOpListSchemaVersionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSchemaVersions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutCodeBindingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutCodeBinding{}, middleware.After)
}

func addOpPutResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResourcePolicy{}, middleware.After)
}

func addOpSearchSchemasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchSchemas{}, middleware.After)
}

func addOpStartDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDiscoverer{}, middleware.After)
}

func addOpStopDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopDiscoverer{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDiscovererValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDiscoverer{}, middleware.After)
}

func addOpUpdateRegistryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateRegistry{}, middleware.After)
}

func addOpUpdateSchemaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSchema{}, middleware.After)
}

func validateOpCreateDiscovererInput(v *CreateDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDiscovererInput"}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRegistryInput(v *CreateRegistryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRegistryInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSchemaInput(v *CreateSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSchemaInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.Content == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Content"))
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

func validateOpDeleteDiscovererInput(v *DeleteDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDiscovererInput"}
	if v.DiscovererId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DiscovererId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRegistryInput(v *DeleteRegistryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRegistryInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSchemaInput(v *DeleteSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSchemaInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSchemaVersionInput(v *DeleteSchemaVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSchemaVersionInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.SchemaVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCodeBindingInput(v *DescribeCodeBindingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCodeBindingInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.Language == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Language"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDiscovererInput(v *DescribeDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDiscovererInput"}
	if v.DiscovererId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DiscovererId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRegistryInput(v *DescribeRegistryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRegistryInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSchemaInput(v *DescribeSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSchemaInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportSchemaInput(v *ExportSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportSchemaInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCodeBindingSourceInput(v *GetCodeBindingSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCodeBindingSourceInput"}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.Language == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Language"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDiscoveredSchemaInput(v *GetDiscoveredSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDiscoveredSchemaInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Events == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Events"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSchemasInput(v *ListSchemasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSchemasInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSchemaVersionsInput(v *ListSchemaVersionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSchemaVersionsInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutCodeBindingInput(v *PutCodeBindingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutCodeBindingInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if v.Language == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Language"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutResourcePolicyInput(v *PutResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResourcePolicyInput"}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchSchemasInput(v *SearchSchemasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchSchemasInput"}
	if v.Keywords == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Keywords"))
	}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDiscovererInput(v *StartDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDiscovererInput"}
	if v.DiscovererId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DiscovererId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopDiscovererInput(v *StopDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopDiscovererInput"}
	if v.DiscovererId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DiscovererId"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdateDiscovererInput(v *UpdateDiscovererInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDiscovererInput"}
	if v.DiscovererId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DiscovererId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateRegistryInput(v *UpdateRegistryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRegistryInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSchemaInput(v *UpdateSchemaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSchemaInput"}
	if v.RegistryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistryName"))
	}
	if v.SchemaName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SchemaName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
