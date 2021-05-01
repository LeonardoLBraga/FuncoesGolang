package controller

import (
	"math"
)

func Adicao(primeiroNumero, segundoNumero int) int {
	return primeiroNumero + segundoNumero
}

func Subtracao(primeiroNumero, segundoNumero int) int {
	return primeiroNumero - segundoNumero
}

func Multiplicacao(primeiroNumero, segundoNumero int) int {
	return primeiroNumero * segundoNumero
}

func Divisao(primeiroNumero, segundoNumero int) float64 {
	cosciente := float64(primeiroNumero) / float64(segundoNumero)
	return cosciente
}

func Fatorial(numero int) int {
	fatorial := 1
	for i := numero; i >= 1; i-- {
		println(fatorial)
		fatorial *= i
	}
	return fatorial
}

func CalculaValorAbsoluto(numero int) int {
	if numero < 0 {
		return -numero
	}
	return numero
}

func CalculaValorAbsolutoSemIf(numero float64) int {
	numeroAoQuadrado := numero * numero
	raiz := math.Sqrt(numeroAoQuadrado)
	return int(raiz)
}
