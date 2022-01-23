# cc rpc

## 相对 go rpc 优势

- 使用function接口而非反射, 提高调用效率
- client 支持服务发现
- client 支持连接池, 可以同时连接多个服务端
- client 支持断链重连, 而 go rpc 的client一旦连接断掉后不在重连, 调用Call会直接报错


NewServer 结构因为有 ServeHTTP 函数所以实现了Handler接口
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

ServeHTTP 包含1个结构体 NewServerSession,接受2个参数 NewServer 返回的实例和 ServeHTTP 的ResponseWriter 参数经过Hijacker之后的连接管道。

