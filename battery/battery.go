package battery

import (
	"fmt"
	"sort"
)

func MaxRunTime(n int, batteries []int) int64 {
	times := distribuite(0, n, batteries)
	return int64(times)
}

func distribuite(times, machines int, batteries []int) int {
	if len(batteries) < machines {
		return times
	}

	if contains(batteries[len(batteries)-machines:], 0) {
		return times
	}

	batteries = consumeBattery(machines, batteries)

	return distribuite(times+1, machines, batteries)
}

func consumeBattery(machines int, batteries []int) []int {
	sort.Ints(batteries)

	for idx := range batteries[len(batteries)-machines:] {
		batteries[idx+len(batteries)-machines]--
	}
	fmt.Println(batteries)
	return batteries
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
