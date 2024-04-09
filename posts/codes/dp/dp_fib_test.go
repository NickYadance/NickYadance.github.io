package dp

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

提示：

1 <= n <= 45
*/
func climbStairs(n int) int {
	fib0 := 0
	fib1 := 1
	fib := 0
	if n == 0 {
		return fib0
	}
	if n == 1 {
		return fib1
	}
	for i := 2; i <= n; i++ {
		fib = fib0 + fib1
		fib0 = fib1
		fib1 = fib
	}
	return fib
}

/*
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n) 。

示例 1：

输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1
示例 2：

输入：n = 3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2
示例 3：

输入：n = 4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3

提示：

0 <= n <= 30
*/
func fib(n int) int {
	fib0 := 0
	fib1 := 1
	fib := 0
	if n == 0 {
		return fib0
	}
	if n == 1 {
		return fib1
	}
	for i := 2; i <= n; i++ {
		fib = fib0 + fib1
		fib0 = fib1
		fib1 = fib
	}
	return fib
}

/*
泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

示例 1：

输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
示例 2：

输入：n = 25
输出：1389537

提示：

0 <= n <= 37
答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
*/
func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1
	tn := 0
	if n == 0 {
		return t0
	}
	if n == 1 {
		return t1
	}
	if n == 2 {
		return t2
	}
	for i := 3; i <= n; i++ {
		tn = t0 + t1 + t2
		t0 = t1
		t1 = t2
		t2 = tn
	}
	return tn
}

/*
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

示例 1：

输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。
示例 2：

输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。

提示：

2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/
func minCostClimbingStairs(cost []int) int {
	/*
		m[i] = Min(m[i-1] + cost[i-1], m[i-2] + cost[i-2])
		m[0] = 0
		m[1] = cost[0]
		ans = m[len(cost)]
	*/
	m0 := 0
	m1 := 0
	m := 0
	for i := 2; i <= len(cost); i++ {
		m = Min(m0+cost[i-2], m1+cost[i-1])
		m0 = m1
		m1 = m
	}
	return m
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。

	偷窃到的最高金额 = 2 + 9 + 1 = 12 。

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/
func rob(nums []int) int {
	/*
		dp[i] = Max(dp[i-2] + nums[i], dp[i-1])
		dp[0] = nums[0]
		dp[1] = Max(nums[0], nums[1])

		max = Max(m + nums[i], n)
		m = n
		n = max
	*/
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	m := nums[0]
	n := Max(nums[0], nums[1])
	max := 0
	for i := 2; i < len(nums); i++ {
		max = Max(m+nums[i], n)
		m = n
		n = max
	}
	return max
}

/*
给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

示例 1：

输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。
示例 2：

输入：nums = [2,2,3,3,3,4]
输出：9
解释：
删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。

提示：

1 <= nums.length <= 2 * 104
1 <= nums[i] <= 104
*/
func deleteAndEarn(nums []int) int {
	/**
	dp[i] = Max(dp[i-2] + sum(i), dp[i-1])
	*/
	sum := make([]int, 10001)
	for _, n := range nums {
		sum[n] += n
	}
	m := sum[0]
	n := Max(sum[0], sum[1])
	max := 0
	for i := 2; i < len(sum); i++ {
		max = Max(m+sum[i], n)
		m = n
		n = max
	}
	return max
}
