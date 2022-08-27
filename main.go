package main

import ( 
		"fmt"
		"banco/contas"
)
func main() {

	contaDoAbel := contas.ContaCorrente{
		"Abel",
		200,
		202021,
		500,
	}

	contaDaSilvia := contas.ContaCorrente{
		Titular: "Silvia",
		Saldo: 300,
	}

	status := contaDoAbel.Transferir(-100, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoAbel)

}