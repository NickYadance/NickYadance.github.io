---
title: "Insertionsort"
description: ""
lead: ""
date: 2023-01-11T14:39:50+08:00
lastmod: 2023-01-11T14:39:50+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "insertionsort-26832919a29379355e037dd2afa238ec"
weight: 999
toc: true
---
Let's learn about how to prove the algorithm right with `Loop Variant`, which contains 4 parts:
* $ Variant(V): $ The variant formula
* $ Start(S): $ The variant is true on start
* $ L(Loop): $ The variant is true before each loop
* $ E(End): $ When loop ends, the variant proves the algorithm right

Take insertion sort for example:
* $ V:\ A[0,j)\ is\ sorted,\ j \in [0,A.length)$
* $ S:\ A[0,j)\ is\ sorted,\ j=1$
* $ L:\ A[0,j+1)\ is\ sorted,\ j \rightarrow j+1$
* $ E:\ A[0,n)\ is\ sorted,\ j=n$

Pseudocode
```
InsertionSort(A):
  for j=1 to A.length:
    i=j-1
    while A[i]>A[j] and i >= 0:
      A[i+1]=A[i]
      i=i-1
    A[i+1]=A[j]
```
