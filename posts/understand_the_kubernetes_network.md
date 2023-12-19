---
title: 'draft-理解k8s网络'
date: '2023-12-18'
description: '理解k8s网络的工作原理'
---
本文旨在从问题出发探索一些k8s网络相关的工作原理。

## 单机上的两个Pod如何通信

![network_in_node.png](/images/understand_the_kubernetes_network/network_in_node.png)

* Pod通过**network namespace**与主机网络隔离
* Pod中存在一个虚拟以太网设备对(veth pair)，一端(eth0)在Pod的network namespace中，另一端(ethxx)在主机网络中
* Host的多个veth通过bridge网络相连，所以单机Pod间的通信本质上是多个虚拟设备构成的桥接网络

## 多机上的两个Pod如何通信

![network_cross_node.png](/images/understand_the_kubernetes_network/network_cross_node.png)

* 多机场景下，需要将不同主机网络相连

## 什么是Linux namespace

Linux namespace是内核提供的功能，是Linux容器化技术的基石，可以通过[unshare](https://man7.org/linux/man-pages/man1/unshare.1.html)命令调用。
![container](https://i0.wp.com/theboreddev.com/wp-content/uploads/2023/02/linux-container.png?w=451&ssl=1)

## 什么是虚拟以太网设备(veth)

linux虚拟以太网设备是[network interface](https://oldwiki.archive.openwrt.org/doc/networking/network.interfaces)的一种，可以在没有物理设备的情况下，对网络进行逻辑分区。

