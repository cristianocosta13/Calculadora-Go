package main

import (
	"fmt"
)

func main() {
  var n1 float32 
  var n2 float32
  var opc string
  escolha:="sim"
  
  fmt.Println("\t\t\tCalculadora\n")
  fmt.Println("+ ( SOMAR ) | - ( SUBTRAIR ) \n× ( MULTIPLICAR ) | ÷ ( DIVIDIR )\n")

  for escolha=="sim"{
    fmt.Print("1° valor: ")
    fmt.Scan(&n1)
    fmt.Print("Operação: ")
    fmt.Scan(&opc) 
    fmt.Print("2° valor: ")
    fmt.Scan(&n2)

    calcular(opc, n1, n2)
    fmt.Println("Deseja realizar outra operação? (sim | nao)")
    fmt.Scan(&escolha)
  }

  fmt.Println("Funcionamento encerrado") 
  
}


func calcular(opc string, n1 float32, n2 float32){
  if opc=="+"{
    fmt.Println("= ", (n1+n2))
  }else if opc=="-"{
    fmt.Println("= ", (n1-n2))
  }else if opc=="×"{
    fmt.Println("= ", (n1*n2))
  }else if opc=="÷"{
    fmt.Println("= ", (n1/n2))
  }else{
    fmt.Println("Operação inválida! :(")
  }
}
