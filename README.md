# Gofire

Gofire是一个基于golang编写的网络框架，支持自定义消息格式和传输协议。目前，内置了tcp和udp两个ConnGenerator。使用者可以轻松地重复使用这两者的功能，实现消息传输。 Gofire不限制消息传输的底层协议。理论上，用户可以通过任何网络传输协议传输消息，只需实现程序的ConnGenerator接口即可。

## 基本结构说明

```
├── README.md
├── core
│   ├── client.go # 内置的客户端类
│   ├── endpoint.go
│   ├── errors.go
│   ├── ifaces.go # 抽象的接口定义
│   ├── inflight_msg.go # 代表发送了但还未返回的消息
│   ├── mqueue_default.go # 默认消息队列
│   ├── mqueue_priority.go # 优先级消息队列
│   ├── pcodec.go # 网络包的编解码器
│   ├── server.go # 内置的服务端类
│   └── transport.go # 负责将写入的msg做encode和decode并投入的conn中
├── doc.md
├── example
│   ├── client
│   │   └── main.go
│   ├── proto
│   │   ├── mcodec.go
│   │   └── message.go
│   └── server
│       └── main.go
├── generator
│   ├── gen_tcp_client.go # 用于客户端建立tcp通道
│   ├── gen_tcp_server.go # 用于服务端监听tcp消息
│   ├── gen_udp_client.go # 用户客户端建立udp通道
│   └── gen_udp_server.go # 用于服务端监听udp消息
├── go.mod
└── go.sum
```

## 已实现功能

- [x] server监听tcp, udp请求
- [x] client端发送tcp, udp请求
- [x] 内置一个简单的packet_codec，用作网络消息编解码
- [x] 提供msg_codec接口，让使用者能够自定义应用消息编解码的逻辑
- [x] 提供conn_generator接口，让使用者能够自定义使用的传输协议

## 待实现功能

- [ ] 增加消息加密
- [ ] client与单个server之间可以建立多个连接，防止单个连接读写堵塞
- [ ] 对象池化
- [ ] 消息支持按优先级发送
- [ ] 增加websocket支持
- [ ] 完善日志记录
- [ ] 增加benchmark测试
- [ ] 网络中断、超时、连接失败等异常状态下恢复
- [ ] 支持配置server配置多个endpoint
