---
title: '容易被忽略的计数排序稳定性'
date: '2023-11-14'
description: '尽管二分查找的基本思想相对简单，但细节可以令人难以招架...'
---
计数排序中，一个容易被忽略的点是保持原数组元素的相对顺序，即排序算法的稳定性。在下面的实现中，数组中每个元素被计数后，被进行无序的还原。
```go
func CountingSortNaive(nums []int, N int) {
	counting := make([]int, N)
	for _, num := range nums {
		counting[num] = counting[num] + 1
	}
	k := 0
	for i, count := range counting {
		for count > 0 {
			nums[k] = i
			k = k + 1
			count = count - 1
		}
	}
}
```
为了保持稳定性，我们在排序前对计数数组做一个**滚动累加**，以计算出每个值在原数组中对应的下标，最后对原数组从后向前做排序。
```go
func CountingSort(nums []int, N int) {
	counting := make([]int, N)
	tmp := make([]int, len(nums))
	for i := range nums {
		counting[nums[i]] = counting[nums[i]] + 1
	}
	for i := range counting {
		counting[i] = counting[i] + counting[i-1]
	}
	for i := len(nums) - 1; i >= 0; i = i - 1 {
		v := nums[i]
		tmp[counting[v]-1] = v
		counting[v] = counting[v] - 1
	}
	for i := 0; i < len(tmp); i = i + 1 {
		nums[i] = tmp[i]
	}
}
```
稳定的计数排序能帮助处理基数排序的问题。

> [算法导论](https://jingyuexing.github.io/Ebook/Algorithm/%E7%AE%97%E6%B3%95%E5%AF%BC%E8%AE%BA.pdf)
