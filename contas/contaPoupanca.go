package contas

import ( 
	"fmt"
	"banco/clientes"
)

type ContaPoupanca struct{
	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {

	podeSacar := valorDoSaque <= float64(c.saldo) && valorDoSaque > 0
	if podeSacar == true {
		c.saldo = c.saldo - float64(valorDoSaque)
		fmt.Println("Saque realizado com sucesso. Seu saldo é de: R$", c.saldo)
		// return "Saque realizado com sucesso."
	} else {
		fmt.Println("Saque inválido. Seu saldo é de: R$", c.saldo)
		// return "Saque inválido." 
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {

	if valorDoDeposito > 0 {
		c.saldo = c.saldo + float64(valorDoDeposito)
		fmt.Println("Depósito efetuado com sucesso. Seu saldo é de: R$", c.saldo)
		// return "Depósito efetuado com sucesso."
	} else {
		fmt.Println("Depósito inválido. Seu saldo é de: R$", c.saldo)
		// return "Depósito inválido."
	}
}

func (c *ContaPoupanca) ObterSaldo() float64{
	return float64(c.saldo)
}