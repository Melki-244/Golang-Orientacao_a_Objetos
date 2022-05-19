package contas

import "github.com/Melki-244/Golang-Orientacao_a_Objetos/Banco/clientes"

type ContaPoupanca struct {
  Titular clientes.Titular  
  Operacao, NumeroDaAgencia, NumeroDaConta int
  saldo float64
}
func (conta *ContaPoupanca) Saca (valorDoSaque float64) (string, float64)  {
  podeSacar := valorDoSaque >= 0 && valorDoSaque <= conta.saldo 

  if podeSacar {
    conta.saldo -= valorDoSaque 
    return "Saque Realizado Com Sucesso", conta.saldo
  } else {
    return "Saldo Insulficiente", conta.saldo
  }
}
func (conta *ContaPoupanca) Deposita (valorDoDeposito float64) (string, float64) {
  podeDepositar := valorDoDeposito >= 0
  if podeDepositar {
    conta.saldo += valorDoDeposito
    return "Deposito Realizado Com Sucesso!" , conta.saldo
  } else {
    return "Deposito Não Realizado" , conta.saldo
  }
}
func (conta *ContaPoupanca) ObterSaldo() float64 {
  return conta.saldo 
}
