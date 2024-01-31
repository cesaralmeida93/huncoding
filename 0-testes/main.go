package main

import "fmt"

func main() {
	// var test int16
	// test = 12345

	// testString := "teste"

	// var testInt int = 123

	// var testFloat64 = 123.45

	// var testInterface interface{}

	// // testInterface = "teste"
	// // testInterface = 123
	// // testInterface = false

	// fmt.Println(test)

	// fmt.Printf("%T\n", testString)
	// fmt.Printf("%T\n", testInt)
	// fmt.Printf("%T\n", testFloat64)
	// fmt.Printf("%T\n", testInterface)

	// var testArray [4]string = [4]string{"teste", "teste", "teste", "teste"}
	// fmt.Printf("%s\n", testArray)
	// fmt.Println(len(testArray))
	// fmt.Println(cap(testArray))

	// var testSlice []string = []string{"teste", "teste", "teste", "teste"}
	// testSlice = append(testSlice, "nome")
	// fmt.Printf("%s\n", testSlice)
	// fmt.Println(len(testSlice))
	// fmt.Println(cap(testSlice))

	// type User struct {
	// 	name  string
	// 	age   int
	// 	test2 string
	// }

	// var user User
	// user = User{
	// 	name:  "nome",
	// 	age:   99,
	// 	test2: "teste2",
	// }
	// fmt.Println(user)

	// for i := 0; i < 11; i++ {
	// 	fmt.Printf("numero: %v\n", i)
	// }

	// teste := 0

	// for teste <= 10 {
	// 	fmt.Printf("valor: %v\n", teste)
	// 	teste++

	// PASSANDO FUNÇÃO COMO PARÂMETRO
	funcaoTest := func(test string, testInt int) {
		fmt.Println(test, testInt)
	}

	test(funcaoTest)

	// RETORNANDO FUNCAO COMO RETORNO
	funcReturn := returnTest()

	funcReturn("nome", 40)

	// FUNCAO ANONIMA
	funcAnonima()

}

// PASSANDO FUNÇÃO COMO PARÂMETRO
func test(value func(string, int)) {

	value("nome", 30)
}

// RETORNANDO FUNCAO COMO RETORNO

func returnTest() func(string, int) {

	impl := func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}

	return impl
}

// FUNCAO ANONIMA
func funcAnonima() {
	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("nome", 50)
}
