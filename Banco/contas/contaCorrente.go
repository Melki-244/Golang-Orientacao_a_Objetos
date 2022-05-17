package contas

/* Criando Uma Estrutura em Go*/
type ContaCorrente struct {
  Titular string
  NumeroDaConta int
  NumeroDaAgencia int
  Saldo float64
}
func (conta *ContaCorrente) Saca (valorDoSaque float64) (string, float64)  {
  podeSacar := valorDoSaque >= 0 && valorDoSaque <= conta.Saldo 

  if podeSacar {
    conta.Saldo -= valorDoSaque 
    return "Saque Realizado Com Sucesso", conta.Saldo
  }else {
    return "Saldo Insulficiente", conta.Saldo
  }
}
func (conta *ContaCorrente) Deposita (valorDoDeposito float64) (string, float64) {
  podeDepositar := valorDoDeposito >= 0
  if podeDepositar {
    conta.Saldo += valorDoDeposito
    return "Deposito Realizado Com Sucesso!" , conta.Saldo
  }else {
    return "Deposito Não Realizado" , conta.Saldo
  }
}
func (conta *ContaCorrente) Transfere (valorDaTranferencia float64, contaDestino *ContaCorrente) (string,bool) {
  podeTranferir := valorDaTranferencia <= conta.Saldo && valorDaTranferencia >= 0   
  if podeTranferir {
    conta.Saldo -= valorDaTranferencia
    contaDestino.Saldo += valorDaTranferencia 
    return "Transação Realizada com Sucesso!", true
  }else {
    return "A Transação Não Pode Ser Efetuada", false
  }
}
