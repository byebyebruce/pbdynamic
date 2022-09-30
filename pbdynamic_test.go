package pbdynamic_test

import (
	"testing"

	"github.com/byebyebruce/pbdynamic"
	"github.com/byebyebruce/pbdynamic/testdata"
	"github.com/byebyebruce/pbdynamic/testdata/testdata2"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

const (
	protosetPath = "testdata/testdata.pb.desc"
	messageName  = "testdata2.TestData1"
)

var (
	samplePB = &testdata2.TestData1{
		A: &testdata.PBMsg{
			ValBool:   true,
			ValString: "okok",
			ValInt64:  321312312324324,
			ValBytes:  []byte("123453321"),
		},
	}
)

func TestNewMessageFromBytes(t *testing.T) {
	err := pbdynamic.LoadFiles(protosetPath)
	assert.Nil(t, err)
	// sample
	b, _ := proto.Marshal(samplePB)

	m, err := pbdynamic.NewMessageFromBytes(b, messageName)
	assert.Nil(t, err)
	assert.NotNil(t, m)
	assert.EqualValues(t, samplePB.String(), m.String())
}
