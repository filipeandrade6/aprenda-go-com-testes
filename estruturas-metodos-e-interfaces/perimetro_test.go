package perimetro

import "testing"

func TestPerimetro(t *testing.T) {
	tests := []struct {
		name  string
		input Forma
		want  float64
	}{
		{"retangulo", Retangulo{12, 6}, 72.0},
		{"circulo", Circulo{10}, 314.1592653589793},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.Perimetro()
			if got != tc.want {
				t.Errorf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		input Forma
		want  float64
	}{
		{"retangulo", Retangulo{12, 6}, 72.0},
		{"circulo", Circulo{10}, 314.1592653589793},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.Area()
			if got != tc.want {
				t.Errorf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}
