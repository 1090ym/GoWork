package main

import (
	"GoWork/thrift/gen-go/demo"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

//定义服务
type Greeter struct {
}

//实现IDL里定义的接口
//SayHello
func (this *Greeter) SayHello(ctx context.Context, u *demo.User) (r *demo.Response, err error) {
	strJson, _ := json.Marshal(u)
	return &demo.Response{Errcode: 0, Errmsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

//GetUser
func (this *Greeter) GetUser(ctx context.Context, uid int32) (r *demo.Response, err error) {
	return &demo.Response{Errcode: 1, Errmsg: "user not exist."}, nil
}

func main() {
	//命令行参数
	flag.Usage = Usage
	addr := flag.String("addr", "192.168.40.124:9090", "Address to listen to")
	flag.Parse()

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//transport
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	//handler
	handler := &Greeter{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Println("error running server:", err)
	}

	//processor
	processor := demo.NewGreeterProcessor(handler)

	fmt.Println("Starting the simple server... on ", *addr)

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}
}
