package AdventOfCode

import "fmt"

var pointPower map[Point10]int

func Main11() {
	coord := Point10{1, 1}
	pointPower = make(map[Point10]int)
	maxPower := 0
	sno := 9005
	maxSize := 1
	sizePower := _3x3Power(Point10{1, 1}, sno, 1)
	powerX := sizePower
	powerY := sizePower
	for size := 1; size <= 100; size++ {
		sizePower = IncrementPowerXY(Point10{1, 1}, sno, size, sizePower)
		for x:=1; x<=300-size+1; x++ {
			if x == 1 {
				powerX = sizePower
			} else {
				powerX = IncrementPowerX(Point10{x, 1}, sno, size, powerX)
			}
			for y:=1; y<=300-size+1; y++ {
				if y == 1 {
					powerY = powerX
				} else {
					powerY = IncrementPowerY(Point10{x, y}, sno, size, powerY)
				}
				tempCoord := Point10{x, y}
				if powerY > maxPower {
					maxPower = powerY
					coord = tempCoord
					maxSize = size
				}
			}
		}
	}
	fmt.Print(coord, maxSize, maxPower)
}

func IncrementPowerXY(p Point10, sno int, size int, prevPower int) int {
	if size == 1 || prevPower == -33 {
		return _3x3Power(p, sno, size)
	}
	power := prevPower
	for x := p.X; x < p.X+size; x++ {
		power += PowerLevel(Point10{x, p.Y+size-1}, sno)
	}
	for y := p.Y; y < p.Y+size-1; y++ {
		power += PowerLevel(Point10{p.X+size-1, y}, sno)
	}
	return power
}

func IncrementPowerX(p Point10, sno int, size int, prevPower int) int {
	power := prevPower

	for y := p.Y; y < p.Y+size; y++ {
		power -= PowerLevel(Point10{p.X-1, y}, sno)
		power += PowerLevel(Point10{p.X+size-1, y}, sno)
	}
	return power
}

func IncrementPowerY(p Point10, sno int, size int, prevPower int) int {
	power := prevPower
	for x := p.X; x < p.X+size; x++ {
		power -= PowerLevel(Point10{x, p.Y-1}, sno)
		power += PowerLevel(Point10{x, p.Y+size-1}, sno)
	}
	return power
}

func _3x3Power(p Point10, sno int, size int) int {
	power := 0
	for x := p.X; x < p.X+size; x++ {
		for y := p.Y; y < p.Y+size; y++ {
			power += PowerLevel(Point10{x, y}, sno)
		}
	}
	return power
}

func PowerLevel(p Point10, sno int) int {
	if pointPower[p] != 0 {
		return pointPower[p]
	}
	rackId := p.X + 10
	power := rackId * p.Y
	power += sno
	power *= rackId
	power = (power%1000) / 100
	power -= 5
	pointPower[p] = power
	return power
}
