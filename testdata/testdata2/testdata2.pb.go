// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: testdata2/testdata2.proto

package testdata2

import (
	testdata "github.com/byebyebruce/pbdynamic/testdata"
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

type TestData1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *testdata.PBMsg `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *TestData1) Reset() {
	*x = TestData1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata2_testdata2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestData1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestData1) ProtoMessage() {}

func (x *TestData1) ProtoReflect() protoreflect.Message {
	mi := &file_testdata2_testdata2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestData1.ProtoReflect.Descriptor instead.
func (*TestData1) Descriptor() ([]byte, []int) {
	return file_testdata2_testdata2_proto_rawDescGZIP(), []int{0}
}

func (x *TestData1) GetA() *testdata.PBMsg {
	if x != nil {
		return x.A
	}
	return nil
}

var File_testdata2_testdata2_proto protoreflect.FileDescriptor

var file_testdata2_testdata2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x32, 0x1a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x31, 0x12, 0x1d, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x42, 0x4d, 0x73, 0x67, 0x52,
	0x01, 0x61, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x79, 0x65, 0x62, 0x79, 0x65, 0x62, 0x72, 0x75, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x32, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata2_testdata2_proto_rawDescOnce sync.Once
	file_testdata2_testdata2_proto_rawDescData = file_testdata2_testdata2_proto_rawDesc
)

func file_testdata2_testdata2_proto_rawDescGZIP() []byte {
	file_testdata2_testdata2_proto_rawDescOnce.Do(func() {
		file_testdata2_testdata2_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata2_testdata2_proto_rawDescData)
	})
	return file_testdata2_testdata2_proto_rawDescData
}

var file_testdata2_testdata2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testdata2_testdata2_proto_goTypes = []interface{}{
	(*TestData1)(nil),      // 0: testdata2.TestData1
	(*testdata.PBMsg)(nil), // 1: testdata.PBMsg
}
var file_testdata2_testdata2_proto_depIdxs = []int32{
	1, // 0: testdata2.TestData1.a:type_name -> testdata.PBMsg
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testdata2_testdata2_proto_init() }
func file_testdata2_testdata2_proto_init() {
	if File_testdata2_testdata2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata2_testdata2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestData1); i {
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
			RawDescriptor: file_testdata2_testdata2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata2_testdata2_proto_goTypes,
		DependencyIndexes: file_testdata2_testdata2_proto_depIdxs,
		MessageInfos:      file_testdata2_testdata2_proto_msgTypes,
	}.Build()
	File_testdata2_testdata2_proto = out.File
	file_testdata2_testdata2_proto_rawDesc = nil
	file_testdata2_testdata2_proto_goTypes = nil
	file_testdata2_testdata2_proto_depIdxs = nil
}
