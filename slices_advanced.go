package main

import "fmt"

func main() {
    // 1) Empezamos con un slice de strings
    frutas := []string{"manzana", "banana", "naranja", "uva", "mango"}
    fmt.Println("Inicial:", frutas)

    // 2) Eliminar un elemento (por ejemplo, índice 2 → "naranja")
    //    a) Tomamos todo lo anterior a índice 2: frutas[:2]
    //    b) Tomamos todo lo posterior al índice 2: frutas[3:]
    //    c) Los juntamos con append
    frutas = append(frutas[:2], frutas[3:]...)
    fmt.Println("Después de eliminar índice 2:", frutas)
    //    Resultado: [manzana banana uva mango]

    // 3) Insertar un elemento en el índice 1 (por ejemplo, "kiwi")
    //    a) Partimos en dos slices: antes de 1 y desde 1 en adelante
    //    b) Creamos un slice con el nuevo elemento: []string{"kiwi"}
    //    c) Lo insertamos usando append anidado
    idx := 1
    frutas = append(frutas[:idx], append([]string{"kiwi"}, frutas[idx:]...)...)
    fmt.Println("Después de insertar 'kiwi' en índice 1:", frutas)
    //    Resultado: [manzana kiwi banana uva mango]

    // 4) Fusionar dos slices distintos
    exoticas := []string{"papaya", "pitahaya"}
    todas := append(frutas, exoticas...)
    fmt.Println("Después de fusionar con slice 'exóticas':", todas)
    //    Resultado: [manzana kiwi banana uva mango papaya pitahaya]

    // 5) Demostrar crecimiento de la capacidad
    fmt.Printf("Len=%d, Cap=%d\n", len(todas), cap(todas))
    // Añadimos muchos elementos para forzar realloc
    for i := 0; i < 5; i++ {
        todas = append(todas, fmt.Sprintf("fruta%d", i))
        fmt.Printf("After append #%d → Len=%d, Cap=%d\n", i+1, len(todas), cap(todas))
    }
}
