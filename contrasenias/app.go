// Generador de Contraseñas Aleatorias:
// Desarrolla un generador de contraseñas que pueda generars contraseñas aleatorias y seguras.

// Desarrollar un generador que me mande letras y numeros de manera random
// En un array mandar todas las letras y numeros
// Pedirle al usuario que mande la longitud que quiera que sea la pw
// Agarrar del array de manera random n cantidad de elementos

// package main

// import(
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	var password []interface{}

// 	for char := 'A'; char <= 'Z'; char++ {
// 		password = append(password, string(char))	
// 	}
// 	for nums := 0; nums <= 9; nums++ {
// 		password = append(password, nums)
// 	}

// 	var longitud int
// 	fmt.Print("Seleccione que tan larga quiere que sea su pw: ")
// 	fmt.Scan(&longitud)

// 	if longitud <= 0 || longitud > len(password) {
// 		fmt.Println("Longitud de password no valida")
// 		return
// 	}

// 	rand.Shuffle(len(password), func(i, j int) {
// 		password[i], password[j] = password[j], password[i]
// 	})

// 	contrasena := password[:longitud]

// 	fmt.Println("Contraseña generada:")
// 	for _, element := range contrasena {
// 		fmt.Print(element, "")
// 	}
	
// }
