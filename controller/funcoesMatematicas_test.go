package controller

import "testing"

func Test_Adicao(t *testing.T) {
	if Adicao(2, 2) != 4 {
		t.Error("Esperado 2 + 2 = 4")
	}
}

func Test_Subtracao(t *testing.T) {
	if Subtracao(2, 2) != 0 {
		t.Error("Esperado 2 - 2 = 0")
	}
}

func Test_Multiplicacao(t *testing.T) {
	if Multiplicacao(2, 3) != 6 {
		t.Error("Esperado 2 * 3 = 6")
	}
}

func Test_Divisao(t *testing.T) {
	if Divisao(4, 2) != 2 {
		t.Error("Esperado 4 / 2 = 2")
	}
}

func Test_Fatorial(t *testing.T) {
	if Fatorial(4) != 24 {
		t.Error("Esperado 4! = 24")
	}
}
