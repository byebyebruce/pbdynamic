package pbdynamic

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

var DefaultFactor = &DescriptorFactory{
	fs: []*protoregistry.Files{},
}

// LoadFiles 读protobuf 描述符文件
// pbDescFiles: protobuf 描述符文件 protoc --descriptor_set_out=xx.pb.desc
func LoadFiles(pbDescFiles ...string) error {
	for _, v := range pbDescFiles {
		if err := DefaultFactor.LoadFile(v); err != nil {
			return err
		}
	}
	return nil
}

// NewMessageFromBytes 从字节new message
// fullname： 全名 testdata.MyType
func NewMessageFromBytes(b []byte, fullname string) (*dynamicpb.Message, error) {
	m := DefaultFactor.NewMessage(fullname)
	if m == nil {
		return nil, fmt.Errorf("no type")
	}
	err := proto.Unmarshal(b, m)
	return m, err
}

// String 打印信息
func String(b []byte, fullname string) string {
	m, err := DefaultFactor.NewMessageFromBytes(b, fullname)
	if err != nil {
		return err.Error()
	}
	return m.String()
}
