package log

import (
	"testing"
)

func TestGetInstace(t *testing.T) {
	logger := GetInstance("test")
	logger.Info(
		"Hello World",
		"type", "greeting",
		"target", "planet",
	)
}