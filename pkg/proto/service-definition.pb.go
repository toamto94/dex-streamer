// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: service-definition.proto

package proto

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

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint       string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ChainId        string `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Address        string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Abi            string `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty"`
	ScrapeInterval uint32 `protobuf:"varint,5,opt,name=scrapeInterval,proto3" json:"scrapeInterval,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_service_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_service_definition_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Contract) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *Contract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Contract) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *Contract) GetScrapeInterval() uint32 {
	if x != nil {
		return x.ScrapeInterval
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tbd string `protobuf:"bytes,1,opt,name=tbd,proto3" json:"tbd,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_definition_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetTbd() string {
	if x != nil {
		return x.Tbd
	}
	return ""
}

var File_service_definition_proto protoreflect.FileDescriptor

var file_service_definition_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x62, 0x69, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x62, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x62, 0x64, 0x32,
	0x39, 0x0a, 0x0b, 0x44, 0x45, 0x58, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x09, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_definition_proto_rawDescOnce sync.Once
	file_service_definition_proto_rawDescData = file_service_definition_proto_rawDesc
)

func file_service_definition_proto_rawDescGZIP() []byte {
	file_service_definition_proto_rawDescOnce.Do(func() {
		file_service_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_definition_proto_rawDescData)
	})
	return file_service_definition_proto_rawDescData
}

var file_service_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_definition_proto_goTypes = []interface{}{
	(*Contract)(nil), // 0: Contract
	(*Response)(nil), // 1: Response
}
var file_service_definition_proto_depIdxs = []int32{
	0, // 0: DEXStreamer.StreamContract:input_type -> Contract
	1, // 1: DEXStreamer.StreamContract:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_definition_proto_init() }
func file_service_definition_proto_init() {
	if File_service_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_service_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_service_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_definition_proto_goTypes,
		DependencyIndexes: file_service_definition_proto_depIdxs,
		MessageInfos:      file_service_definition_proto_msgTypes,
	}.Build()
	File_service_definition_proto = out.File
	file_service_definition_proto_rawDesc = nil
	file_service_definition_proto_goTypes = nil
	file_service_definition_proto_depIdxs = nil
}
