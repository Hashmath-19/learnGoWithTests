package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6
		assertCorrectMessage(t, sum, expected)
	})
	t.Run("Fail", func(t *testing.T) {
		sum := Add(2, 8)
		expected := 6
		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output:6
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
