package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2

	fmt.Println(floodFill(image, sr, sc, newColor))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	node := image[sr][sc]

	if node == newColor {
		return image
	}

	m, n := len(image), len(image[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= m || c >= n {
			return
		}

		if image[r][c] != node {
			return
		}

		image[r][c] = newColor
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(sr, sc)

	return image
}
