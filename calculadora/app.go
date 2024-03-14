// Calculadora: la calculadora se debe poder ingresar numeros y elegir que tipo de de cuenta hacer (suma, resta, multiplicacion, division) y mostrar el resultado.

// package main

// import(
// 	"fmt"
// 	"os"
// )

// func suma(primerNum, segundoNum int) int{
// 	resultado := primerNum + segundoNum
// 	return resultado
// }
// func resta(primerNum, segundoNum int) int{
// 	resultado := primerNum - segundoNum
// 	return resultado
// }
// func multiplicacion(primerNum, segundoNum int) int{
// 	resultado := primerNum * segundoNum
// 	return resultado
// }
// func division(primerNum, segundoNum int) (int, error){
// 	resultado := primerNum / segundoNum
// 	if segundoNum == 0 {
// 		return 0, fmt.Errorf("Error, no se puede dividir por cero")
// 	}
// 	return resultado, nil
// }

// func main() {

// 	fmt.Println("1. Suma")
// 	fmt.Println("2. Resta")
// 	fmt.Println("3. Multiplicacion")
// 	fmt.Println("4. Division")
// 	var operacionBasica string
// 	fmt.Print("Elija la operacion que desee realizar: ")
// 	fmt.Scan(&operacionBasica)

// 	var primerNum int
// 	fmt.Print("Ingrese un numero: ")
// 	fmt.Scan(&primerNum)

// 	var segundoNum int
// 	fmt.Print("Ingrese el segundo numero: ")
// 	fmt.Scan(&segundoNum)

// 	switch operacionBasica {
// 	case "1", "suma":
// 		resultado := suma(primerNum, segundoNum)
// 		fmt.Println("El resultado de la suma es:", resultado)
// 	case "2", "resta":
// 		resultado := resta(primerNum, segundoNum)
// 		fmt.Println("El resultado de la resta es:", resultado)
// 	case "3", "multiplicacion":
// 		resultado := multiplicacion(primerNum, segundoNum)
// 		fmt.Println("El resultado de la multiplicación es:", resultado)
// 	case "4", "division":
// 		resultado, err := division(primerNum, segundoNum)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		fmt.Println("El resultado de la división es:", resultado)
// 	default:
// 		fmt.Println("Error: Operación no válida")
// 		os.Exit(1)
// 	}
// }