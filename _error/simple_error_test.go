// @Author:冯铁城 [17615007230@163.com] 2023-04-13 14:54:07
package _error

import "testing"

func TestTryUnitTest(t *testing.T) {

	//1.value > 0 情况用例
	if TryUnitTest(1) {
		t.Log("the value 1's res should be false")
		t.Fail()
	}

	//2.value = 0 情况用例
	if TryUnitTest(0) {
		t.Log("the value 0's res should be false")
		t.Fail()
	}

	//3.value < 0 情况用例
	if !TryUnitTest(-1) {
		t.Log("the value -1's res should be true")
		t.Fail()
	}
}
