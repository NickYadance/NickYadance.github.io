## 单机上的两个 Pod 如何通信

![network_in_node.png](/images/understand_the_kubernetes_network/network_in_node.png)

* Pod 通过**network namespace**与主机网络隔离
* Pod 中存在一个虚拟以太网设备对（veth pair），一端（eth0）在 Pod 的 network namespace 中，另一端（ethxx）在主机网络中
* Host 的多个 veth 通过 bridge 网络相连，所以单机 Pod 间的通信本质上是多个虚拟设备构成的桥接网络

## 多机上的两个Pod如何通信

![network_cross_node.png](/images/understand_the_kubernetes_network/network_cross_node.png)

* 多机场景下，需要将不同主机网络相连

## Service 的原理

![services.png](/images/understand_the_kubernetes_network/services.png)

* services 本身与网络拓扑没有关系
* services 包括一个虚拟的 IP 地址
* kube-proxy（每个节点一个）根据 services 的定义修改 iptables 规则，将发往 services IP 的包，路由到对应的 pod 地址

## Ingress 的原理

![img.png](/images/understand_the_kubernetes_network/ingress.png)

* 与 `LoadBalancer` 不同，Ingress 只需要一个独立 IP 地址，可以路由到多个 Service
* Ingress 需要一个配套的 `IngressController`，来完成实际的流量分配
* IngressController 需要将自身端口暴露到集群外部以供访问，可能是通过 `NodePort`

## 什么是 Linux namespace

Linux namespace 是内核提供的功能，是 Linux 容器化技术的基石，可以通过 [unshare](https://man7.org/linux/man-pages/man1/unshare.1.html) 命令调用。
![container](https://i0.wp.com/theboreddev.com/wp-content/uploads/2023/02/linux-container.png?w=451&ssl=1)

## 什么是虚拟以太网设备（veth）

Linux 虚拟以太网设备是 [network interface](https://oldwiki.archive.openwrt.org/doc/networking/network.interfaces) 的一种，可以在没有物理设备的情况下，对网络进行逻辑分区。

参考：
* [kubernetes-in-action](https://www.manning.com/books/kubernetes-in-action)