package dp

var PATH_COST = [4][4]int{
	{0, 2, 5, 18},
	{-1, 0, 5, 15},
	{-1, -1, 0, 6},
	{-1, -1, -1, 0},
}

func MinCost(src, dst int, memo *[][]int) int {
	min := PATH_COST[src][dst]
	if src == dst || src == dst-1 {

		return min
	}

	if ((*memo)[src][dst]) != 0 {
		return (*memo)[src][dst]
	}

	for i := src + 1; i < dst; i++ {
		minPath := MinCost(src, i, memo) + MinCost(i, dst, memo)

		if minPath < min {
			min = minPath
		}
		(*memo)[src][dst] = minPath
	}
	return min
}

func MinCostDP() int {
	stations := len(PATH_COST[0])
	memo := make([]int, stations)
	memo[0] = 0
	memo[1] = PATH_COST[0][1]

	for i := 2; i < stations; i++ {
		memo[i] = PATH_COST[0][i]
		for j := 1; j < i; j++ {
			cost := memo[j] + PATH_COST[j][i]
			if memo[i] > cost {
				memo[i] = cost
			}
		}
	}
	return memo[stations-1]
}
