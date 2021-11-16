include "user.thrift"

// 标记各语言的命名空间（包名），不同语言需要单独声明
namespace go demo

// 重新定义类型名称，同c语言
typedef map<string, string> Data

// 定义响应体结构
struct Response {
    1:required i32 errcode,
    2:required string errmsg,
    3:required Data data,
}

// 定义服务接口，相当于go的interface
service Greeter {
    Response SayHello(
        1:required user.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}