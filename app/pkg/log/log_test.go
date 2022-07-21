package log

import (
	"testing"
	"fmt"
)

func TestGetInstace(t *testing.T) {
	err := fmt.Errorf("this one!!")
	wrap := fmt.Errorf("Wrapped: %w", err)
	logger := GetInstance("test")
	logger.Error(wrap, "logger Message", nil)
}