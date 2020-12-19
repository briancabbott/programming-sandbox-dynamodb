// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpDeleteLexicon struct {
}

func (*validateOpDeleteLexicon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLexicon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLexiconInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLexiconInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLexicon struct {
}

func (*validateOpGetLexicon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLexicon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLexiconInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLexiconInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSpeechSynthesisTask struct {
}

func (*validateOpGetSpeechSynthesisTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSpeechSynthesisTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSpeechSynthesisTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSpeechSynthesisTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutLexicon struct {
}

func (*validateOpPutLexicon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutLexicon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutLexiconInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutLexiconInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSpeechSynthesisTask struct {
}

func (*validateOpStartSpeechSynthesisTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSpeechSynthesisTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSpeechSynthesisTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSpeechSynthesisTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSynthesizeSpeech struct {
}

func (*validateOpSynthesizeSpeech) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSynthesizeSpeech) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SynthesizeSpeechInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSynthesizeSpeechInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteLexiconValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLexicon{}, middleware.After)
}

func addOpGetLexiconValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLexicon{}, middleware.After)
}

func addOpGetSpeechSynthesisTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSpeechSynthesisTask{}, middleware.After)
}

func addOpPutLexiconValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutLexicon{}, middleware.After)
}

func addOpStartSpeechSynthesisTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSpeechSynthesisTask{}, middleware.After)
}

func addOpSynthesizeSpeechValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSynthesizeSpeech{}, middleware.After)
}

func validateOpDeleteLexiconInput(v *DeleteLexiconInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLexiconInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLexiconInput(v *GetLexiconInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLexiconInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSpeechSynthesisTaskInput(v *GetSpeechSynthesisTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSpeechSynthesisTaskInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutLexiconInput(v *PutLexiconInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutLexiconInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Content == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Content"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSpeechSynthesisTaskInput(v *StartSpeechSynthesisTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSpeechSynthesisTaskInput"}
	if len(v.OutputFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OutputFormat"))
	}
	if v.OutputS3BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputS3BucketName"))
	}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if len(v.VoiceId) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("VoiceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSynthesizeSpeechInput(v *SynthesizeSpeechInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SynthesizeSpeechInput"}
	if len(v.OutputFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OutputFormat"))
	}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if len(v.VoiceId) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("VoiceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
