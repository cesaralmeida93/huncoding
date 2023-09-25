package main

import "fmt"

func main() {
	//FOR
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//WHILE
	test := 0
	for test <= 10 {
		fmt.Println("VALOR: ", test)
		test++
	}

	//DO-WHILE
	// o código printa "PASSOU AQUI" porque é um do-while, mesmo se a condição do loop for falsa, ele executa ao menos uma vez o bloco
	anExpression := false

	for ok := true; ok; ok = anExpression {
		fmt.Println("PASSOU AQUI")
	}

	//FOR-COM-RANGE
	testRange := []string{"teste1", "teste2", "teste3"}

	for key, value := range testRange {
		fmt.Println(key, value)
	}

	value, err := testFunc()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	funcParam := func(test string, testInt int) {
		fmt.Println(test, testInt)
	}

	testParam(funcParam)

	returnTest := testReturn()

	returnTest("cesar", 30)

	testAnonimo()

	receivingStringParameters("TEST1", "TEST2", "TEST3")

	funcaoTest := func() {
		fmt.Println("TEST DE DENTRO DA FUNCAO")
	}

	receivingFunctionParameters(funcaoTest, funcaoTest)

	funcaoTestWithParameters := func(valueString string, valueInt int) {
		fmt.Println(valueInt, valueString)
	}

	receivingFunctionWithParameterByParameters(funcaoTestWithParameters, funcaoTestWithParameters)
}

//FUNÇÕES
func testFunc() (retornoString string, retornoErro error) {
	retornoErro = nil
	retornoString = "test"

	return retornoString, retornoErro
}

// PASSANDO FUNÇÕES POR PARAMETRO
func testParam(value func(string, int)) {
	value("cesar", 30)
}

// RETORNAR UMA FUNÇÃO
func testReturn() func(string, int) {
	funcaoTest := func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}

	return funcaoTest
}

// FUNÇÃO ANÔNIMA
func testAnonimo() {
	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("cesar", 30)
}

//PASSANDO VÁRIOS VALORES POR PARÂMETRO
func receivingStringParameters(stringValues ...string) {
	for _, x := range stringValues {
		fmt.Println(x)
	}
}

func receivingFunctionParameters(stringsValues ...func()) {
	for _, x := range stringsValues {
		x()
	}
}

func receivingFunctionWithParameterByParameters(stringsValues ...func(string, int)) {
	for _, x := range stringsValues {
		x("TEST", 300)
	}
}
