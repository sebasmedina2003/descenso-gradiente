package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateFunction(number float64) float64 {
	return (number - 3) * (number - 3)
}

func calculateDerivative(number float64) float64 {
	return 2 * (number - 3)
}

func readFloat(message string) float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	input_buffer, _ := reader.ReadString('\n')       // Leemos hasta el salto de linea
	input := strings.TrimRight(input_buffer, "\r\n") // Remover el salto de línea de la entrada del usuario

	number, err := strconv.ParseFloat(input, 64) // Convertir la cadena a un número de punto flotante

	if err != nil {
		fmt.Println("Error al convertir la cadena:", err)
	}

	return number
}

func readNumber(message string) int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	input_buffer, _ := reader.ReadString('\n')       // Leemos hasta el salto de linea
	input := strings.TrimRight(input_buffer, "\r\n") // Remover el salto de línea de la entrada del usuario

	number, err := strconv.Atoi(input) // Convertir la cadena a un número entero

	if err != nil {
		fmt.Println("Error al convertir la cadena:", err)
	}

	return number

}

func calculateMinimunValue(number float64, size_paso float64, max_iteraciones int) float64 {
	for i := 0; i < max_iteraciones; i++ {
		// Calcular el gradiente
		gradiente := 2 * (number - 2)

		// Actualizar el peso
		x1 := number - size_paso*gradiente

		// Convergencia
		if math.Abs(x1-number) < 1e-5 {
			fmt.Println("Mínimo encontrado:", x1)
			break
		}

		// Repetir para la siguiente iteración
		number = x1
	}

	return number
}
