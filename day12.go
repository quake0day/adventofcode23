package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Parses the input into springs' conditions and group sizes.
func parseInput(row string) (string, []int) {
	parts := strings.Split(row, " ")
	conditions := parts[0]
	groupsStr := strings.Split(parts[1], ",")

	var groups []int
	for _, g := range groupsStr {
		size, _ := strconv.Atoi(g)
		groups = append(groups, size)
	}

	fmt.Printf("Parsed Input - Conditions: [%s], Groups: %v\n", conditions, groups)
	return conditions, groups
}

// Determine springs' conditions that can be definitively fixed based on group sizes.
func determineFixedSprings(conditions string, groups []int) string {
	// Implementation of logic to fix certain springs based on group sizes
	// This is a simplified version and may need enhancements based on detailed requirements
	for _, groupSize := range groups {
		index := strings.Index(conditions, strings.Repeat("?", groupSize))
		if index != -1 {
			conditions = conditions[:index] + strings.Repeat("#", groupSize) + conditions[index+groupSize:]
		}
	}
	fmt.Printf("After Determining Fixed Springs - Conditions: %s\n", conditions)
	return conditions
}

func countArrangementsForRow(conditions string, groups []int) int {
	totalSprings := 0
	for _, g := range groups {
		totalSprings += g
	}

	// 已存在的温泉数量
	existingSprings := strings.Count(conditions, "#")
	fmt.Printf("existingSprings: %d\n", existingSprings)
	unknowns := strings.Count(conditions, "?")
	fmt.Printf("unknowns: %d\n", unknowns)

	extraSprings := totalSprings - existingSprings
	fmt.Printf("extraSprings: %d\n", extraSprings)

	// 初始化 dp 数组
	dp := make([][][]int, unknowns+1)
	for i := range dp {
		dp[i] = make([][]int, extraSprings+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, len(groups)+1) // 根据温泉组的数量调整第三维的大小
		}
	}

	dp[0][0][0] = 1 // 从第一个温泉组开始

	// 动态规划状态转移
	for i := 1; i <= unknowns; i++ {
		for j := 0; j <= extraSprings; j++ {
			for k := 0; k < len(groups); k++ {
				// 如果当前位置放置温泉
				if j+1 <= extraSprings && (k == 0 || dp[i-1][j][k-1] > 0) {
					dp[i][j+1][k+1] += dp[i-1][j][k] // 注意更新状态为 k+1
				}
				// 如果当前位置放置空地
				if k > 0 {
					dp[i][j][k-1] += dp[i-1][j][k] // 考虑转移到前一个温泉组
				}
			}
		}
	}

	// 输出用于调试的结果
	fmt.Printf("DP table for '%s':\n", conditions)
	for i := 0; i <= unknowns; i++ {
		for j := 0; j <= extraSprings; j++ {
			for k := 0; k <= len(groups); k++ {
				fmt.Printf("dp[%d][%d][%d] = %d\n", i, j, k, dp[i][j][k])
			}
		}
	}
	return dp[unknowns][extraSprings][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Place a group of broken springs at the given position.
func placeGroup(conditions string, pos int, groupSize int) string {
	return conditions[:pos] + strings.Repeat("#", groupSize) + conditions[pos+groupSize:]
}

// Handles special cases like rows with no `?` or all `?`.
func handleEdgeCases(conditions string, groups []int) (bool, int) {
	// Implement logic for edge cases
	fmt.Printf("Handling Edge Cases - Conditions: %s\n", conditions)
	return false, 0 // Placeholder return values
}

// Sums up the arrangements from all rows.
func sumArrangements(rows []string) int {
	total := 0
	for _, row := range rows {
		conditions, groups := parseInput(row)

		total += countArrangementsForRow(conditions, groups)
	}
	return total
}

func main() {
	rows := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	fmt.Printf("Total Arrangements: %d\n", sumArrangements(rows))
}
