package leetcode

// 62. 不同路径
// https://leetcode-cn.com/problems/unique-paths/
func UniquePaths(m int, n int) int {
	temp := make([][]int, m)
	for i := 0; i < m; i++ {
		//初始化每一行
		temp[i] = make([]int, n)
		//设置每一行的第一个元素是1
		temp[i][0] = 1
	}
	//设置每一列的第一个元素是1
	for i := 0; i < n; i++ {
		temp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			temp[i][j] = temp[i - 1][j] + temp[i][j - 1]
		}
	}
	return temp[m - 1][n - 1]
}