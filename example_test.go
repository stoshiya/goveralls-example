package goveralls_example

import (
	"math/rand"
	"testing"
)

func TestSum(t *testing.T) {
	var x = rand.Int()
	var y = rand.Int()
	var expected = x + y;
	var actual = Sum(x, y);
	if (expected != actual) {
		t.Error("not match")
	}
}
