package dp

import (
	"math"
)

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？



示例 1：


输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6


提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109
*/

func uniquePaths(m int, n int) int {
	/*
		dp[m][n] = dp[m-1][n] + dp[m][n-1]
		dp[0][0] = 0
		dp[i][0] = 1
		dp[0][i] = 1
	*/
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsRollingArray(m int, n int) int {
	/*
		dp[m][n] = dp[m-1][n] + dp[m][n-1]
		aka
		f[n] = f[n] + f[n-1]
		 -------->
		|    ^
		|    |
		| <- x
		|
	*/
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[j] += f[j-1]
		}
	}
	return f[n-1]
}

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 200
*/

func minPathSum(grid [][]int) int {
	/*
		dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		dp[0][i] = sum(grid[0][i])
		dp[i][0] = sum(grid[i][0])
	*/
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func minPathSumRollingArray(grid [][]int) int {
	/*
		dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		dp[0][i] = sum(grid[0][i])
		dp[i][0] = sum(grid[i][0])
	*/
	m := len(grid)
	n := len(grid[0])
	f := make([]int, n)
	f[0] = grid[0][0]
	for i := 1; i < len(f); i++ {
		f[i] = f[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		f[0] += grid[i][0]
		for j := 1; j < n; j++ {
			f[j] = min(f[j-1], f[j]) + grid[i][j]
		}
	}
	return f[n-1]
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

示例 1：

输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：

输入：obstacleGrid = [[0,1],[0,0]]
输出：1

提示：

m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*
		let G(i,j) = 1 - grid[i][j]
		dp[i][j] = G[i][j]*(dp[i-1][j] + dp[i][j-1]))
		dp[0][j] = G(0,j)
		dp[i][0] = G(i,0)
	*/
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([]int, n)
	G := func(i, j int) int {
		return 1 - obstacleGrid[i][j]
	}
	f[0] = G(0, 0)
	for i := 1; i < n; i++ {
		f[i] = G(0, i) * f[i-1]
	}
	for i := 1; i < m; i++ {
		f[0] = G(i, 0) * f[0]
		for j := 1; j < n; j++ {
			f[j] = G(i, j) * (f[j-1] + f[j])
		}
	}
	return f[n-1]
}

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：

	  2
	 3 4
	6 5 7

4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10

提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

进阶：

你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
*/
func minimumTotal(triangle [][]int) int {
	/*
		   2
		  3 4
		 6 5 7
		4 1 8 3

		dp[i][j] = t[i][j] + min(dp[i-1][j], dp[i-1][j-1])
	*/
	m := len(triangle)
	f := make([]int, m)
	for i := 0; i < m; i++ {
		f[i] = math.MaxInt
	}

	f[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		x := math.MaxInt
		for j := 0; j < len(triangle[i]); j++ {
			tmp := f[j]
			f[j] = triangle[i][j] + min(f[j], x)
			x = tmp
		}
	}

	ans := math.MaxInt
	for i := 0; i < m; i++ {
		ans = min(ans, f[i])
	}
	return ans
}

func minimumTotalReverse(triangle [][]int) int {
	/*
			   2
			  3 4
			 6 5 7
			4 1 8 3
		   0 0 0 0 0

		dp[i][j] = t[i][j] + min(dp[i+1][j], dp[i+1][j+1])
	*/
	m := len(triangle)
	dp := make([]int, m+1)
	for i := m - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j+1], dp[j])
		}
	}
	return dp[0]
}

/*
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

示例 1：

输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
输出：13
解释：如图所示，为和最小的两条下降路径
示例 2：

输入：matrix = [[-19,57],[-40,-5]]
输出：-59
解释：如图所示，为和最小的下降路径

提示：

n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100
*/
func minFallingPathSum(matrix [][]int) int {
	/*
		dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
		dp[i] = min(x + dp[i] + dp[i+1]) + matrix[i][j]
	*/
	n := len(matrix)
	dp := make([]int, n+1)
	dp[0] = math.MaxInt
	for i := 1; i <= n; i++ {
		x := dp[0]
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if j < n {
				dp[j] = min(x, dp[j], dp[j+1]) + matrix[i-1][j-1]
			} else {
				dp[j] = min(x, dp[j]) + matrix[i-1][j-1]
			}
			x = tmp
		}
	}
	ans := math.MaxInt
	for i := 0; i <= n; i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4
示例 2：

输入：matrix = [["0","1"],["1","0"]]
输出：1
示例 3：

输入：matrix = [["0"]]
输出：0

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
func maximalSquare(matrix [][]byte) int {
	/*
		let G(x) = x > 0 ? 1 : 0
		dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) * G(matrix[i][j]) + G(matrix[i][j])
				 = G(matrix[i][j])(1 + min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]))
	*/
	m := len(matrix)
	n := len(matrix[0])
	dp := make([]int, n+1)
	ans := 0
	G := func(i, j int) int {
		if matrix[i][j] == '1' {
			return 1
		} else {
			return 0
		}
	}

	for i := 1; i <= m; i++ {
		x := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			dp[j] = G(i-1, j-1) * (1 + min(x, dp[j-1], dp[j]))
			x = tmp

			if dp[j] > ans {
				ans = dp[j]
			}
		}
	}
	return ans * ans
}
