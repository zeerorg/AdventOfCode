package AdventOfCode

import "testing"

var (
	first = Rect{Point{1, 3}, Point{4, 6}, "#1"}
	second = Rect{Point{3, 1}, Point{6, 4}, "#1"}
	third = Rect{Point{5, 5}, Point{6, 6}, "#1"}
	fourth = Rect{Point{1, 3}, Point{4, 4}, "#1"}
)

func TestGetRect(t *testing.T) {
	t.Run("Creating a sample Rectangle", func(t *testing.T) {
		r := Rect{Point{1,3}, Point{4, 6}, "#1"}
		if GetRect("#1 @ 1,3: 4x4") != r {
			t.Log(r)
			t.Log(GetRect("#1 @ 1,3: 4x4"))
			t.Fatal("Not working")
		}
	})
}

func TestDoOverlap(t *testing.T) {
	if DoOverlap(second, third) {
		t.Fatal("You are wrong")
	}
	if !DoOverlap(first, second) {
		t.Fatal("Again wrong")
	}
	if DoOverlap(second, third) {
		t.Fatal("Enough of this")
	}
	if !DoOverlap(first, fourth) {
		t.Fatal("Very Wrong")
	}
}

func TestIntersectingArea(t *testing.T) {
	if IntersectingArea(first, fourth) != 3 {
		t.Fatal("Wrong Algo")
	}
}

func TestIntersectingAreaArray(t *testing.T) {
	var rects [2]Rect
	rects[0] = first
	rects[1] = second

	if IntersectingAreaArray(rects[:]) != 4 {
		t.Log(IntersectingAreaArray(rects[:]))
		t.Fatal("Not working")
	}
}

func GetSlice(points <-chan Point) []Point {
	var ans []Point
	for point := range points {
		ans = append(ans, point)
	}
	return ans[:]
}