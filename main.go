package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64)  {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64)
}

func main() {
	
	contaDoAbel := contas.ContaCorrente{}
	contaDoAbel.Depositar(100)
	PagarBoleto(&contaDoAbel, 50)

	fmt.Println(contaDoAbel.ObterSaldo())
}