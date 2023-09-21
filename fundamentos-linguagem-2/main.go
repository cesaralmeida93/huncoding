package main

import (
	"errors"
	"fmt"
)

func main() {
	//IF - ELSE - ELSE IF

	// if 1 > 1 {
	// 	fmt.Println("true")
	// } else if 3 > 1 {
	// 	fmt.Println("test")
	// } else {
	// 	fmt.Println("false")
	// }

	//IF COM INIT
	//atribui o valor de test() para a variável retorno, se err for diferente de nil, printa o retorno
	// o valor de retorno não é acessível fora do escopo do if com init
	// if retorno, err := test(); err != nil {
	// 	fmt.Println("true")
	// 	fmt.Println("retorno: ", retorno)
	// }

	// fmt.Println(retorno)

	//SWITCH CASE
	// fallthrough - se o valor do case for igual ao valor da variável do switch case, ele passa para o próximo case
	// break - se o valor do case for igual ao valor da variável do switch case, ele sai do switch
	// default - se nenhum dos cases for atendido, ele passa para o default

	test := "TESTE"

	switch test {

	case "test1", "teste2", "teste3", "TESTE":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")
		// fallthrough

	case "TESTE@":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
		// fallthrough

	case "TESTE@@@@@@@@":
		fmt.Println("CAIU NA TERCEIRA CONDIÇÃO")
		break

	default:
		fmt.Println("CAIU NO DEFAULT")

	}
}

func test() (string, error) {
	return "test", errors.New("test")
}
