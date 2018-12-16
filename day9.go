package AdventOfCode

import "fmt"

func Main9() {
	fmt.Println(HighScore(468, 71843*100))
}

type Link struct {
	Val int
	NextLink *Link
	PrevLink *Link
}

func HighScore(numPlayers int, numMarbles int) int {
	players := make([]int, numPlayers+1)
	CurrentLink := &Link{Val: 0}
	CurrentLink.NextLink = CurrentLink
	CurrentLink.PrevLink = CurrentLink
	currentPlayer := 0
	for i := 1; i <= numMarbles; i++ {
		currentPlayer  = NextPlayer(currentPlayer, numPlayers)
		if i % 23 == 0 {
			players[currentPlayer] += i
			for j := 0; j < 7; j ++ {
				CurrentLink = CurrentLink.PrevLink
			}
			_TempLink, removed := RemoveLink(CurrentLink)
			CurrentLink = _TempLink
			players[currentPlayer] += removed
		} else {
			CurrentLink = InsertLink(CurrentLink, i)
		}
	}
	return MaxScore(players)
}

func RemoveLink(link *Link) (*Link, int) {
	removed := link.Val
	link.PrevLink.NextLink = link.NextLink
	link.NextLink.PrevLink = link.PrevLink
	return link.NextLink ,removed
}

func MaxScore(players []int) int {
	max := 0
	for _, score := range players {
		max = GetMax(max, score)
	}
	return max
}

func InsertLink(link *Link, val int) *Link {
	NewLink := &Link{Val: val, NextLink: link.NextLink.NextLink, PrevLink: link.NextLink}
	link.NextLink.NextLink.PrevLink = NewLink
	link.NextLink.NextLink = NewLink
	return link.NextLink.NextLink
}

func NextPlayer(currentPlayer, maxPlayers int) int {
	ans := (currentPlayer+1) % (maxPlayers + 1)
	if ans == 0 {
		return 1
	}
	return ans
}