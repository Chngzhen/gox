package arrays

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	arr := []interface{}{"1", "2", 2, 3}
	fmt.Println(Contains(arr, 2))
}

func TestContainsString(t *testing.T) {
	arr := []string{"1", "2", "3"}
	fmt.Println(ContainsString(arr, "2"))
}

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
