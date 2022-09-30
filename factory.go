package pbdynamic

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

// DescriptorFactory 类型描述符工厂
type DescriptorFactory struct {
	fs []*protoregistry.Files
}

// NewDescriptorFactory 构造
// pbDescFiles: protobuf 描述符文件 protoc --descriptor_set_out=xx.pb.desc
func NewDescriptorFactory(pbDescFiles ...string) (*DescriptorFactory, error) {
	ret := &DescriptorFactory{
		fs: []*protoregistry.Files{},
	}
	for _, v := range pbDescFiles {
		if err := ret.LoadFile(v); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

// LoadFile 读protobuf 描述符文件
func (pfb *DescriptorFactory) LoadFile(f string) error {
	b, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	var fileSet descriptor.FileDescriptorSet
	if err := proto.Unmarshal(b, &fileSet); err != nil {
		return err
	}
	return pfb.Add(&fileSet)
}

// Add 读protobuf 描述符文件
func (pfb *DescriptorFactory) Add(set *descriptor.FileDescriptorSet) error {
	files, err := protodesc.NewFiles(set)
	if err != nil {
		return err
	}
	pfb.fs = append(pfb.fs, files)
	return nil
}

// FindDesc 查找描述符
// fullname： 全名 testdata.MyType
func (pfb *DescriptorFactory) FindDesc(fullname string) protoreflect.Descriptor {
	for _, v := range pfb.fs {
		ret, _ := v.FindDescriptorByName(protoreflect.FullName(fullname))
		if ret != nil {
			return ret
		}
	}
	return nil
}

// NewMessage 动态构造一个message
func (pfb *DescriptorFactory) NewMessage(fullname string) *dynamicpb.Message {
	d := pfb.FindDesc(fullname)
	if d == nil {
		return nil
	}
	return dynamicpb.NewMessage(d.(protoreflect.MessageDescriptor))
}

// NewMessageFromBytes 反序列化
func (pfb *DescriptorFactory) NewMessageFromBytes(b []byte, fullname string) (*dynamicpb.Message, error) {
	m := pfb.NewMessage(fullname)
	if m == nil {
		return nil, fmt.Errorf("no type")
	}
	err := proto.Unmarshal(b, m)
	return m, err
}
