---
title: '理解动态规划：从切钢管开始'
date: '2023-11-13'
description: ''
---

动态规划(dynamic programming)是一种分治算法，可以解决一类拥有特定性质的问题，针对这些问题，动态规划提供一种高效的解决思路。特定性质包括两点：

1. 重复子问题
2. 最优子结构

两种性质按字面意义理解都不难，1，有被重复处理的子问题；2，当父问题的解是最优时，子问题也是最优，具体问题中的难点是发现问题具有这两种性质，并做出子问题的定义。
本文以切钢管问题为例子，介绍动态规划的问题解决原理和编码模板。

钢管问题：已知长度为$i$的钢管的的售价为$p(i)$，现在可将长度为$n$的钢管切成任意段，求所有片段的售价总和最大值。

直观的看，对于长度为$n$的钢管，我们可以分别尝试在长度$1,2,...,n$的位置“砍一刀”得到一个切分方案，得到一个递归式。
$$
P(n)=Max(p(i) + P(n-i)), i \in [1,n]
$$

画出$P(4)$的递归树，树中有多个重复的子节点$P(1),P(2),P(0)$，说明问题满足动态规划的性质之一：**重复子问题**。

![dp_cutting_recursive_tree.png](/images/dp_cutting_recursive_tree.png)

切钢管问题中，每次切分后只会产生一个不重叠的子问题，可以通过**cut-paste**方法进行反证最优子结构。对于一般
的动态规划问题，也可以通过这种方式进行证明。

假设$P(n)=p(i)+P(n-i)$为最优解， 而$P(n-i)$不是最优解，那么会存在另外一个$P1(n-i)$，使得$P1(n)=p(i)+P1(n-i)
$为最优解，与$P(n)$为最优解矛盾。

确定问题的递归式和最优子结构性质后，就可以自底向上递推得出最优解。

```go
func CuttingProblem(price []int, length int) int {
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = max(dp[i], price[j]+dp[i-j])
		}
	}

	return dp[length]
}
```

另一个问题是**最优解路径**。由于具有**最优子结构**性质，每个问题的最优解都是由当前问题的解和子问题最优解组成，因此我们
可以**从后向前回溯子问题**，将所有解拼接起来得到最优解路径。
```go
func Path(price, cuttingPoint []int, length int) string {
	if length > 0 {
		return fmt.Sprintf("%s+%d[%d]", Path(price, cuttingPoint, length-cuttingPoint[length]), cuttingPoint[length], price[cuttingPoint[length]])
	}
	return ""
}

func CuttingProblem(price []int, length int) int {
	dp := make([]int, length+1)
	cuttingPoint := make([]int, length+1)
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			if price[j]+dp[i-j] > dp[i] {
				dp[i] = price[j] + dp[i-j]
				cuttingPoint[i] = j
			}
		}
	}

	fmt.Println(Path(price, cuttingPoint, length))

	return dp[length]
}
```