package common_problem_types

import common_types "aoc.2023/lib/common/types"

type PlatformParameters struct {
	TotalRows      int
	TotalLoads     int
	RocksStack     common_types.Stack[[2]int]
	BlockPointsMap map[int][]int
}
