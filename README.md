
```golang
// 哪一部分的内容是交给用户自定义
// 那一部分内容是交由框架来控制

// 典型的时候场景就是，客户端想发送一个消息，仅需指定消息本身和传输的协议
// 考虑消息传递过程中可以由用户自定义消息混淆方式，消息的编码，解码方式可以交给用户自定义
// 消息发送方和接收方都需要知道从流里读取多少内容，因此每次读多少是需要预先知道的
// 读的数量和读的内容是一起的
// |协议名称-version-内容长度|

type TransProtocol struct {
    Name string
    Version uint32
    ContentLength uint32
}

// 设置消息的传输方式
// 0. 设置传输的协议
// 1. 决定消息如何编码
// 2. 将消息写入底层的网络连接中
// 3. 从网络连接中等待并取出消息
// 4. 将消息解码并传递给调用方
type ITransport interface {
    TransProtocol() TransProtocol
}

func NewTransport(protocol string, codec MsgCodec)


func initClient() {
	// 从全局系统上注册好channel的generator，client使用的时候仅需指定使用的传输方式
	gofire.Regist()

	// 对于client而言，client的职责仅仅是将用户发送的消息传入我们抽象出来的管道中
	// 并等待从管道中收到回复
	// client不应该关心管道的generator
	client := gofire.NewClient()

    client.SetTransport()
    client.SetServerEndpoint()
	// client将message通过某种途径传输

	client.Send(msg)
}

func initServer() {
    server := gofire.NewServer()

    // 设置支持的传输途径
    server.SetTransport()
    server.SetEndpoint()

    server.Listen()
}
```

// 1. 要想ch被复用，则可以将消息排队执行
