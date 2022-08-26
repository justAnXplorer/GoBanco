package main

import ( "fmt"
		 "G:\Meu Drive\2022\GO\GoBanco\contas"
)
func main() {

	contaDoAbel := contas.contaCorrente{
		"Abel",
		200,
		202021,
		500,
	}

	contaDaSilvia := contas.contaCorrente{
		titular: "Silvia",
		saldo: 300,
	}

	status := contaDoAbel.transferir(-100, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoAbel)

}