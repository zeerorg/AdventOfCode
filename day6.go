package AdventOfCode

import (
	"fmt"
	"strconv"
	"strings"
)

type Map [][]int

func Main6() {
	points := ToPoints(strings.Split(Day6Data, "\n"))
	maxHamming := 0
	maxX := 0
	maxY := 0
	for i, point := range points {
		for _, secP := range points[i+1:] {
			maxHamming = GetMax(maxHamming, GetDistance(point, secP))
		}
		maxX = GetMax(maxX, point.x)
		maxY = GetMax(maxY, point.y)
	}
	fmt.Println(points)

	Land := make([][]int, maxX+maxHamming)
	for x := 0; x < maxX+maxHamming; x += 1 {
		Land[x] = make([]int, maxY+maxHamming)
		for y := 0; y < maxY+maxHamming; y += 1 {
			Land[x][y] = GetLowestHammingInd(Point{x, y}, points)
		}
	}

	isInfinite := make(map[int]bool)
	for x := 0; x < maxX+maxHamming; x += 1 {
		for y := 0; y < maxY+maxHamming; y += 1 {
			isInfinite[Land[0][y]] = true
			isInfinite[Land[x][0]] = true
			isInfinite[Land[maxX+maxHamming-1][y]] = true
			isInfinite[Land[x][maxY+maxHamming-1]] = true
		}
	}

	areas := make([]int, len(points)+1)
	for x := 0; x < maxX+maxHamming; x += 1 {
		for y := 0; y < maxY+maxHamming; y += 1 {
			areas[Land[x][y]] += 1
		}
	}

	maxArea := 0
	for i, area := range areas {
		if i > 0 && !isInfinite[i] && maxArea < area {
			maxArea = area
		}
	}

	fmt.Println(maxArea)

	ans := 0
	for x := 0; x < maxX+maxHamming; x += 1 {
		for y := 0; y < maxY+maxHamming; y += 1 {
			if GetTotalDistance(Point{x, y}, points) < 10000 {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}

func GetLowestHammingInd(p Point, points []Point) int {
	minDist := 10000
	for _, point := range points {
		if GetDistance(point, p) < minDist {
			minDist = GetDistance(point, p)
		}
	}

	var ind []int
	for i, point := range points {
		if GetDistance(point, p) == minDist {
			ind = append(ind, i)
		}
	}
	if len(ind) > 1 {
		return 0
	}

	return ind[0] + 1
}

func GetTotalDistance(p Point, points []Point) int {
	ans := 0
	for _, point := range points {
		ans += GetDistance(p, point)
	}
	return ans
}

func ToPoints(s []string) []Point {
	var p []Point
	for _, str := range s {
		x, _ := strconv.Atoi(strings.Split(str, ",")[0])
		y, _ := strconv.Atoi(strings.Split(str, ",")[1][1:])
		p = append(p, Point{x, y})
	}
	return p
}

func GetDistance(p, q Point) int {
	return GetAbs(p.x-q.x) + GetAbs(p.y-q.y)
}

func GetAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const Day6Data = `242, 112
292, 356
66, 265
73, 357
357, 67
44, 303
262, 72
220, 349
331, 301
338, 348
189, 287
285, 288
324, 143
169, 282
114, 166
111, 150
251, 107
176, 196
254, 287
146, 177
149, 213
342, 275
158, 279
327, 325
201, 70
145, 344
227, 345
168, 261
108, 236
306, 222
174, 289
67, 317
316, 302
248, 194
67, 162
232, 357
300, 193
229, 125
326, 234
252, 343
51, 263
348, 234
136, 337
146, 82
334, 62
255, 152
326, 272
114, 168
292, 311
202, 62`
