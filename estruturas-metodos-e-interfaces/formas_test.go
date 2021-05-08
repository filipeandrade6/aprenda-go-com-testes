package formas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		input Forma
		want  float64
	}{
		{"retangulo", Retangulo{12, 6}, 72.0},
		{"circulo", Circulo{10}, 314.1592653589793},
		{"triangulo", Triangulo{2, 2}, 2.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := tc.input.Area()
			require.Equal(tt, tc.want, got)
		})
	}
}
