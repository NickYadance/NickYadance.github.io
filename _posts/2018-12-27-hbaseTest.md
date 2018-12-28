---
layout: post
title:  "HBase 跨RS使用分页过滤器"
categories: HBase
date:   2018-11-28 17:20:54
tags:  Java HBase
---

* content
{:toc}

![apachelogo.png](https://nickyadance.github.io/img/apachelogo.png)

## 问题
某业务需要对HBase一张大表做全表扫描，为了避免对业务产生影响，同时保证扫描效率，需采用多线程+分页的方式进行扫描。
参照<<HBase权威指南>>一书，使用PageFilter进行扫描，却发现10多亿的数据很快就扫完了，远远短于预估时间。
经过一番调试，终于找到原因，详细的解决过程见：
[HBASE-21332](https://issues.apache.org/jira/browse/HBASE-21332)

近日把上次的问题翻出来，想起之前测试的是HBase 1.2，最新的HBase2.0是不是对分页过滤器作了优化以避免上述问题呢？
于是把代码扒下来，改成2.1.1版本，测试结果如下：

![ut2x.png](https://nickyadance.github.io/img/result.png)

1. 第一个region上只取了一个分页大小数量的数据，而1.2版本中会先把第一个region扫描完
2. 一次扫描取到了所有region上的数据，而1.2版本会依次扫描数据

结果和1.2版本有些许不同，但是分页过滤器的行为更符合其定义：
```
/**
 * Implementation of Filter interface that limits results to a specific page
 * size. It terminates scanning once the number of filter-passed rows is &gt;
 * the given page size.
 * <p>
 * Note that this filter cannot guarantee that the number of results returned
 * to a client are &lt;= page size. This is because the filter is applied
 * separately on different region servers. It does however optimize the scan of
 * individual HRegions by making sure that the page size is never exceeded
 * locally.
 */
```
HBase是分布式的，分页过滤器在不同RS(region server)之间没有一致性保障。当在多个RS上使用分页过滤器，返回的结果
是作用在所有RS上的分页过滤器的共同结果，即：**N * PAGE_SIZE**，其中N是扫描范围跨过的RS数量。这样就能对丢数据给出合理的解释了。

而1.1版本中，第一个region会被扫描完，在结束时才会出现于2.0版本相同的结果。我估计是1.1版本中对合并的结果作了处理，
限制结果大小为PAGE_SIZE，当然这些纯属猜测。

## 建议的使用方式
1. 手动限制单个scan的范围在同一个region上，或者同一个RS上。可以参考测试代码中获取region范围的方式。
2. （待测试）不使用分页过滤器，客户端限定获取的数据条数。