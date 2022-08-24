package main

import "fmt"

func main() {
	contaDoLeozao := contaCorrente{
		"Leonardo",
		155,
		123456,
		125.5,
	}

	contaDoAbel := contaCorrente{
		"Abel",
		200,
		202021,
		156548.7,
	}

	contaDoAbreu := contaCorrente{
		"Abreu",
		201,
		202181,
		15654.7,
	}
	fmt.Println(contaDoLeozao,"\n", contaDoAbel)
	fmt.Println(contaDoAbreu)
}

type contaCorrente struct{
	titular string
	numeroAgencia int
	numeroConta int
	saldo float32
}
