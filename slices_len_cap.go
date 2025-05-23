// slices_len_cap.go
package main

import "fmt"

func main() {
    // 1) Slice literal
    s1 := []int{1, 2, 3}
    fmt.Printf("s1 → %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
    // len=3, cap=3 (cap igual a la longitud inicial)

    // 2) Sub-slice
    arr := [5]string{"A", "B", "C", "D", "E"}
    s2 := arr[1:4] // elementos B, C, D
    fmt.Printf("s2 → %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
    // len=3, cap=4 (desde índice 1 hasta el final del arreglo)

    // 3) Append dentro de la capacidad
    //    Primero, preasignamos con make(cap=5)
    s3 := make([]string, 0, 5)
    fmt.Printf("s3 inicial → %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
    s3 = append(s3, "x", "y", "z")
    fmt.Printf("s3 tras append 3 elems → %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
    // len crece a 3, cap sigue en 5

    // 4) Append que supera la capacidad
    s3 = append(s3, "u", "v", "w")
    fmt.Printf("s3 tras append 3 elems más → %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
    // cap se redimensiona (normalmente dobla) al superar el límite

    // 5) Slice de slices
    s4 := [][]int{
        {1, 2},
        {3, 4, 5},
    }
    fmt.Printf("s4 → %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))
    // En slices de slices, cap indica cuántos elementos de primer nivel caben

    // 6) Sub-slice de s3 y su capacidad
    sub := s3[2:6] // toma de índice 2 hasta 5
    fmt.Printf("sub de s3 → %v, len=%d, cap=%d\n", sub, len(sub), cap(sub))
    // cap refleja lo que queda de capacidad desde el nuevo inicio

    // 7) Slicing vacío con make sin preasignar cap
    s5 := make([]int, 2) // len=2, cap=2 por defecto
    fmt.Printf("s5 → %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
    // s5 inicia con dos ceros y no puede crecer sin realloc

    // 8) Creación con make sólo capacidad
    s6 := make([]int, 0, 3) // len=0, cap=3
    fmt.Printf("s6 inicial → %v, len=%d, cap=%d\n", s6, len(s6), cap(s6))
    // Luego de append 3 elems:
    s6 = append(s6, 7, 8, 9)
    fmt.Printf("s6 tras append 3 elems → %v, len=%d, cap=%d\n", s6, len(s6), cap(s6))
}
