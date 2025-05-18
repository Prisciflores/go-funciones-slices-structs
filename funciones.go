package main

import "fmt"

// Función que suma dos enteros
func sumar(a int, b int) int {
    return a + b
}

// Función que dice si un número es par
func esPar(n int) bool {
    return n%2 == 0
}

func main() {
    fmt.Println("Resultado de sumar(4, 7):", sumar(4, 7))

    numero := 10
    if esPar(numero) {
        fmt.Printf("%d es par\n", numero)
    } else {
        fmt.Printf("%d es impar\n", numero)
    }
}
