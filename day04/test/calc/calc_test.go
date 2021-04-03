package calc  // 单元测试 go test ./calc/  

import "testing"

func TestAdd01(t *testing.T) {
	if 3 != Add(1, 2) {
		t.Error("1 + 2 != 3")
	}
}