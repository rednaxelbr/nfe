package nfe

import "testing"

func TestIsNumber(t *testing.T) {

	letras := "ABC"
	ans := isNumber(letras)
	if ans {
		t.Errorf("isNumber(%s) deveria ser false", letras)
	}

	numeros := "1234567890"
	ans = isNumber(numeros)
	if !ans {
		t.Errorf("isNumber(%s) deveria ser true", numeros)
	}
}
