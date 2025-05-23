package main

import "fmt"

func main() {
    // 1) If / Else

    // Pedimos un número al usuario
    var n int
    fmt.Print("Ingresa un número entero: ")
    fmt.Scanln(&n)

    // Verificamos si es par o impar
    if n%2 == 0 {
        fmt.Printf("%d es par\n", n)
    } else {
        fmt.Printf("%d es impar\n", n)
    }

    // 2) Switch

    // Pedimos un día (1–7) y mostramos su nombre
    var dia int
    fmt.Print("Ingresa un día de la semana (1-7): ")
    fmt.Scanln(&dia)

    switch dia {
    case 1:
        fmt.Println("Lunes")
    case 2:
        fmt.Println("Martes")
    case 3:
        fmt.Println("Miércoles")
    case 4:
        fmt.Println("Jueves")
    case 5:
        fmt.Println("Viernes")
    case 6:
        fmt.Println("Sábado")
    case 7:
        fmt.Println("Domingo")
    default:
        fmt.Println("Día inválido")
    }

    // 3) For clásico (inicialización; condición; incremento)
    fmt.Println("Contando del 1 al 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // 4) For como while
    fmt.Println("Contando mientras j < 3:")
    j := 0
    for j < 3 {
        fmt.Println("j =", j)
        j++
    }

    // 5) For-Range sobre un slice
    frutas := []string{"manzana", "banana", "uva"}
    fmt.Println("Recorriendo slice de frutas:")
    for idx, fruta := range frutas {
        fmt.Printf("  índice %d → %s\n", idx, fruta)
    }
}
