---
layout: post
title:  "设计模式"
categories: DesignPattern
date:   2018-10-30 17:20:54
tags:  Java
---

* content
{:toc}

![dplogo.jpg](https://nickyadance.github.io/img/设计模式.jpg)

## 建造者模式

* 简单工厂

![simpleFatory](https://nickyadance.github.io/img/简单工厂.png)

* 工厂方法

![simpleFatory](https://nickyadance.github.io/img/工厂方法.png)

* 抽象工厂

![simpleFatory](https://nickyadance.github.io/img/抽象方法.png)

* 单例

JDK 1.5之后，最简洁的单例模式是使用枚举：
```
public enum Singleton {
  INSTANCE;
}
```
相比其他单例模式，其有以下优点：

1. 编码简单
2. 不用担心因为序列化引起的问题，JVM会保证枚举类的序列化
3. 枚举类的创建是语言级别上保证的线程安全

## 建造者模式

* 适配器模式

![adapter](https://nickyadance.github.io/img/适配器模式.png)  

* 装饰器模式

![decorator](https://nickyadance.github.io/img/装饰器模式.png)