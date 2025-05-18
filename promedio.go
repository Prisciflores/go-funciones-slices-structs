package main

import "fmt"

// Función que calcula el promedio de tres números
func promedio(a, b, c float64) float64 {
    suma := a + b + c              // Sumamos los tres números
    return suma / 3                // Dividimos entre 3 para obtener el promedio
}

func main() {
    // Definimos tres valores para calcular el promedio
    var x, y, z float64 = 6.5, 8.0, 9.5

    // Llamamos a la función y guardamos el resultado
    resultado := promedio(x, y, z)

    // Imprimimos el resultado con dos decimales
    fmt.Printf("El promedio de %.1f, %.1f y %.1f es: %.2f\n", x, y, z, resultado)
}
