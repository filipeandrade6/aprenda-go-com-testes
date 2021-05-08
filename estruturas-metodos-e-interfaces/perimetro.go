package perimetro

import "math"

type Forma interface {
	Perimetro() float64
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Altura + r.Largura)
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Perimetro() float64 {
	return 2.0 * math.Pi * c.Raio
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
