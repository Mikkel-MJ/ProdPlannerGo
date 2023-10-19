package utils

import (
	"fmt"
	"sort"
	"testing"
)

func TestLexorank(t *testing.T) {
	a := GenerateRankArray(25)
	fmt.Println(a)
	sort.Strings(a)
	fmt.Println(a)
	v, _ := Rank("aa", "cc")
	println(v)
}
