## request/limit

request 的值有三个作用：

1. 影响调度策略
2. 影响cpu资源配比
3. 影响oom打分值
`
> The scheduler ensures that, for each resource type, the sum of the resource requests of the scheduled containers is less than the capacity of the node. Note that although actual memory or CPU resource usage on nodes is very low, the scheduler still refuses to place a Pod on a node if the capacity check fails.

request 在调度策略中的作用可以总结为：节点上所有容器的 request 总和不能超过节点容量。这是一个保守的硬性条件，并不会考虑作决策时节点的容器真实资源使用率，这也是实际使用中优化 Kubernetes 集群资源使用率的一个点，毕竟多数容器会设定一个大的预估值以应对峰值流量。

cpu.request 会影响运行时，容器可能分配到的 CPU 时间片比值，即两个容器的实际 CPU 资源使用率比值与 cpu.request 的比值相等。
![cpu_request](/images/understand_the_kubernetes_resource/cpu_request.png)

memory.request并不是在容器一开始就分配到的内存值大小，与cpu.request类似memory.request主要用于调度策略。内存相对于cpu属于不可压缩的“静态资源”，在分配给进程后内存不能动态的伸缩，只能由进程GC或者被操作系统完整回收。memory.request会参与到操作系统oom打分的计算，memory.request使用占比越高的容器越容易被oom。