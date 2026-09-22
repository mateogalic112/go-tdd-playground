package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("m", 3)
	fmt.Println(repeated)
	// Output: mmm
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}
