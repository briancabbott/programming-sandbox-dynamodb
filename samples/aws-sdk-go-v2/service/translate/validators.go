// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpDeleteTerminology struct {
}

func (*validateOpDeleteTerminology) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTerminology) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTerminologyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTerminologyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTextTranslationJob struct {
}

func (*validateOpDescribeTextTranslationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTextTranslationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTerminology struct {
}

func (*validateOpGetTerminology) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTerminology) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTerminologyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTerminologyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportTerminology struct {
}

func (*validateOpImportTerminology) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportTerminology) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportTerminologyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportTerminologyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTextTranslationJob struct {
}

func (*validateOpStartTextTranslationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTextTranslationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopTextTranslationJob struct {
}

func (*validateOpStopTextTranslationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopTextTranslationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTranslateText struct {
}

func (*validateOpTranslateText) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTranslateText) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TranslateTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTranslateTextInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteTerminologyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTerminology{}, middleware.After)
}

func addOpDescribeTextTranslationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTextTranslationJob{}, middleware.After)
}

func addOpGetTerminologyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTerminology{}, middleware.After)
}

func addOpImportTerminologyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportTerminology{}, middleware.After)
}

func addOpStartTextTranslationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTextTranslationJob{}, middleware.After)
}

func addOpStopTextTranslationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopTextTranslationJob{}, middleware.After)
}

func addOpTranslateTextValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTranslateText{}, middleware.After)
}

func validateEncryptionKey(v *types.EncryptionKey) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionKey"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
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

func validateInputDataConfig(v *types.InputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputDataConfig"}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutputDataConfig(v *types.OutputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputDataConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTerminologyData(v *types.TerminologyData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TerminologyData"}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if v.File == nil {
		invalidParams.Add(smithy.NewErrParamRequired("File"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTerminologyInput(v *DeleteTerminologyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTerminologyInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTextTranslationJobInput(v *DescribeTextTranslationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTextTranslationJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTerminologyInput(v *GetTerminologyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTerminologyInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.TerminologyDataFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("TerminologyDataFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportTerminologyInput(v *ImportTerminologyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportTerminologyInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.TerminologyData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TerminologyData"))
	} else if v.TerminologyData != nil {
		if err := validateTerminologyData(v.TerminologyData); err != nil {
			invalidParams.AddNested("TerminologyData", err.(smithy.InvalidParamsError))
		}
	}
	if v.EncryptionKey != nil {
		if err := validateEncryptionKey(v.EncryptionKey); err != nil {
			invalidParams.AddNested("EncryptionKey", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.MergeStrategy) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MergeStrategy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTextTranslationJobInput(v *StartTextTranslationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTextTranslationJobInput"}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.TargetLanguageCodes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetLanguageCodes"))
	}
	if v.SourceLanguageCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceLanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopTextTranslationJobInput(v *StopTextTranslationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopTextTranslationJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTranslateTextInput(v *TranslateTextInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TranslateTextInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if v.TargetLanguageCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetLanguageCode"))
	}
	if v.SourceLanguageCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceLanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
