package pbdynamic_test

import (
	"fmt"

	"github.com/byebyebruce/pbdynamic"
	"github.com/byebyebruce/pbdynamic/testdata"
	"github.com/byebyebruce/pbdynamic/testdata/testdata2"
	"google.golang.org/protobuf/proto"
)

func Example() {
	err := pbdynamic.LoadFiles(protosetPath)
	if err != nil {
		panic(err)
	}

	// 模拟生成一些测试数据
	tmp := &testdata2.TestData1{
		A: &testdata.PBMsg{
			ValBool:   true,
			ValString: "okok",
			ValInt64:  321312312324324,
			ValBytes:  []byte("123453321"),
		},
	}
	b, _ := proto.Marshal(tmp)

	// 通过类型名字和字节动态反序列化
	fmt.Println(pbdynamic.String(b, messageName))
}
