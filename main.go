package main

import "fmt"

func main() {

	number := readFloat("Ingrese el valor de x para la funcion F(x) = (x - 3)^2: ")
	size_paso := readFloat("Ingrese el tamaño del paso: ")
	max_iteraciones := readNumber("Ingrese el maximo de iteraciones: ")

	fmt.Println("El resultado de la función F(x) = (x - 3)^2 es:", calculateFunction(number))
	fmt.Println("El resultado de la derivada de la función F'(x) = 2(x - 3) es:", calculateDerivative(number))
	fmt.Println("El valor que minimiza la funcion es x =", 3)

	minimum := calculateMinimunValue(number, size_paso, max_iteraciones)
	fmt.Println("El valor que minimiza la funcion es x =", minimum)
	fmt.Println("El valor minimo de la funcion es F(x) =", calculateFunction(minimum))

}
