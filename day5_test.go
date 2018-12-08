package AdventOfCode

import (
	"testing"
)

func TestDestroyAdj(t *testing.T) {
	t.Parallel()
	//if DestroyAdj("sS") != "" {
	//	t.Fatal(DestroyAdj("sS"))
	//}
	if DestroyAdj("sSpQq") != "p" {
		t.Fatal(DestroyAdj("sSpQq"))
	}
	if DestroyAdj("sQa") != "sQa" {
		t.Fatal(DestroyAdj("sQa"))
	}
	if DestroyAdj("aa") != "aa" {
		t.Fatal(DestroyAdj("aa"))
	}
}

func TestRemoveUnit(t *testing.T) {
	t.Parallel()
	testStr := [3][3]string{
		{"abbc", "b", "ac"},
		{"bbcc", "c", "bb"},
		{"cCc", "c", ""},
	}
	for _, testcase := range testStr {
		if RemoveUnit(testcase[0], testcase[1]) != testcase[2] {
			t.Fatal(RemoveUnit(testcase[0], testcase[1]), testcase[0], testcase[1], testcase[2])
		}
	}

}
