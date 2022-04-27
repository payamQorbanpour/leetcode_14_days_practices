package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

	fmt.Println(updateMatrix(mat))
}

func updateMatrix(mat [][]int) [3][3]int {
	var resultMatrix = [3][3]int{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			result := 0

			if mat[i][j] == 1 {
				result = min(result, dfs(i, j, mat))
				resultMatrix[i][j] = result
			}
		}
	}

	return resultMatrix
}

func dfs(i, j int, mat [][]int) int {
	m, n := len(mat), len(mat[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}

	if mat[i][j] == 0 {
		return 1
	}

	mat[i][j] = 0
	x := 1 + dfs(i-1, j, mat) + dfs(i+1, j, mat) + dfs(i, j+1, mat) + dfs(i, j-1, mat)
	return x
}

func min(m, n int) int {
	if m < n {
		return m
	}

	return n
}
