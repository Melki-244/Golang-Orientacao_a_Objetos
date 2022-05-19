package main

import (
	f "fmt"

	 "github.com/Melki-244/Golang-Orientacao_a_Objetos/Banco/clientes"
	 "github.com/Melki-244/Golang-Orientacao_a_Objetos/Banco/contas"
)

func main()  {
  clienteKatia := clientes.Titular{"Katia", "000.000.000-00", "Sei Lá"}
  contaDaKatia := contas.ContaPoupanca{Titular: clienteKatia, NumeroDaAgencia: 222333, NumeroDaConta: 444000222}

  contaDaKatia.Deposita(200)
  f.Println(contaDaKatia.ObterSaldo())
}
