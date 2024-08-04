package reflection

import (
	"testing"
)

func TestIsReflectionWithOneException(t *testing.T) {
	values := []int{1, 0, 1, 1, 0, 0, 1}
	if !IsReflectionWithOneException(2, values) {
		t.Errorf("error should be reflection")
	}
	if IsReflectionWithOneException(3, values) {
		t.Errorf("should not be reflection")
	}
}
