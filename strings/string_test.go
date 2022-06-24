package xstrings

import (
	"testing"
)

func TestArrayToDelimitedString(t *testing.T) {
	arr := []string{"1", "2", "3"}
	xcp := "'1', '2', '3'"
	rst := ArrayToDelimitedString(arr, "'", "'", ", ")
	if xcp != rst {
		t.Error("期望结果：;", xcp, "实际结果：", rst)
	}
}

func TestArrayToCommaDelimitedString(t *testing.T) {
	arr := []string{"1", "2", "3"}
	xcp := "1,2,3"
	rst := ArrayToCommaDelimitedString(arr)
	if xcp != rst {
		t.Error("期望结果：;", xcp, "实际结果：", rst)
	}
}
