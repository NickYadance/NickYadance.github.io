---
layout: post
title:  "HBase-发现新大陆"
categories: HBase
date:   2018-08-8 22:14:54
tags:  Java HBase
---

* content
{:toc}

![nick-icon-smaller.png](https://i.loli.net/2018/08/08/5b6a63d67942f.png)

# HBase-优化初体验
> [Apache](http://www.apache.org/) HBase™ is the [Hadoop](http://hadoop.apache.org/) database, a distributed, scalable, big data store.  

## HBase配置项
> 本节转自：http://blog.csdn.net/odailidong/article/details/41794403

在与HBase打过招呼后，尝试对HBase作一些性能上的优化，核心内容是配置项，先贴一点在测试过程中参考的配置项：




```
1.hbase.regionserver.handler.count：rpc请求的线程数量，默认值是10，生产环境建议使用100，也不是越大越好，特别是当请求内容很大的时候，比如scan/put几M的数据，会占用过多的内存，有可能导致频繁的GC，甚至出现内存溢出。
2.hbase.master.distributed.log.splitting：默认值为true，建议设为false。关闭hbase的分布式日志切割，在log需要replay时，由master来负责重放
3.hbase.regionserver.hlog.splitlog.writer.threads：默认值是3，建议设为10，日志切割所用的线程数
4.hbase.snapshot.enabled：快照功能，默认是false(不开启)，建议设为true，特别是对某些关键的表，定时用快照做备份是一个不错的选择。
5.hbase.hregion.max.filesize：默认是10G， 如果任何一个column familiy里的StoreFile超过这个值, 那么这个Region会一分为二，因为region分裂会有短暂的region下线时间(通常在5s以内)，为减少对业务端的影响，建议手动定时分裂，可以设置为60G。
6.hbase.hregion.majorcompaction：hbase的region主合并的间隔时间，默认为1天，建议设置为0，禁止自动的major主合并，major合并会把一个store下所有的storefile重写为一个storefile文件，在合并过程中还会把有删除标识的数据删除，在生产集群中，主合并能持续数小时之久，为减少对业务的影响，建议在业务低峰期进行手动或者通过脚本或者api定期进行major合并。
7.hbase.hregion.memstore.flush.size：默认值128M，单位字节，一旦有memstore超过该值将被flush，如果regionserver的jvm内存比较充足(16G以上)，可以调整为256M。
8.hbase.hregion.memstore.block.multiplier：默认值2，如果一个memstore的内存大小已经超过hbase.hregion.memstore.flush.size *  hbase.hregion.memstore.block.multiplier，则会阻塞该memstore的写操作，为避免阻塞，建议设置为5，如果太大，则会有OOM的风险。如果在regionserver日志中出现"Blocking updates for '<threadName>' on region <regionName> : memstore size <多少M> is >= than blocking <多少M> size"的信息时，说明这个值该调整了。
9.hbase.hstore.compaction.min：默认值为3，如果任何一个store里的storefile总数超过该值，会触发默认的合并操作，可以设置5~8，在手动的定期major compact中进行storefile文件的合并，减少合并的次数，不过这会延长合并的时间，以前的对应参数为hbase.hstore.compactionThreshold。
10.hbase.hstore.compaction.max：默认值为10,一次最多合并多少个storefile，避免OOM。
11.hbase.hstore.blockingStoreFiles：默认为7，如果任何一个store(非.META.表里的store)的storefile的文件数大于该值，则在flush memstore前先进行split或者compact，同时把该region添加到flushQueue，延时刷新，这期间会阻塞写操作直到compact完成或者超过hbase.hstore.blockingWaitTime(默认90s)配置的时间，可以设置为30，避免memstore不及时flush。当regionserver运行日志中出现大量的“Region <regionName> has too many store files; delaying flush up to 90000ms"时，说明这个值需要调整了
12.hbase.regionserver.global.memstore.upperLimit：默认值0.4，regionserver所有memstore占用内存在总内存中的upper比例，当达到该值，则会从整个regionserver中找出最需要flush的region进行flush，直到总内存比例降到该数以下，采用默认值即可。
13.hbase.regionserver.global.memstore.lowerLimit：默认值0.35，采用默认值即可。
14.hbase.regionserver.thread.compaction.small：默认值为1，regionserver做Minor Compaction时线程池里线程数目,可以设置为5。
15.hbase.regionserver.thread.compaction.large：默认值为1，regionserver做Major Compaction时线程池里线程数目，可以设置为8。
16.hbase.regionserver.lease.period：默认值60000(60s)，客户端连接regionserver的租约超时时间，客户端必须在这个时间内汇报，否则则认为客户端已死掉。这个最好根据实际业务情况进行调整
17.hfile.block.cache.size：默认值0.25，regionserver的block cache的内存大小限制，在偏向读的业务中，可以适当调大该值，需要注意的是hbase.regionserver.global.memstore.upperLimit的值和hfile.block.cache.size的值之和必须小于0.8。
18.dfs.socket.timeout：默认值60000(60s)，建议根据实际regionserver的日志监控发现了异常进行合理的设置，比如我们设为900000，这个参数的修改需要同时更改hdfs-site.xml
19.dfs.datanode.socket.write.timeout：默认480000(480s),有时regionserver做合并时，可能会出现datanode写超时的情况，480000 millis timeout while waiting for channel to be ready for write，这个参数的修改需要同时更改hdfs-site.xml

jvm和垃圾收集参数：

export HBASE_REGIONSERVER_OPTS="-Xms36g -Xmx36g -Xmn1g -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:+UseCMSCompactAtFullCollection -XX:CMSFullGCsBeforeCompaction=15 -XX:CMSInitiatingOccupancyFraction=70 -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCTimeStamps -Xloggc:/data/logs/gc-$(hostname)-hbase.log"
```

## HBase配置项测试
* 测试工具：YCSB
* 测试流程：控制配置项变量；向数据库中插入300万条数据，做100万次查询，10次为一组；读写9:1

数据载入脚本：
```

#!/bin/bash
recordcount=3000000
operationcount=100000

ycsb load hbase20 \
-P ../workloads/workloadb \
-cp ~/hbase/hbase-2.1.0/conf/ \
-p table=user_portrait \
-p columnfamily=up \
-p recordcount=$recordcount \
-p operationcount=$operationcount \
-threads 50 \
-s
```
运行脚本：
```

#!/bin/bash
operation= {在这里填测试的配置项}
recordcount=3000000
operationcount=1000000

for i in {1..10}
do

ycsb run hbase20 \
-P ../workloads/workloadb \
-cp ~/hbase/hbase-2.1.0/conf/ \
-p table=user_portrait \
-p columnfamily=up \
-p measurementtype=hdrhistogram \
-p hdrhistogram.fileoutput=true \
-p hdrhistogram.output.path=./test \
-p recordcount=$recordcount \
-p operationcount=$operationcount \
-threads 50 \
-s \
| tee -a ${operation}-${recordcount}-${operationcount}-log.txt

done
```
在每组测试完成后得到log，用 RapidMiner Studio 提取出[READ]节点，导出到excel。以下是最明显的几个配置项

读平均延时

![avg](https://i.loli.net/2018/08/10/5b6d7591b4b77.png)


读95%百分比延时

![95th](https://i.loli.net/2018/08/10/5b6d759344309.png)

1. 提高rpc处理线程和block缓存比例是可以有效提高读性能的，这个是意料之中。但是线程数不能开得太高
2. Java GC合理配置也能提高读性能
3. 针对hdfs的优化反倒使性能降低了，可能是参数调太大了
4. 采用Bucketcache作为L2缓存提性能效果也可以


HBase还有非常多其他配置，做这个主要是熟悉一下其工作原理，为生产环境的调试作准备。
