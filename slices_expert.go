package main

import "fmt"

// removeString elimina el elemento en 'index' usando copy y ajusta la longitud
func removeString(s []string, index int) []string {
    // Desplaza todos los elementos desde index+1 a la posición index
    copy(s[index:], s[index+1:])
    // Reduce la longitud en 1 (el último queda duplicado hasta que se recolecte)
    return s[:len(s)-1]
}

func main() {
    // 1) Preasignar un slice vacío con capacidad para 5 elementos
    frutas := make([]string, 0, 5)
    fmt.Printf("Inicial: len=%d cap=%d %v\n", len(frutas), cap(frutas), frutas)
    // => Inicial: len=0 cap=5 []

    // 2) Agregar elementos
    frutas = append(frutas, "manzana", "banana", "naranja")
    fmt.Printf("Tras primer append: len=%d cap=%d %v\n", len(frutas), cap(frutas), frutas)
    // => Tras primer append: len=3 cap=5 [manzana banana naranja]

    // 3) Añadir más allá de la capacidad para forzar realloc
    frutas = append(frutas, "uva", "mango", "papaya")
    fmt.Printf("Tras segundo append: len=%d cap=%d %v\n", len(frutas), cap(frutas), frutas)
    // => Cap crecerá (p.ej. de 5 a 10)

    // 4) Eliminar el elemento en índice 2 ("naranja")
    frutas = removeString(frutas, 2)
    fmt.Printf("Tras remove índice 2: %v\n", frutas)
    // => [manzana banana uva mango papaya]

    // 5) Copiar el slice a uno nuevo e independiente
    copia := make([]string, len(frutas))
    copy(copia, frutas)
    fmt.Printf("Copia antes de modificar: %v\n", copia)
    // => Igual que 'frutas'

    // 6) Modificar ambos para demostrar que ya no comparten arreglo subyacente
    frutas[0] = "cereza"
    copia[1] = "kiwi"
    fmt.Printf("Original modificado: %v\n", frutas)
    fmt.Printf("Copia modificada:   %v\n", copia)
    // => Ves que cada slice mantiene sus propios cambios
}
