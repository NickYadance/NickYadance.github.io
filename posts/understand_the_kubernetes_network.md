---
title: '理解k8s网络'
date: '2023-12-18'
description: '理解k8s网络的工作原理'
---
## 单机上的两个Pod如何通信

![network_in_node.png](/images/understand_the_kubernetes_network/network_in_node.png)

* Pod通过**network namespace**与主机网络隔离
* Pod中存在一个虚拟以太网设备对(veth pair)，一端(eth0)在Pod的network namespace中，另一端(ethxx)在主机网络中
* Host的多个veth通过bridge网络相连，所以单机Pod间的通信本质上是多个虚拟设备构成的桥接网络

## 多机上的两个Pod如何通信

![network_cross_node.png](/images/understand_the_kubernetes_network/network_cross_node.png)

* 多机场景下，需要将不同主机网络相连

## Service 的原理
![services.png](/images/understand_the_kubernetes_network/services.png)
* services本身与网络拓扑没有关系
* services包括一个虚拟的ip地址
* kube-proxy(one per node)根据services的定义修改iptables规则，将发往services ip的包，路由到对应的pod地址

## Ingress 的原理
![img.png](/images/understand_the_kubernetes_network/ingress.png)
* 与`LoadBalancer`不同，Ingress只需要一个独立Ip地址，可以路由到多个Service
* Ingress需要一个配套的`IngressController`，来完成实际的流量分配
* IngressController需要将自身端口暴露到集群外部以供访问，可能是通过`NodePort`

## 什么是Linux namespace

Linux namespace是内核提供的功能，是Linux容器化技术的基石，可以通过[unshare](https://man7.org/linux/man-pages/man1/unshare.1.html)命令调用。
![container](https://i0.wp.com/theboreddev.com/wp-content/uploads/2023/02/linux-container.png?w=451&ssl=1)

## 什么是虚拟以太网设备(veth)

linux虚拟以太网设备是[network interface](https://oldwiki.archive.openwrt.org/doc/networking/network.interfaces)的一种，可以在没有物理设备的情况下，对网络进行逻辑分区。

> * [kubernetes-in-action](https://www.manning.com/books/kubernetes-in-action)