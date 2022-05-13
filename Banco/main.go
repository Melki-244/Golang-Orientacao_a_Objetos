package main

import "fmt"

type ContaCorrente struct {
  nome string
  numeroDaConta int
  numeroDaAgencia int
  saldo float64
}

func main()  {
  contaGuilherme := ContaCorrente{nome: "Guilherme", numeroDaConta: 111333, numeroDaAgencia:999000, saldo: 30.58} 
  fmt.Println(contaGuilherme)

  contaMelki := ContaCorrente{"Melki", 111222, 999111, 99.08}
  fmt.Println(contaMelki)
}
