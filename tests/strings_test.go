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

func TestMulti(t *testing.T) {
	t.Run("Group 1", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("Parallel simple")
			var x, y, result = 2, 2, 4
			realResult := strings.Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("Parallel medium")
			var x, y, result = 222, 222, 49284
			realResult := strings.Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("Parallel negative")
			var x, y, result = -2, -2, 4
			realResult := strings.Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})
}
