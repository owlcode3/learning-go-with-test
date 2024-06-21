package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("expected result", func(t *testing.T) {

		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("expected no of time character was repeated", func(t *testing.T) {
		noOfTime := 5
		repeated := Repeat("a", noOfTime)

		if noOfTime != len(repeated) {
			t.Errorf("expected %d but got %d", noOfTime, len(repeated))
		}
	})
}

func Repeat(character string, noOfTime int) (repeated string) {
	for i := 0; i < noOfTime; i++ {
		repeated += character
	}
	return
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 5)
	fmt.Println(result)
	// Output: bbbbb

}
