---
title: "Insertion Sort"
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
## Loop Variant
Let's learn about how to prove the algorithm right with `Loop Variant`, which contains 4 parts:
$$
\begin{cases}
Variant(V): &T(i) \\\\
Start(S): &T(0) \\\\
Loop(L): &T(i+1)=f(T(i)) \\\\
End(E): &Answer=f(T(n)) \\\\
\end{cases}
$$

For insertion sort
$$
\begin{cases}
V: &T(i)=sorted(A[0,i]) \\\\
S: &T(0)=sorted(A[0,0]) \\\\
L: &T(i+1)=orderd\\_insertion(T(i), A[i+1]) \\\\
E:  &Answer=T(n)=sorted(A[0,n]) \\\\
\end{cases}
$$

Pseudocode
```
InsertionSort(A):
  for i=1 to A.length:
    j=i-1
    while A[j]>A[i] and j >= 0:
      A[j+1]=A[j]
      j=j-1
    A[j+1]=A[i]
```
## Analysis
![anayze](images/insertion_sort_analyze.png)
For the best condition when A is ascending sorted
$$t_j=1,\ T(n)=an+b$$
For the worst condition when A is descending sorted
$$t_j=j,\ T(n)=an^2+bn+c$$
