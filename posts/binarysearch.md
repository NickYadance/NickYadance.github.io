---
title: '彻底理解二分搜索的上下界问题'
date: '2023-11-10'
description: '尽管二分查找的基本思想相对简单，但细节可以令人难以招架...'
---
> 尽管二分查找的基本思想相对简单，但细节可以令人难以招架 ... — 高德纳

二分查找是一个迭代算法，基本思想是通过不断折半缩小查找范围，最终找到目标值。

如果只对二分查找的代码实现死记硬背，在编码过程中非常容易对上下界和终结条件发晕，
导致无法写出BugFree的代码。我们用数学的思想：**循环不变量**来描述整个迭代过程，避免发晕。循环不变量的原理类似多米诺骨牌，包括四个要素：
1. $Variant$，定义
2. $Start$，初始化
3. $Loop$，保持
4. $End$，终止

首先我们定义二分查找的循环不变量为：对于一个有序数组 $a[l,r)$ 和目标值 $n$，在**左闭右开的数组区间$[l, r)$内，区间左边的元素都比$n$小(或者不存在)，区间右边的元素都比$n$大(或者不存在)。**
$$ a[i]<n, i < l $$
$$ a[i]>=n, i >= r $$
算法开始时有$l=0,r=len(a)$，循环不变量成立。

当我们不断缩小这个区间直到$l=r$，算法终止。此时$l$左边的元素都小于$n$，$l$ 右边的元素(包括自身)都大于等于$n$，我们得到一个有用的数学性质，即$a[l]$是目标值n在数组$a$中的**下界**。
循环不变量定义的区间，实际上就是**目标值下界所在的区间**。

> 数组中N的下界是第一个大于或等于N的元素下标

因此，保持过程就是要不断寻找下一个下界所在的子区间。
1. 如果$a[mid]<n$，则n的下界在mid右边且不包括mid，区间$[mid + 1, r)$满足循环不变量
2. 如果$a[mid]>=n$，则n的下界在mid左边且包括mid，区间$[l, mid)$满足循环不变量

对照四要素，就可以写出**寻找目标值下界**的二分搜索算法。
```go
func BinarySearchLowerBound(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if arr[mid] < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
```

二分查找可以在二分搜索下界的基础上实现。
```go
func BinarySearch(arr []int, n int) int {
	lowerBound := BinarySearchLowerBound(arr, n)
	if lowerBound >= 0 && lowerBound < len(arr) && arr[lowerBound] == n {
		return lowerBound
	}
	return -1
}
```

也可以在迭代过程中提前对$a[mid]=n$做判断，减少迭代次数，但也丢失掉了算法的**稳定性**。

最后总结一下二分查找的循环不变量四要素。
1. $Variant$， 目标值下界所在的区间
2. $Start$， $l=0,r=len(a)$
3. $Loop$， 如果$a[mid]<n$，则n的下界在mid右边且不包括mid，反之在左边
4. $End$， $l=r$

> [二分查找有几种写法？它们的区别是什么？ - Jason Li的回答 - 知乎](https://www.zhihu.com/question/36132386/answer/530313852)
> 
> [算法导论](https://jingyuexing.github.io/Ebook/Algorithm/%E7%AE%97%E6%B3%95%E5%AF%BC%E8%AE%BA.pdf)