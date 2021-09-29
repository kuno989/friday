// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: clamav.proto

package clamav_api

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

type ScanFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *ScanFileRequest) Reset() {
	*x = ScanFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clamav_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanFileRequest) ProtoMessage() {}

func (x *ScanFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clamav_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanFileRequest.ProtoReflect.Descriptor instead.
func (*ScanFileRequest) Descriptor() ([]byte, []int) {
	return file_clamav_proto_rawDescGZIP(), []int{0}
}

func (x *ScanFileRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output   string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Infected bool   `protobuf:"varint,2,opt,name=infected,proto3" json:"infected,omitempty"`
	Update   int64  `protobuf:"varint,3,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clamav_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clamav_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_clamav_proto_rawDescGZIP(), []int{1}
}

func (x *ScanResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ScanResponse) GetInfected() bool {
	if x != nil {
		return x.Infected
	}
	return false
}

func (x *ScanResponse) GetUpdate() int64 {
	if x != nil {
		return x.Update
	}
	return 0
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clamav_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clamav_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_clamav_proto_rawDescGZIP(), []int{2}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clamav_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clamav_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_clamav_proto_rawDescGZIP(), []int{3}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_clamav_proto protoreflect.FileDescriptor

var file_clamav_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x5a, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x32, 0x8d, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x6d, 0x41, 0x56, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76,
	0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6c, 0x61, 0x6d, 0x61, 0x76, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x75, 0x6e, 0x6f, 0x39, 0x38, 0x39, 0x2f, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79, 0x3b, 0x63, 0x6c,
	0x61, 0x6d, 0x61, 0x76, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clamav_proto_rawDescOnce sync.Once
	file_clamav_proto_rawDescData = file_clamav_proto_rawDesc
)

func file_clamav_proto_rawDescGZIP() []byte {
	file_clamav_proto_rawDescOnce.Do(func() {
		file_clamav_proto_rawDescData = protoimpl.X.CompressGZIP(file_clamav_proto_rawDescData)
	})
	return file_clamav_proto_rawDescData
}

var file_clamav_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clamav_proto_goTypes = []interface{}{
	(*ScanFileRequest)(nil), // 0: clamav.ScanFileRequest
	(*ScanResponse)(nil),    // 1: clamav.ScanResponse
	(*VersionRequest)(nil),  // 2: clamav.VersionRequest
	(*VersionResponse)(nil), // 3: clamav.VersionResponse
}
var file_clamav_proto_depIdxs = []int32{
	0, // 0: clamav.ClamAVScanner.ScanFile:input_type -> clamav.ScanFileRequest
	2, // 1: clamav.ClamAVScanner.GetVersion:input_type -> clamav.VersionRequest
	1, // 2: clamav.ClamAVScanner.ScanFile:output_type -> clamav.ScanResponse
	3, // 3: clamav.ClamAVScanner.GetVersion:output_type -> clamav.VersionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clamav_proto_init() }
func file_clamav_proto_init() {
	if File_clamav_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clamav_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanFileRequest); i {
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
		file_clamav_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
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
		file_clamav_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_clamav_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_clamav_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clamav_proto_goTypes,
		DependencyIndexes: file_clamav_proto_depIdxs,
		MessageInfos:      file_clamav_proto_msgTypes,
	}.Build()
	File_clamav_proto = out.File
	file_clamav_proto_rawDesc = nil
	file_clamav_proto_goTypes = nil
	file_clamav_proto_depIdxs = nil
}
