// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ptypes.proto

package types

import (
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

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_ptypes_proto_rawDescGZIP(), []int{0}
}

type GetInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuID int32 `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_ptypes_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvoiceRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

type AggregateDistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuID int32   `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Unix  int64   `protobuf:"varint,3,opt,name=Unix,proto3" json:"Unix,omitempty"`
}

func (x *AggregateDistanceRequest) Reset() {
	*x = AggregateDistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateDistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateDistanceRequest) ProtoMessage() {}

func (x *AggregateDistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateDistanceRequest.ProtoReflect.Descriptor instead.
func (*AggregateDistanceRequest) Descriptor() ([]byte, []int) {
	return file_ptypes_proto_rawDescGZIP(), []int{2}
}

func (x *AggregateDistanceRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

func (x *AggregateDistanceRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AggregateDistanceRequest) GetUnix() int64 {
	if x != nil {
		return x.Unix
	}
	return 0
}

var File_ptypes_proto protoreflect.FileDescriptor

var file_ptypes_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06,
	0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x62, 0x75, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62, 0x75, 0x49,
	0x44, 0x22, 0x5a, 0x0a, 0x18, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4f, 0x62, 0x75, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62,
	0x75, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x78, 0x32, 0x4b, 0x0a,
	0x12, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ptypes_proto_rawDescOnce sync.Once
	file_ptypes_proto_rawDescData = file_ptypes_proto_rawDesc
)

func file_ptypes_proto_rawDescGZIP() []byte {
	file_ptypes_proto_rawDescOnce.Do(func() {
		file_ptypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptypes_proto_rawDescData)
	})
	return file_ptypes_proto_rawDescData
}

var file_ptypes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ptypes_proto_goTypes = []interface{}{
	(*None)(nil),                     // 0: None
	(*GetInvoiceRequest)(nil),        // 1: GetInvoiceRequest
	(*AggregateDistanceRequest)(nil), // 2: AggregateDistanceRequest
}
var file_ptypes_proto_depIdxs = []int32{
	2, // 0: DistanceAggregator.AggregateDistance:input_type -> AggregateDistanceRequest
	0, // 1: DistanceAggregator.AggregateDistance:output_type -> None
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ptypes_proto_init() }
func file_ptypes_proto_init() {
	if File_ptypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ptypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*None); i {
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
		file_ptypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceRequest); i {
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
		file_ptypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateDistanceRequest); i {
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
			RawDescriptor: file_ptypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ptypes_proto_goTypes,
		DependencyIndexes: file_ptypes_proto_depIdxs,
		MessageInfos:      file_ptypes_proto_msgTypes,
	}.Build()
	File_ptypes_proto = out.File
	file_ptypes_proto_rawDesc = nil
	file_ptypes_proto_goTypes = nil
	file_ptypes_proto_depIdxs = nil
}
