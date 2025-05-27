// structs_intro.go
package main

import "fmt"

// Definimos un struct llamado Persona con tres campos
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}

func main() {
    // 2. Instanciar un struct usando campos nombrados
    p1 := Persona{
        Nombre: "Priscila Flores",
        Edad:   32,
        Email:  "priscila@example.com",
    }
    fmt.Printf("Persona 1: %+v\n", p1)
    // %+v imprime nombres de campo y valores

    // 3. Instanciar un struct con orden posicional (menos legible)
    p2 := Persona{"Juan Pérez", 28, "juan@example.com"}
    fmt.Printf("Persona 2: %+v\n", p2)

    // 4. Acceder y modificar campos
    fmt.Println("Nombre de p1:", p1.Nombre)
    p1.Edad = 33
    fmt.Println("p1 ahora tiene edad:", p1.Edad)

    // 5. Puntero a struct
    p3 := &Persona{Nombre: "Ana", Edad: 24}
    fmt.Printf("Antes, p3: %+v\n", *p3)
    p3.Email = "ana@mail.com" // Go dereferencia automáticamente
    fmt.Printf("Después, p3: %+v\n", *p3)

    // 6. Slice de structs
    agenda := []Persona{p1, p2, *p3}
    fmt.Println("\nAgenda de personas:")
    for i, persona := range agenda {
        fmt.Printf(" %d) %s (%d años)\n", i+1, persona.Nombre, persona.Edad)
    }

    // 7. Map de structs (clave string → Persona)
    contactos := map[string]Persona{
        p1.Email: p1,
        p2.Email: p2,
    }
    fmt.Println("\nContacto por email:")
    if person, ok := contactos["juan@example.com"]; ok {
        fmt.Printf("Juan existe: %+v\n", person)
    }
}
