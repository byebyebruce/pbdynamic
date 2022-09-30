# pbdynamic
动态解析 protobuf 类型

## 原理解释  
通过加载pb 描述符文件(protoc --descriptor_set_out=xxx.pb.desc xxx.proto)，不需要重新编译xxx.pb.go的情况下动态的构造pb类型，  
进而反序列化字节，并打印信息

## 快速使用
1. 生成filedesc文件
* [示例](testdata/gen.sh)
```shell
protoc --descriptor_set_out=xxx.pb.desc xxx.proto
```
2. 依赖包
```shell
go get github.com/byebyebruce/pbdynamic@latest
```

3. 解析从字节动态反序列化pb并打印
```go
err := pbdynamic.LoadFiles("xxx.pb.desc")
if err != nil {
    panic(err)
}

// bs 是从其他地方来的pb序列化后的字节
fmt.Println(pbdynamic.String(bs, "mypkg.SomeTypeName"))
```