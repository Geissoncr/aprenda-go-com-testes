package estruturas

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
	Altura float64
	Base   float64
}

func (t Triangulo) Area() float64 {
	return (t.Altura * t.Base) * 0.5
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Largura * retangulo.Altura
}
