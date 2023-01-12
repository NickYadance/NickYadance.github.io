---
title: "Mergesort"
description: ""
lead: ""
date: 2023-01-11T15:27:52+08:00
lastmod: 2023-01-11T15:27:52+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "mergesort-9d010f5fd3d4b892c5f9b7bedc320a4c"
weight: 999
toc: true
---
## Divide And Conquer
1. Divide: divide the original problem to several smaller sub problems
2. Resolve: resolve sub problems if they are small enough, or else goto step1
3. Conquer: merge the result of sub problems

## Merge Sort
```
MergeSort(A,l,r):
if l < r:
  m=(l+r)/2
  MergeSort(A,l,m)
  MergeSort(A,m,r)
  Merge(A,l,m,r)
```
## Analyze
For merge sort, when $n=2^n$:
$$
T(n)=
\begin{cases}
1, &n=1 \\\\
2T(n/2)+cn &n>1 \\\\
\end{cases}
$$
The `recursive tree` of the merge sort complexity: ![recursive tree](images/recursive_tree.png)
