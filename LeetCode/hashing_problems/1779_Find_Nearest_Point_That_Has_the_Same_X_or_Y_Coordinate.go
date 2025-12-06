package hashing_problems

import (
	"math"
)

func NearestValidPoint(x int, y int, points [][]int) int {
	result := 999999
	index := -1
	for i, point := range points {
		if point[0] == x || point[1] == y {

			x2 := point[0]
			y2 := point[1]

			distance := int(math.Abs(float64(x)-float64(x2)) + math.Abs(float64(y)-float64(y2)))

			if distance < result {
				result = distance
				index = i
			}
		}
	}

	return index
}
