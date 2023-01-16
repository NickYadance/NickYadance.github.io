---
title: "Dag"
description: ""
lead: ""
date: 2022-09-15T22:09:07+08:00
lastmod: 2022-09-15T22:09:07+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "dag-0229befd436efdbb633063ff449d678b"
weight: 999
toc: true
---
Shuffle&Exchange
```
scala> val n2 = spark.range(1, 1000000)
scala> val n2split = n2.repartition(7);
scala> n2split.take(2).foreach(println)
```
![spark_ui_1](images/spark_ui_1.png)
* stage2 run `take` in one task, or in one concurrency

Partitions
```
scala> val ds1 = spark.range(1, 1000000)
scala> val ds2 = spark.range(1, 1000000, 2)
scala> val ds3 = ds1.repartition(7)
scala> val ds4 = ds2.repartition(9)
scala> val ds5 = ds3.selectExpr("id * 5 as id")
scala> val joined = ds5.join(ds4, "id")
scala> val sum = joined.selectExpr("sum(id)")
```

![spark_ui_2](images/spark_ui_2.png)

* Job3 & Job4 both run `range() -> partition()`, 10 tasks, producing partitions
* Job5 reads partitions with 9 tasks, broadcast result in memory cuz the result is small enough
* Job6 reads partitions with 7 tasks, run `id * 5 -> join(in memory)` ![spark_ui_3](images/spark_ui_3.png)
* Job7 run `take` with 1 task
## Reference
[Nice Gitbook about Spark](https://spark-internals.books.yourtion.com/markdown/1-Overview.html)
[How To Read Spark DAGs](https://www.youtube.com/watch?v=LoFN_Q224fQ)
