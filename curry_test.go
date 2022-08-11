package curry

import (
	"math"
	"testing"

	"github.com/vengor/curry/curryone"
)

func TestCurry(t *testing.T) {
	t.Run("curryone.FromTwoArgsOneRet", func(t *testing.T) {
		f := curryone.FromTwoArgsOneRet(1, math.Min)
		if v := f(0.0); v != 0.0 {
			t.Fatalf("Expected float 0.0, got %f", v)
		}
	})
}
