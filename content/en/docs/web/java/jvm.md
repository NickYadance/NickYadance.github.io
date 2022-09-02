---
title: "JVM"
description: ""
lead: ""
date: 2022-09-02T13:19:04+08:00
lastmod: 2022-09-02T13:19:04+08:00
draft: true
images: []
menu:
  docs:
    parent: ""
    identifier: "jvm-ea09810dbe490b8bf3fefeac568e102f"
weight: 999
toc: true
---
内存结构
堆内存
虚拟机把堆内存逻辑上划分成三块区域以优化 GC 性能
新生带 新对象和没达到一定年龄的对象都在新生代。
老年代 被长时间使用的对象，一般情况下老年代的内存空间比年轻代更大。
元空间 像一些方法中的操作临时对象等，JDK1.8 之前是占用 JVM 内存，JDK1.8 之后直接使用物理内存。
堆内存分配过程
分代与复制 #分代与复制
复制算法将内存划分为两个区间，在^^任意时间点，所有动态分配的对象都只能分配在其中一个区间，成为活动区间^^，而另外一个区间称为空闲区间。
当有效内存空间耗尽时，JVM将暂停程序运行，开启复制算法GC线程。接下来GC线程会将活动区间内的存活对象，全部复制到空闲区间，且严格按照内存地址依次排列，与此同时，GC线程将更新存活对象的内存引用地址指向新的内存地址。此时，空闲区间已经与活动区间交换，而垃圾对象现在已经全部留在了原来的活动区间，也就是现在的空闲区间。事实上，在活动区间转换为空间区间的同时，垃圾对象已经被一次性全部回收。
分配过程
新对象分配在Eden。
当Eden被填满，发生Minor GC，将伊甸园中的剩余对象移动到S0。
当再次发生Minor GC，会将S0复制到S1，并清空S0。
当S0|S1区域年龄达到阈值，复制内存至Old。
当Old内存不足时，发生Major GC。
当Major GC后Old内存仍不足以分配新对象，抛出OOM。
> 考虑一个问题，为什么需要S0和S1两个空间呢？

**`TLAB`** TLAB是JVM为每个线程分配的私有缓存区域。
**`方法区`** 方法区用于存储已被虚拟机加载的类型信息、常量、运行时常量、静态变量、即时编译器编译后的代码缓存等。
GC算法
标记清除
将存活的对象进行标记，然后清理掉未被标记的对象。
标记整理
让所有存活的对象都向一端移动，然后直接清理掉端边界以外的内存。
复制
将内存划分为大小相等的两块，每次只使用其中一块，当这一块内存用完了就将还存活的对象复制到另一块上面，然后再把使用过的内存空间进行一次清理。
现在的商业虚拟机都采用这种收集算法来回收新生代，但是并不是将新生代划分为大小相等的两块，而是分为一块较大的 Eden 空间和两块较小的 Survivor 空间，每次使用 Eden 空间和其中一块 Survivor。在回收时，将 Eden 和 Survivor 中还存活着的对象一次性复制到另一块 Survivor 空间上，最后清理 Eden 和使用过的那一块 Survivor。
HotSpot 虚拟机的 Eden 和 Survivor 的大小比例默认为 8:1，保证了内存的利用率达到 90%。如果每次回收有多于 10% 的对象存活，^^那么一块 Survivor 空间就不够用了，此时需要依赖于老年代进行分配担保^^，也就是借用老年代的空间存储放不下的对象。
> 为什么要有两个Survivor区域：From 和 To。
这个问题不应该是这样的，而是为什么HotSpot虚拟机为什么要增加Eden区域。
在复制算法中，有一个很大的缺点就在于堆的使用效率问题，如果是按照五五分成，总有一半是不能用的。那为什么不干脆八二分？那么问题来了，某次触发YGC，需要把From(8分)中的存活对象(假设占据1分)复制到To空间，这个时候，交换指针之后，From空间还剩(1分)可以使用，相信很快就会再次触发Ygc，STW发生频次提高
矛盾点：五五分成，堆利用率实在不高。不五五分成，由于复制算法的特性，STW的频次变高。为了解决内存分配空间它的capacity并不是恒定的这个问题，HotSpot引入了Eden区作为对复制算法的优化，其实eden区域可以看作是From和To区域的缓冲和共享区域。

混合分代收集
分代收集算法根据对象存活周期将内存划分为几块，^^不同块采用适合的收集算法^^。一般将堆分为新生代和老年代。新生代使用复制算法，老年代使用标记清除或者标记整理 算法。
GCRoot #gcroot
GCRoot指这样的类对象，即^^它们直接被JVM引用^^。
JVM中有四类对象可以作为GCRoot对象。
**Local variables** are kept alive by the stack of a thread. This is not a real object virtual reference and thus is not visible. For all intents and purposes, local variables are GC roots.
**Active Java threads** are always considered live objects and are therefore GC roots. This is especially important for thread local variables.
**Static variables** are referenced by their classes. This fact makes them de facto GC roots. Classes themselves can be garbage-collected, which would remove all referenced static variables. This is of special importance when we use application servers, OSGi containers or class loaders in general. We will discuss the related problems in the Problem Patterns section.
**JNI References** are Java objects that the native code has created as part of a JNI call. Objects thus created are treated specially because the JVM does not know if it is being referenced by the native code or not. Such objects represent a very special form of GC root, which we will examine in more detail in the Problem Patterns section below.
垃圾收集器
CMS *Concurrent Mark Sweep* #cms
初始标记: 仅标记GC Roots能直接关联到的对象，速度很快但需要停顿。
并发标记: 进行 GC Roots Tracing，在整个回收过程中耗时最长但不需要停顿。
重新标记: 修正并发标记期间因用户程序继续运作而导致标记产生变动对象的标记记录，需要停顿。
并发清除: 不需要停顿。
G1 *Garbage-First* #g1
G1 把堆划分成多个大小相等的**独立区域(Region)**，新生代和老年代不再物理隔离，将原来的整块内存空间划分成多个的小空间，使得每个小空间可以^^单独进行垃圾回收，使得可预测的停顿时间模型成为可能^^。
通过记录每个 Region 垃圾回收时间以及回收所获得的空间并维护一个优先列表，每次根据设置的收集时间，^^优先回收价值最大的 Region^^。
整体来看是基于标记 - 整理算法实现的收集器，从局部(两个 Region 之间)上来看是基于复制算法实现，运行期间不会产生内存空间碎片。
可预测的停顿: 能让使用者明确指定在一个长度为 M 毫秒的时间片段内，消耗在 GC 上的时间不得超过 N 毫秒。
三色标记法
CMS和G1在并发标记时使用的是同一个算法：三色标记法，使用黑白灰三种颜色标记对象。
白色：未标记。
灰色：自身被标记，引用的对象未标记。
黑色：自身与引用对象都已标记。
GC 开始前所有对象都是白色，GC 一开始所有根能够直达的对象被压到栈中，待搜索，此时颜色是灰色。然后灰色对象依次从栈中取出搜索子对象，子对象也会被涂为灰色，入栈。当其所有的子对象都涂为灰色之后该对象被涂为黑色。当 GC 结束之后灰色对象将全部没了，剩下黑色的为存活对象，白色的为垃圾。
重标记
在重标记中，黑色指向了白色，如果不对黑色重新扫描，则会漏标。如图，会把白色D对象当作没有新引用指向从而回收掉。要解决漏标问题，打破图中黄色区域所说的两个条件之一即可。
跟踪黑指向白的增量更新，关注引用的增加，把黑色重新标记为灰色，下次重新扫描属性。CMS采用该方法。
记录灰指向白的消失 SATB *snapshot at the beginning* 关注引用的删除，当灰–>白消失时，要把这个引用推到GC的堆栈，保证白还能被GC扫描到。G1采用该方法。
> 为什么G1采用SATB而不用incremental update？
因为采用incremental update把黑色重新标记为灰色后，之前扫描过的还要再扫描一遍，效率太低。G1有RSet与SATB相配合。Card Table里记录了RSet，**RSet里记录了其他对象指向自己的引用**，这样就不需要再扫描其他区域，只要扫描RSet就可以了。

调优

| 参数 | 类型 | 意义 | 指导配置 |
|---|---|---|---|
| -Xms | 基本参数 | 堆最小值(初始化值) | 与 -Xmx 配置为相等，避免内存震荡 |
| -Xmx | 基本参数 | 堆最大值 | 与 -Xms 配置为相等，避免内存震荡 |
| -XX:NewRatio | 分代配置参数 | 新生代与老年代比值 | 如果young gc非常频繁，考虑调大 |
| -XX:MaxTenuringThreshold | 分代配置参数 | 新生代晋升阈值 | 如果young gc非常频繁，考虑调大 |
| -XX:SurvivorRatio | 新生代配置参数 | Eden比值 | 如果Eden作为内存缓冲上涨过快，考虑调大 |
| -XX:+PrintGCDetails | 工具配置 | 打印详细GC日志 | |
| -XX:+PrintGCDateStamps | 工具配置 | 打印详细GC日期和时间 | |
| -XX:+UseG1GC | G1配置 | 启用G1收集器 | |
| -XX:+MaxGCPauseMillis=200 | G1配置 | 最大暂停时间 | |
