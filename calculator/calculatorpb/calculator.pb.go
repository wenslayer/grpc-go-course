// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Operands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ints []int64 `protobuf:"zigzag64,1,rep,packed,name=ints,proto3" json:"ints,omitempty"`
}

func (x *Operands) Reset() {
	*x = Operands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operands) ProtoMessage() {}

func (x *Operands) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operands.ProtoReflect.Descriptor instead.
func (*Operands) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Operands) GetInts() []int64 {
	if x != nil {
		return x.Ints
	}
	return nil
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operands *Operands `protobuf:"bytes,1,opt,name=operands,proto3" json:"operands,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SumRequest) GetOperands() *Operands {
	if x != nil {
		return x.Operands
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"zigzag64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SumResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeFactorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"zigzag64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeFactorizationRequest) Reset() {
	*x = PrimeFactorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeFactorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorizationRequest) ProtoMessage() {}

func (x *PrimeFactorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorizationRequest.ProtoReflect.Descriptor instead.
func (*PrimeFactorizationRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeFactorizationRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeFactorizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Factor int64 `protobuf:"zigzag64,1,opt,name=factor,proto3" json:"factor,omitempty"`
}

func (x *PrimeFactorizationResponse) Reset() {
	*x = PrimeFactorizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeFactorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorizationResponse) ProtoMessage() {}

func (x *PrimeFactorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorizationResponse.ProtoReflect.Descriptor instead.
func (*PrimeFactorizationResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *PrimeFactorizationResponse) GetFactor() int64 {
	if x != nil {
		return x.Factor
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"zigzag64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *AverageRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float64 `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *AverageResponse) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type MaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"zigzag64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MaximumRequest) Reset() {
	*x = MaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumRequest) ProtoMessage() {}

func (x *MaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumRequest.ProtoReflect.Descriptor instead.
func (*MaximumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *MaximumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int64 `protobuf:"zigzag64,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *MaximumResponse) Reset() {
	*x = MaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumResponse) ProtoMessage() {}

func (x *MaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumResponse.ProtoReflect.Descriptor instead.
func (*MaximumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{8}
}

func (x *MaximumResponse) GetMaximum() int64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x1e, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x12,
	0x52, 0x04, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3e, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a,
	0x19, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x34, 0x0a, 0x1a, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x06, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x28, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0xc8, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x03,
	0x53, 0x75, 0x6d, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x46, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x19, 0x5a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculatorpb_calculator_proto_rawDescData = file_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*Operands)(nil),                   // 0: calculator.Operands
	(*SumRequest)(nil),                 // 1: calculator.SumRequest
	(*SumResponse)(nil),                // 2: calculator.SumResponse
	(*PrimeFactorizationRequest)(nil),  // 3: calculator.PrimeFactorizationRequest
	(*PrimeFactorizationResponse)(nil), // 4: calculator.PrimeFactorizationResponse
	(*AverageRequest)(nil),             // 5: calculator.AverageRequest
	(*AverageResponse)(nil),            // 6: calculator.AverageResponse
	(*MaximumRequest)(nil),             // 7: calculator.MaximumRequest
	(*MaximumResponse)(nil),            // 8: calculator.MaximumResponse
}
var file_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.SumRequest.operands:type_name -> calculator.Operands
	1, // 1: calculator.CalculatorService.Sum:input_type -> calculator.SumRequest
	3, // 2: calculator.CalculatorService.PrimeFactorization:input_type -> calculator.PrimeFactorizationRequest
	5, // 3: calculator.CalculatorService.Average:input_type -> calculator.AverageRequest
	7, // 4: calculator.CalculatorService.Maximum:input_type -> calculator.MaximumRequest
	2, // 5: calculator.CalculatorService.Sum:output_type -> calculator.SumResponse
	4, // 6: calculator.CalculatorService.PrimeFactorization:output_type -> calculator.PrimeFactorizationResponse
	6, // 7: calculator.CalculatorService.Average:output_type -> calculator.AverageResponse
	8, // 8: calculator.CalculatorService.Maximum:output_type -> calculator.MaximumResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_calculatorpb_calculator_proto_init() }
func file_calculator_calculatorpb_calculator_proto_init() {
	if File_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operands); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeFactorizationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeFactorizationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculatorpb_calculator_proto = out.File
	file_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeFactorization(ctx context.Context, in *PrimeFactorizationRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorizationClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeFactorization(ctx context.Context, in *PrimeFactorizationRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorizationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeFactorization", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeFactorizationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeFactorizationClient interface {
	Recv() (*PrimeFactorizationResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeFactorizationClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeFactorizationClient) Recv() (*PrimeFactorizationResponse, error) {
	m := new(PrimeFactorizationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/Maximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaximumClient{stream}
	return x, nil
}

type CalculatorService_MaximumClient interface {
	Send(*MaximumRequest) error
	Recv() (*MaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaximumClient) Send(m *MaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaximumClient) Recv() (*MaximumResponse, error) {
	m := new(MaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeFactorization(*PrimeFactorizationRequest, CalculatorService_PrimeFactorizationServer) error
	Average(CalculatorService_AverageServer) error
	Maximum(CalculatorService_MaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeFactorization(*PrimeFactorizationRequest, CalculatorService_PrimeFactorizationServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeFactorization not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServiceServer) Maximum(CalculatorService_MaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method Maximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeFactorization_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeFactorizationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeFactorization(m, &calculatorServicePrimeFactorizationServer{stream})
}

type CalculatorService_PrimeFactorizationServer interface {
	Send(*PrimeFactorizationResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeFactorizationServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeFactorizationServer) Send(m *PrimeFactorizationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Maximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Maximum(&calculatorServiceMaximumServer{stream})
}

type CalculatorService_MaximumServer interface {
	Send(*MaximumResponse) error
	Recv() (*MaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaximumServer) Send(m *MaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaximumServer) Recv() (*MaximumRequest, error) {
	m := new(MaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeFactorization",
			Handler:       _CalculatorService_PrimeFactorization_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Maximum",
			Handler:       _CalculatorService_Maximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
