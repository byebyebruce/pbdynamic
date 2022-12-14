// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: testdata.proto

package testdata

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

type DummyEnum int32

const (
	DummyEnum_TYPE_DUMMY_0 DummyEnum = 0
	DummyEnum_TYPE_DUMMY_1 DummyEnum = 1
	DummyEnum_TYPE_DUMMY_2 DummyEnum = 2
)

// Enum value maps for DummyEnum.
var (
	DummyEnum_name = map[int32]string{
		0: "TYPE_DUMMY_0",
		1: "TYPE_DUMMY_1",
		2: "TYPE_DUMMY_2",
	}
	DummyEnum_value = map[string]int32{
		"TYPE_DUMMY_0": 0,
		"TYPE_DUMMY_1": 1,
		"TYPE_DUMMY_2": 2,
	}
)

func (x DummyEnum) Enum() *DummyEnum {
	p := new(DummyEnum)
	*p = x
	return p
}

func (x DummyEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DummyEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_proto_enumTypes[0].Descriptor()
}

func (DummyEnum) Type() protoreflect.EnumType {
	return &file_testdata_proto_enumTypes[0]
}

func (x DummyEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DummyEnum.Descriptor instead.
func (DummyEnum) EnumDescriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

type DummyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DummyMsg) Reset() {
	*x = DummyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyMsg) ProtoMessage() {}

func (x *DummyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyMsg.ProtoReflect.Descriptor instead.
func (*DummyMsg) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *DummyMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PBMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValDouble   float64   `protobuf:"fixed64,1,opt,name=val_double,json=valDouble,proto3" json:"val_double,omitempty"`
	ValFloat    float32   `protobuf:"fixed32,2,opt,name=val_float,json=valFloat,proto3" json:"val_float,omitempty"`
	ValInt64    int64     `protobuf:"varint,3,opt,name=val_int64,json=valInt64,proto3" json:"val_int64,omitempty"`
	ValUint64   uint64    `protobuf:"varint,4,opt,name=val_uint64,json=valUint64,proto3" json:"val_uint64,omitempty"`
	ValInt32    int32     `protobuf:"varint,5,opt,name=val_int32,json=valInt32,proto3" json:"val_int32,omitempty"`
	ValFixed64  uint64    `protobuf:"fixed64,6,opt,name=val_fixed64,json=valFixed64,proto3" json:"val_fixed64,omitempty"`
	ValFixed32  uint32    `protobuf:"fixed32,7,opt,name=val_fixed32,json=valFixed32,proto3" json:"val_fixed32,omitempty"`
	ValBool     bool      `protobuf:"varint,8,opt,name=val_bool,json=valBool,proto3" json:"val_bool,omitempty"`
	ValString   string    `protobuf:"bytes,9,opt,name=val_string,json=valString,proto3" json:"val_string,omitempty"`
	ValMessage  *DummyMsg `protobuf:"bytes,10,opt,name=val_message,json=valMessage,proto3" json:"val_message,omitempty"`
	ValBytes    []byte    `protobuf:"bytes,11,opt,name=val_bytes,json=valBytes,proto3" json:"val_bytes,omitempty"`
	ValUint32   uint32    `protobuf:"varint,12,opt,name=val_uint32,json=valUint32,proto3" json:"val_uint32,omitempty"`
	ValEnum     DummyEnum `protobuf:"varint,13,opt,name=val_enum,json=valEnum,proto3,enum=testdata.DummyEnum" json:"val_enum,omitempty"`
	ValSfixed32 int32     `protobuf:"fixed32,14,opt,name=val_sfixed32,json=valSfixed32,proto3" json:"val_sfixed32,omitempty"`
	ValSfixed64 int64     `protobuf:"fixed64,15,opt,name=val_sfixed64,json=valSfixed64,proto3" json:"val_sfixed64,omitempty"`
	ValSint32   int32     `protobuf:"zigzag32,16,opt,name=val_sint32,json=valSint32,proto3" json:"val_sint32,omitempty"`
	ValSint64   int64     `protobuf:"zigzag64,17,opt,name=val_sint64,json=valSint64,proto3" json:"val_sint64,omitempty"`
}

func (x *PBMsg) Reset() {
	*x = PBMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBMsg) ProtoMessage() {}

func (x *PBMsg) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBMsg.ProtoReflect.Descriptor instead.
func (*PBMsg) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *PBMsg) GetValDouble() float64 {
	if x != nil {
		return x.ValDouble
	}
	return 0
}

func (x *PBMsg) GetValFloat() float32 {
	if x != nil {
		return x.ValFloat
	}
	return 0
}

func (x *PBMsg) GetValInt64() int64 {
	if x != nil {
		return x.ValInt64
	}
	return 0
}

func (x *PBMsg) GetValUint64() uint64 {
	if x != nil {
		return x.ValUint64
	}
	return 0
}

func (x *PBMsg) GetValInt32() int32 {
	if x != nil {
		return x.ValInt32
	}
	return 0
}

func (x *PBMsg) GetValFixed64() uint64 {
	if x != nil {
		return x.ValFixed64
	}
	return 0
}

func (x *PBMsg) GetValFixed32() uint32 {
	if x != nil {
		return x.ValFixed32
	}
	return 0
}

func (x *PBMsg) GetValBool() bool {
	if x != nil {
		return x.ValBool
	}
	return false
}

func (x *PBMsg) GetValString() string {
	if x != nil {
		return x.ValString
	}
	return ""
}

func (x *PBMsg) GetValMessage() *DummyMsg {
	if x != nil {
		return x.ValMessage
	}
	return nil
}

func (x *PBMsg) GetValBytes() []byte {
	if x != nil {
		return x.ValBytes
	}
	return nil
}

func (x *PBMsg) GetValUint32() uint32 {
	if x != nil {
		return x.ValUint32
	}
	return 0
}

func (x *PBMsg) GetValEnum() DummyEnum {
	if x != nil {
		return x.ValEnum
	}
	return DummyEnum_TYPE_DUMMY_0
}

func (x *PBMsg) GetValSfixed32() int32 {
	if x != nil {
		return x.ValSfixed32
	}
	return 0
}

func (x *PBMsg) GetValSfixed64() int64 {
	if x != nil {
		return x.ValSfixed64
	}
	return 0
}

func (x *PBMsg) GetValSint32() int32 {
	if x != nil {
		return x.ValSint32
	}
	return 0
}

func (x *PBMsg) GetValSint64() int64 {
	if x != nil {
		return x.ValSint64
	}
	return 0
}

var File_testdata_proto protoreflect.FileDescriptor

var file_testdata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x04, 0x0a, 0x05, 0x50,
	0x42, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09,
	0x76, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c,
	0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61,
	0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x76,
	0x61, 0x6c, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76,
	0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4d, 0x73, 0x67, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61,
	0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76,
	0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x6c,
	0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x2e, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x76,
	0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0b, 0x76, 0x61,
	0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c,
	0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x10, 0x52,
	0x0b, 0x76, 0x61, 0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x10, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x11, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x2a, 0x41, 0x0a, 0x09, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x55, 0x4d, 0x4d, 0x59, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x55, 0x4d, 0x4d, 0x59, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x55, 0x4d, 0x4d, 0x59, 0x5f, 0x32, 0x10, 0x02, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x65, 0x62,
	0x79, 0x65, 0x62, 0x72, 0x75, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_proto_rawDescOnce sync.Once
	file_testdata_proto_rawDescData = file_testdata_proto_rawDesc
)

func file_testdata_proto_rawDescGZIP() []byte {
	file_testdata_proto_rawDescOnce.Do(func() {
		file_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_proto_rawDescData)
	})
	return file_testdata_proto_rawDescData
}

var file_testdata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testdata_proto_goTypes = []interface{}{
	(DummyEnum)(0),   // 0: testdata.DummyEnum
	(*DummyMsg)(nil), // 1: testdata.DummyMsg
	(*PBMsg)(nil),    // 2: testdata.PBMsg
}
var file_testdata_proto_depIdxs = []int32{
	1, // 0: testdata.PBMsg.val_message:type_name -> testdata.DummyMsg
	0, // 1: testdata.PBMsg.val_enum:type_name -> testdata.DummyEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_testdata_proto_init() }
func file_testdata_proto_init() {
	if File_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyMsg); i {
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
		file_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBMsg); i {
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
			RawDescriptor: file_testdata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_proto_goTypes,
		DependencyIndexes: file_testdata_proto_depIdxs,
		EnumInfos:         file_testdata_proto_enumTypes,
		MessageInfos:      file_testdata_proto_msgTypes,
	}.Build()
	File_testdata_proto = out.File
	file_testdata_proto_rawDesc = nil
	file_testdata_proto_goTypes = nil
	file_testdata_proto_depIdxs = nil
}
