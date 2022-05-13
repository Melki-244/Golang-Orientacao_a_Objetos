package main

import "fmt"

type ContaCorrente struct {
  nome string
  numeroDaConta int
  numeroDaAgencia int
  saldo float64
}

func main()  {
  fmt.Println(ContaCorrente{})
}
