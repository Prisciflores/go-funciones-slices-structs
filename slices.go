package main

import "fmt"

func main() {
    // 1) Creamos un slice de strings (lista de frutas) con tres elementos iniciales
    frutas := []string{"manzana", "banana", "naranja"}
    fmt.Println("Slice inicial:", frutas)
    // => Imprime: Slice inicial: [manzana banana naranja]

    // 2) Agregamos una nueva fruta al slice usando append
    frutas = append(frutas, "uva", "mango")
    fmt.Println("Después de append:", frutas)
    // => Imprime: Después de append: [manzana banana naranja uva]

    // 3) Consultamos la longitud (len) y la capacidad (cap) del slice
    // len(frutas) nos dice cuántos elementos hay
    // cap(frutas) nos dice cuánta “reserva” de espacio hay antes de reubicar
    fmt.Printf("Len: %d, Cap: %d\n", len(frutas), cap(frutas))
    // => Ejemplo: Len: 4, Cap: 4 (o >4, si Go ya pre-asignó espacio extra)

    // 4) Creamos un sub-slice (rebanada) seleccionando índices 1 a 3
    //   frutas[1:4] toma los elementos en posición 1,2 y 3 (excluye el 4)
    sub := frutas[1:4]
    fmt.Println("Sub-slice [1:4]:", sub)
    // => Imprime: Sub-slice [1:4]: [banana naranja uva]

    // 5) Recorremos el slice completo con for-range
    for i, fruta := range frutas {
        // i es el índice, fruta es el valor en esa posición
        fmt.Printf("Fruta %d: %s\n", i+1, fruta)
    }
    // => Imprime:
    //    Fruta 1: manzana
    //    Fruta 2: banana
    //    Fruta 3: naranja
    //    Fruta 4: uva

    // 6) Finalmente, mostramos de nuevo la cantidad total de elementos
    fmt.Println("Cantidad total:", len(frutas))
}
