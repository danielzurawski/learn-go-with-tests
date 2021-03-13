package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 7)
	fmt.Println(sum)
	// Output: 8
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}
