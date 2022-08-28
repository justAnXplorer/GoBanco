package contas

import ( 
		"fmt"
		"banco/clientes"
)
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {

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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {

	if valorDoDeposito > 0 {
		c.saldo = c.saldo + float64(valorDoDeposito)
		fmt.Println("Depósito efetuado com sucesso. Seu saldo é de: R$", c.saldo)
		// return "Depósito efetuado com sucesso."
	} else {
		fmt.Println("Depósito inválido. Seu saldo é de: R$", c.saldo)
		// return "Depósito inválido."
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferencia < float64(c.saldo) && valorDaTransferencia > 0 {
		c.saldo = c.saldo - float64(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		fmt.Println("Tranferência realizada com sucesso. Seu saldo é de: R$", c.saldo)
		return true
	} else {
		fmt.Println("Transferencia inválida. Seu saldo é de: R$", c.saldo)
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64{
	return float64(c.saldo)
}