// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/encoding/json"
	"github.com/awslabs/smithy-go/middleware"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpBatchMeterUsage struct {
}

func (*awsAwsjson11_serializeOpBatchMeterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpBatchMeterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchMeterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.BatchMeterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentBatchMeterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpMeterUsage struct {
}

func (*awsAwsjson11_serializeOpMeterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpMeterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*MeterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.MeterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentMeterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpRegisterUsage struct {
}

func (*awsAwsjson11_serializeOpRegisterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpRegisterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RegisterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.RegisterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentRegisterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpResolveCustomer struct {
}

func (*awsAwsjson11_serializeOpResolveCustomer) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpResolveCustomer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ResolveCustomerInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.ResolveCustomer")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentResolveCustomerInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentUsageAllocation(v *types.UsageAllocation, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AllocatedUsageQuantity != nil {
		ok := object.Key("AllocatedUsageQuantity")
		ok.Integer(*v.AllocatedUsageQuantity)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson11_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentUsageAllocations(v []types.UsageAllocation, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentUsageAllocation(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentUsageRecord(v *types.UsageRecord, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomerIdentifier != nil {
		ok := object.Key("CustomerIdentifier")
		ok.String(*v.CustomerIdentifier)
	}

	if v.Dimension != nil {
		ok := object.Key("Dimension")
		ok.String(*v.Dimension)
	}

	if v.Quantity != nil {
		ok := object.Key("Quantity")
		ok.Integer(*v.Quantity)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.UsageAllocations != nil {
		ok := object.Key("UsageAllocations")
		if err := awsAwsjson11_serializeDocumentUsageAllocations(v.UsageAllocations, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentUsageRecordList(v []types.UsageRecord, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentUsageRecord(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentBatchMeterUsageInput(v *BatchMeterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.UsageRecords != nil {
		ok := object.Key("UsageRecords")
		if err := awsAwsjson11_serializeDocumentUsageRecordList(v.UsageRecords, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentMeterUsageInput(v *MeterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DryRun != nil {
		ok := object.Key("DryRun")
		ok.Boolean(*v.DryRun)
	}

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.UsageAllocations != nil {
		ok := object.Key("UsageAllocations")
		if err := awsAwsjson11_serializeDocumentUsageAllocations(v.UsageAllocations, ok); err != nil {
			return err
		}
	}

	if v.UsageDimension != nil {
		ok := object.Key("UsageDimension")
		ok.String(*v.UsageDimension)
	}

	if v.UsageQuantity != nil {
		ok := object.Key("UsageQuantity")
		ok.Integer(*v.UsageQuantity)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentRegisterUsageInput(v *RegisterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Nonce != nil {
		ok := object.Key("Nonce")
		ok.String(*v.Nonce)
	}

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.PublicKeyVersion != nil {
		ok := object.Key("PublicKeyVersion")
		ok.Integer(*v.PublicKeyVersion)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentResolveCustomerInput(v *ResolveCustomerInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RegistrationToken != nil {
		ok := object.Key("RegistrationToken")
		ok.String(*v.RegistrationToken)
	}

	return nil
}
