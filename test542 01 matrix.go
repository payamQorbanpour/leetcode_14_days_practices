package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

	fmt.Println(updateMatrix(mat))
}

func updateMatrix(mat [][]int) [3][3]int {
	m, n := len(mat), len(mat[0])
	var resultMat = [3][3]int{}

	if m == 0 {
		return resultMat
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				resultMat[i][j] = 0
			} else {
				if i > 0 {
					resultMat[i][j] = min(resultMat[i][j], resultMat[i-1][j]+1)
				}
				if j > 0 {
					resultMat[i][j] = min(resultMat[i][j], resultMat[i][j-1]+1)
				}
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				resultMat[i][j] = min(resultMat[i][j], resultMat[i+1][j]+1)
			}
			if j < n-1 {
				resultMat[i][j] = min(resultMat[i][j], resultMat[i][j+1]+1)
			}
		}
	}
	return resultMat
}

func min(m, n int) int {
	if m < n {
		return m
	}

	return n
}
