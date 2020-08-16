package min_time_to_visit_all_points

import "math"

func minTimeToVisitsAllPoints(points [][]int) int {
	result := 0
	upperbound := len(points)
	for idx := 1; idx < upperbound; idx++ {
		result += calculateDistance(points[idx-1], points[idx])
	}
	return result
}
func calculateDistance(p1 []int, p2 []int) int {
	xdiff := int(math.Abs(float64(p1[0] - p2[0])))
	ydiff := int(math.Abs(float64(p1[1] - p2[1])))
	return max(xdiff, ydiff)
}
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
