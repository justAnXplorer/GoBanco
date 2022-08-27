package main

import ( 
		"fmt"
		"GoBanco/contas"
)
func main() {

	contaDoAbel := contas.contaCorrente{
		"Abel",
		200,
		202021,
		500,
	}

	contaDaSilvia := contas.contaCorrente{
		Titular: "Silvia",
		Saldo: 300,
	}

	status := contaDoAbel.transferir(-100, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoAbel)

}