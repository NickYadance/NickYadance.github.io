---
title: 'Unix Socket'
date: '2023-12-21'
description: 'Unix Socket'
---

## Socket Domain and Type
socket domain包括`unix domain socket`和`internet domain socket`。
![domain](/images/unix_socket/domain.png)

unix domain socket数据通过内核传输，通过套接字文件寻址，不经过网络协议栈(如TCP/IP)。
![unix_domain_socket_io](/images/unix_socket/unix_domain_socket_io.png)

internet domain socket数据通过网络传输，通过网络层寻址(Ipv4/Ipv6)，常见的传输方式基于TCP协议。
![internet_domain_socket_io.png](/images/unix_socket/internet_domain_socket_io.png)

socket类型包括`stream`和`datagram`，`stream` socket的主要特性为：
* 可靠性：发包方和收包方内容完全一致，不乱序、丢包
* 双向：发包方和收包方可双向传输数据
* 无界：包数据无边界

![type](/images/unix_socket/type.png)

## unix domain socket的使用场景
适用于同一个unix操作系统上的进程间通信，由于不经过网络栈，传输效率比通过loopback的单机网络传输要高。

## internet domain socket是基于TCP/IP协议的吗？
可能是。internet domain socket可以基于Ipv4/Ipv6进行网络寻址，TCP协议簇进行网络传输，但是也可以基于其他协议。socket
对内核的网络栈操作进行了抽象，方便开发者的使用，开发者可以自行选择合适的协议簇。

## 与Http/Websocket协议的关系
没有关系。Http是常见的应用层协议，基于TCP/IP连接做**单向**数据传输，同时多数Http连接只会进行一次数据收发，随后便断开连接，这种方式称为`短连接`。
相对的，可以设置Keep Alive参数进行单个Http(TCP)连接上的多次数据收发，称为`长连接`。

由于单向传输，Server端只负责收包，Client端只负责发包。Server/Client端可以利用socket与内核交互操作TCP/IP连接，但也可以通过其他方式，针对场景获得更好的网络性能。

Websocket是Http连接的升级版，与stream socket类似，可以进行双向无界字节流传输。

## 与grpc的关系
grpc也是跨平台的应用层协议，不同的是grpc以二进制数据进行数据传输，同时原生支持stream，grpc也支持直接通过socket套接字进行数据传输。

> * [the-linux-programming-interface-michael-kerrisk-1](https://sciencesoftcode.files.wordpress.com/2018/12/the-linux-programming-interface-michael-kerrisk-1.pdf)