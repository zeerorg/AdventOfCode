package AdventOfCode

import (
	"fmt"
	"strconv"
)

func Main14() {
	FirstElf := &Link{3, nil, nil}
	SecondElf := &Link{7, FirstElf, FirstElf}
	StartRecipe := FirstElf
	StartRecipe.NextLink = SecondElf
	StartRecipe.PrevLink = SecondElf
	EndRecipe := SecondElf

	//times := 110220
	nextLink := 0
	//num := 110201
	//sb := strings.Builder{}

	nextLink = FirstElf.Val + SecondElf.Val
	firstAdd := nextLink/10
	secondLink := nextLink%10
	if firstAdd > 0 {
		EndRecipe = AddRecipe(EndRecipe, firstAdd)
	}
	EndRecipe = AddRecipe(EndRecipe, secondLink)
	links := 4
	toCompare := "110201"

	for !DoesCome(EndRecipe, toCompare) {
		nextLink = FirstElf.Val + SecondElf.Val
		firstAdd := nextLink/10
		secondLink := nextLink%10
		if firstAdd > 0 {
			EndRecipe = AddRecipe(EndRecipe, firstAdd)
			links += 1
		}
		if DoesCome(EndRecipe, toCompare) {
			break
		}
		EndRecipe = AddRecipe(EndRecipe, secondLink)
		links += 1
		runTill := FirstElf.Val
		for j:=0; j<=runTill; j++ {
			FirstElf = FirstElf.NextLink
		}
		runTill = SecondElf.Val
		for j:=0; j<=runTill; j++ {
			SecondElf = SecondElf.NextLink
		}
		//PrintRecipes(StartRecipe, EndRecipe)
		_ = StartRecipe
	}
	fmt.Println(links-len(toCompare))
	//for i := 0; i<num; i++ {
	//	StartRecipe = StartRecipe.NextLink
	//}
	//
	//for i:=0; i<10; i++ {
	//	sb.WriteString(strconv.Itoa(StartRecipe.Val))
	//	StartRecipe = StartRecipe.NextLink
	//}
	//fmt.Println(sb.String())
}

func DoesCome(link *Link, data string) bool {
	end := link
	for i:=0; i < len(data); i++ {
		if strconv.Itoa(end.Val) != string(data[len(data)-1-i]) {
			return false
		}
		end = end.PrevLink
	}
	return true
}

func PrintRecipes(recipe *Link, end *Link) {
	start := recipe
	for start != end {
		fmt.Print(start.Val)
	}
	fmt.Println()
}

func AddRecipe(toAdd *Link, val int) *Link {
	NewRecipe := &Link{val, toAdd.NextLink, toAdd}
	toAdd.NextLink.PrevLink = NewRecipe
	toAdd.NextLink = NewRecipe
	toAdd = NewRecipe
	return toAdd
}