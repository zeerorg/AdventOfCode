package AdventOfCode

import (
	"fmt"
	"strings"
)

type Rule struct {
	Match []bool
	Val bool
}

func Main12() {
	currentState := GetCurrentState(Day12Init, 5000, 2500)
	rules := GetCurrentRules(Day12Rules)
	//PrintState(currentState, 150)
	for i := 1; i <= 1000; i++ {
		currentState = NextState(currentState, rules)
		PrintState(currentState, 2500)
		//fmt.Println()
	}
}

func PrintState(state []bool, middle int) {
	//flag := false
	pots := 0
	for i := 0; i < len(state); i ++ {
		if state[i] {
			//fmt.Print("#")
			pots += i - middle - 1
		} else {
			//fmt.Print(".")
		}
	}
	fmt.Println(pots)
}

func NextState(state []bool, rules []Rule) []bool {
	nextState := make([]bool, len(state))
	for i:=2; i<=len(state)-3; i++ {
		nextState[i] = false
		for _, rule := range rules {
			if MatchRule(state, i, rule) {
				nextState[i] = rule.Val
				break
			}
		}
	}
	return nextState
}

func MatchRule(state []bool, pos int, rule Rule) bool {
	for i := -2; i<3; i++ {
		if state[pos+i] != rule.Match[i+2] {
			return false
		}
	}
	return true
}

func GetCurrentRules(data string) []Rule {
	_tempRules := strings.Split(data, "\n")
	rules := make([]Rule, 0)
	for _, _tempRule := range _tempRules {
		_tempMatch := strings.Split(_tempRule, " => ")[0]
		next := strings.Split(_tempRule, " => ")[1] == "#"
		match := make([]bool, 0)
		for _, chr := range _tempMatch {
			match = append(match, chr == rune('#'))
		}
		rules = append(rules, Rule{match, next})
	}
	return rules
}

func GetCurrentState(data string, max int, middle int) []bool {
	currentState := make([]bool, max)
	i := middle
	for _, chr := range data {
		i += 1
		currentState[i] = chr == rune('#')
	}
	return currentState
}

const Test12Init = `#..#.#..##......###...###`
const Test12Rules = `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
const Day12Init = `######....##.###.#..#####...#.#.....#..#.#.##......###.#..##..#..##..#.##..#####.#.......#.....##..`
const Day12Rules = `...## => #
###.. => .
#.#.# => .
##### => .
....# => .
##.## => .
##.#. => #
##... => #
#..#. => #
#.#.. => .
#.##. => .
..... => .
##..# => .
#..## => .
.##.# => #
..### => #
..#.# => #
.#### => #
.##.. => .
.#..# => #
..##. => .
#.... => .
#...# => .
.###. => .
..#.. => .
####. => #
.#.## => .
###.# => .
#.### => #
.#... => #
.#.#. => .
...#. => .`
