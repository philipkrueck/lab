package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got: %.2f, want: %.2f", got, want)
		}
	})
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := Perimeter(circle)
	// 	want := 40.0
	//
	// 	if got != want {
	// 		t.Errorf("got: %.2f, want: %.2f", got, want)
	// 	}
	// })
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 73},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v, got: %g, want: %g", test.shape, got, test.want)
			}
		})
	}
}
