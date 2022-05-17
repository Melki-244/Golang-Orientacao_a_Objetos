package main

import (
 f "fmt"
 c "github.com/Melki-244/Golang-Orientacao_a_Objetos/Banco/contas"
)

func main()  {
  contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300.00}
  contaDoMarcos := c.ContaCorrente{Titular: "Marcos", Saldo: 1200.00}

  mensagem, status := contaDaSilvia.Transfere(-300,&contaDoMarcos)

  f.Println(mensagem,status)
  f.Println(contaDaSilvia.Saldo)
  f.Println(contaDoMarcos.Saldo)
}

