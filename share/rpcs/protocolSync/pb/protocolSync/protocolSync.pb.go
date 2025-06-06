// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: proto/protocolSync.proto

package protocolSync

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocolSync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocolSync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_protocolSync_proto_rawDescGZIP(), []int{0}
}

type SyncDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"` //
}

func (x *SyncDeviceReq) Reset() {
	*x = SyncDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocolSync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncDeviceReq) ProtoMessage() {}

func (x *SyncDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocolSync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncDeviceReq.ProtoReflect.Descriptor instead.
func (*SyncDeviceReq) Descriptor() ([]byte, []int) {
	return file_proto_protocolSync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncDeviceReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type SyncDeviceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceNames []string `protobuf:"bytes,1,rep,name=deviceNames,proto3" json:"deviceNames,omitempty"`
}

func (x *SyncDeviceResp) Reset() {
	*x = SyncDeviceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocolSync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncDeviceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncDeviceResp) ProtoMessage() {}

func (x *SyncDeviceResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocolSync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncDeviceResp.ProtoReflect.Descriptor instead.
func (*SyncDeviceResp) Descriptor() ([]byte, []int) {
	return file_proto_protocolSync_proto_rawDescGZIP(), []int{2}
}

func (x *SyncDeviceResp) GetDeviceNames() []string {
	if x != nil {
		return x.DeviceNames
	}
	return nil
}

var File_proto_protocolSync_proto protoreflect.FileDescriptor

var file_proto_protocolSync_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x53, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2d, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x22, 0x32, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x32, 0x90, 0x01, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x37, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53,
	0x79, 0x6e, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_protocolSync_proto_rawDescOnce sync.Once
	file_proto_protocolSync_proto_rawDescData = file_proto_protocolSync_proto_rawDesc
)

func file_proto_protocolSync_proto_rawDescGZIP() []byte {
	file_proto_protocolSync_proto_rawDescOnce.Do(func() {
		file_proto_protocolSync_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocolSync_proto_rawDescData)
	})
	return file_proto_protocolSync_proto_rawDescData
}

var file_proto_protocolSync_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_protocolSync_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: protocolSync.Empty
	(*SyncDeviceReq)(nil),  // 1: protocolSync.SyncDeviceReq
	(*SyncDeviceResp)(nil), // 2: protocolSync.SyncDeviceResp
}
var file_proto_protocolSync_proto_depIdxs = []int32{
	0, // 0: protocolSync.protocolSync.SyncProduct:input_type -> protocolSync.Empty
	1, // 1: protocolSync.protocolSync.SyncDevice:input_type -> protocolSync.SyncDeviceReq
	0, // 2: protocolSync.protocolSync.SyncProduct:output_type -> protocolSync.Empty
	2, // 3: protocolSync.protocolSync.SyncDevice:output_type -> protocolSync.SyncDeviceResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_protocolSync_proto_init() }
func file_proto_protocolSync_proto_init() {
	if File_proto_protocolSync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocolSync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_protocolSync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncDeviceReq); i {
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
		file_proto_protocolSync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncDeviceResp); i {
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
			RawDescriptor: file_proto_protocolSync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_protocolSync_proto_goTypes,
		DependencyIndexes: file_proto_protocolSync_proto_depIdxs,
		MessageInfos:      file_proto_protocolSync_proto_msgTypes,
	}.Build()
	File_proto_protocolSync_proto = out.File
	file_proto_protocolSync_proto_rawDesc = nil
	file_proto_protocolSync_proto_goTypes = nil
	file_proto_protocolSync_proto_depIdxs = nil
}
