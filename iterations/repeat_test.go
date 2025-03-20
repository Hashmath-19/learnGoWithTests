package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)	
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("expected: %s, but got: %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 19)
	}
}

func ExampleRepeat(){
	repeated := Repeat("H", 9)
	fmt.Println(repeated)
	//Output:HHHHHHHHH
}
