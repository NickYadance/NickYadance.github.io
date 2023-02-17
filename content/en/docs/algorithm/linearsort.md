---
title: "Linearsort"
description: ""
lead: ""
date: 2023-02-17T20:36:31+08:00
lastmod: 2023-02-17T20:36:31+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "linearsort-e5e66c5b3b67f1e27391e1ba1c06b9fd"
weight: 999
toc: true
---

## Decision tree in sort algorithms
The **decision tree** can be used to describe all effective sort algorithms based on:
* Each leaf node is one possible answer to a sort algorithm
* Leaf node num >= $N!,N=size$, so the algorithm can get all sorts orders
* O(algo) = length of longest_simple_path in decision tree

![img.png](images/decisiontree.png)

## Countingsort
The `countingsort` is easy to understand but also easy to misunderstand. The countingsort, based on the fact
that the input number is limited to $N$, counts all numbers and relay them back in sorted order.

However, the key thing about countingsort is to keep the relative order in input, alias **stability**.So there is no sense writing countingsort as below
even if it works.
```text
MeaninglessCountingsort(A):
    B,C = new array
    for i = 1 to A.length:
        ++B[A[i]]
    k = 0
    for i = 1 to B.length:
        while --B[i] > 0:
            A[k++] = i

Countingsort(A):
    B = array
    for i = 1 to A.length:
        ++B[A[i]]
    for i = 2 to B.length:
        B[i] = B[i] + B[i - 1]
    for i = A.length to 1:
        C[B[A[i]]--] = A[i]
```

## Radixsort
The radixsort is most useful in sorting numbers with same length, by sorting each column from right to left.

It is based on the fact that countingsort(or other sort algorithm) is stable.

![img.png](images/radixsort.png)

## Bucketsort
Similar to countingsort,the input in bucketsort is distributed in a certain range, say [0, 1], but is uncountable. So
the bucketsort distributes the inputs into bucket in which an ordered list receive them, and then combine results of all
the buckets(lists).

![img.png](images/bucketsort.png)
