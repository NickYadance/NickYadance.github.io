---
title: '理解快速排序的分治问题'
date: '2023-11-11'
description: ''
---
快速排序是一个分治算法，基本思路是在每次分治中，对数组$A[l, r)$找到一个分治点$p$，**满足$p$左边的元素都小于等于$A[p]$，$p$右边的元素都大于$A[p]$**。
```
QuickSort(A, l, r):
     if l < r :
         p = Partition(A, l, r)
         QuickSort(A, l, p - 1)
         QuickSort(A, p + 1, r)
```
至于为什么这样分治能对数组排序我们不作分析，只介绍Partition的实现。为了理解Partition的过程并能写出BugFree的代码，
我们以[Lomuto](https://en.wikipedia.org/wiki/Quicksort)的实现为例，使用循环不变量描述算法过程。

> 注意：此实现下的快速排序性能并不是最优的。

首先，我们将数组划分为四个部分$A[l, i, j, r)$，每个部分和目标值$x$的大小关系不同，其中：
$$
\begin{cases}
&k \in [l, i), A[k] <= x \\\\ &k \in [i, j), A[k] > x \\\\ &k \in [j, r), A[k] \ ? \ x 
\end{cases}
$$
![quicksort_partition.png](/images/quicksort_partition.png)

初始条件为$i=j=l, x=A[r-1]$，此时$[l,i),[i,j)$都为空，循环不变量成立。

终结条件为$j=r$，此时$i$即为我们要寻找的分治点。

当$j=j+1$时，有两种情况：
1. $A[j]>x$，此时对 $k \in [i, j+1), A[k] > x$ ，不变量成立
2. $A[j]<=x$，此时交换$A[i], A[j]$，$i=i+1$。即把$A[j]$放到小于$x$的区间末尾，并将此区间右移一位，保持循环不变量成立
![quicksort_loop.png](/images/quicksort_loop.png)

按照四要素对Partition编码并实现快排。

> 注意：对下标的处理

```go
func swap(A []int, i, j int) {
	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}

func Partition(A []int, l, r int) (pivot int) {
	i := l
	j := l
	x := A[r-1]
	for j < r {
		if A[j] <= x {
			swap(A, i, j)
			i = i + 1
		}
		j = j + 1
	}
	return i - 1
}

func QSort(A []int, l, r int) {
	if l < r {
		pivot := Partition(A, l, r)
		QSort(A, l, pivot)
		QSort(A, pivot+1, r)
	}
}
```
> [算法导论](https://jingyuexing.github.io/Ebook/Algorithm/%E7%AE%97%E6%B3%95%E5%AF%BC%E8%AE%BA.pdf)
