package main

import "fmt"

func main() {
	//TIPOS DE VARIÁVEIS

	var testString string = "ewewewewewewew"
	fmt.Printf("%T\n", testString)

	testInt := 30
	fmt.Printf("%T\n", testInt)

	testFloat := 30.55
	fmt.Printf("%T\n", testFloat)

	testMap := map[string]string{
		"test": "test",
	}
	fmt.Printf("%T\n", testMap)

	//VARIÁVEIS PÚBLICAS E PRIVADAS

	// //privada
	// func test() {}

	// //publica
	// func Test() {}

	// TIPO INTERFACE

	var testInterface interface{}
	testInterface = "teste"
	fmt.Printf("%T\n", testInterface)
	testInterface = 30
	fmt.Printf("%T\n", testInterface)
	testInterface = 30.55
	fmt.Printf("%T\n", testInterface)
	testInterface = true
	fmt.Printf("%T\n", testInterface)

	testJson := map[string]interface{}{
		"test123": "Test",
		"test456": 50,
	}
	fmt.Printf("%v\n", testJson["test456"])

	// ARRAYS E SLICES

	var testArray [4]string = [4]string{"test1", "test2", "test3", "teste4"}
	// testArray = append(testArray, "test5")
	fmt.Println(testArray)
	fmt.Println(len(testArray)) //tamanho que o array tem
	fmt.Println(cap(testArray)) // capacidade que o array suporta

	// quando você "estoura" a capacidade do slice, ele cria um novo slice com o dobro do tamanho
	var testSlice []string = []string{"test1", "test2", "test3", "test4"}
	testSlice = append(testSlice, "test5")
	fmt.Println(testSlice)
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))

	// STRUCTS

	type User struct {
		name  string
		age   int
		test2 string
	}

	var user User = User{
		name:  "Test",
		age:   30,
		test2: "test2",
	}

	fmt.Println(user)

}
