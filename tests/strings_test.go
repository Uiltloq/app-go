package tests

import (
	"testing"

	"github.com/Uiltloq/app-go/pkg/strings"
)

func TestMultiple(t *testing.T) {
	var x, y, result = 2, 2, 4

	realResult := strings.Multiple(x, y)

	if realResult != result {
		t.Errorf("%d != %d", result, realResult)
	}
}
