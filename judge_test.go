package main

import "testing"

func TestJudge(t *testing.T) {
	b := [3][3]int{
		{1, 1, 1},
		{4, 0, 4},
		{1, 1, 4},
	}
	str := judge(b)
	if str != "O win" {
		t.Errorf("Test01 is failed")
	}
}
