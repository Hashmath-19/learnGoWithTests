package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{2, 4}
	actual := Perimeter(rect)
	expected := 12.0
	if expected != actual {
		t.Errorf("expected: %.2f, but got: %.2f", expected, actual)
	}
}

func TestArea2(t *testing.T) {
	// By declaring an interface, the helper is decoupled from the
	// concrete types and only has the method it needs to do its job.
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("expected: %g, but got: %g", want, got)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{10, 10}
		checkArea(t, rect, 100)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// Table driven tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 100},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{3, 4}, hasArea: 6},
	}

	// we can run individual test cases by running the command:  go test -run TestArea/<test case name>
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v - expected: %g, but got: %g", tt.shape, tt.hasArea, got)
			}
		})
	}

}
