// todo.go
// CLI TODO List: gestiona tareas en memoria con comandos sencillos

package main

import (
    "bufio"   // Para leer línea por línea desde stdin
    "fmt"     // Para imprimir en consola
    "os"      // Para acceder a os.Stdin
    "strconv" // Para convertir strings a enteros
    "strings" // Para manipular cadenas de texto
)

// Tarea representa una entrada en nuestra lista TODO.
// ID          → identificador único de la tarea
// Descripcion → texto descriptivo de la tarea
// Hecho       → estado: false = pendiente, true = completada
type Tarea struct {
    ID          int
    Descripcion string
    Hecho       bool
}

func main() {
    // Creamos un Scanner para leer la entrada del usuario línea por línea
    scanner := bufio.NewScanner(os.Stdin)

    // Slice dinámico para almacenar todas las tareas
    var tareas []Tarea

    // Variable para asignar IDs incrementales a cada nueva tarea
    siguienteID := 1

    // Mensaje de bienvenida y comandos disponibles
    fmt.Println("CLI TODO List 🌟")
    fmt.Println("----------------")
    fmt.Println("Comandos: add, list, done, delete, help, exit")

    // Bucle principal: lee y procesa comandos hasta que el usuario salga
    for {
        fmt.Print("> ")            // Prompt de entrada
        if !scanner.Scan() {       // Leer línea; si falla, salimos
            break
        }
        input := strings.TrimSpace(scanner.Text()) // Obtener texto sin espacios extras
        partes := strings.Fields(input)            // Dividir en palabras

        // Si no escribió nada, volvemos a pedir comando
        if len(partes) == 0 {
            continue
        }

        // Primer elemento = comando; resto = argumentos
        cmd := strings.ToLower(partes[0])
        args := partes[1:]

        switch cmd {

        case "add":
            // Añade una nueva tarea: "add Comprar leche"
            if len(args) == 0 {
                fmt.Println("Uso: add <descripción de la tarea>")
                continue
            }
            // Reconstruir la descripción con espacios
            desc := strings.Join(args, " ")
            // Crear un nuevo struct Tarea y añadirlo al slice
            tarea := Tarea{
                ID:          siguienteID,
                Descripcion: desc,
                Hecho:       false,
            }
            tareas = append(tareas, tarea)
            fmt.Printf("✅ Tarea #%d añadida: %s\n", siguienteID, desc)
            siguienteID++ // Incrementar ID para la próxima tarea

        case "list":
            // Lista todas las tareas con su estado
            if len(tareas) == 0 {
                fmt.Println("📭 No hay tareas pendientes")
                continue
            }
            fmt.Println("📋 Tareas:")
            for _, t := range tareas {
                estado := " "    // Por defecto, pendiente
                if t.Hecho {     // Si Hecho == true
                    estado = "✔" // Marcar como completada
                }
                // Mostrar: [✔] #ID: Descripción
                fmt.Printf("  [%s] #%d: %s\n", estado, t.ID, t.Descripcion)
            }

        case "done":
            // Marca una tarea como completada: "done 3"
            if len(args) != 1 {
                fmt.Println("Uso: done <ID de la tarea>")
                continue
            }
            // Convertir argumento a entero
            id, err := strconv.Atoi(args[0])
            if err != nil {
                fmt.Println("ID inválido, debe ser un número")
                continue
            }
            // Buscar la tarea y cambiar su estado
            encontrado := false
            for i := range tareas {
                if tareas[i].ID == id {
                    tareas[i].Hecho = true
                    fmt.Printf("✅ Tarea #%d marcada como hecha\n", id)
                    encontrado = true
                    break
                }
            }
            if !encontrado {
                fmt.Printf("Tarea #%d no encontrada\n", id)
            }

        case "delete":
            // Elimina una tarea: "delete 2"
            if len(args) != 1 {
                fmt.Println("Uso: delete <ID de la tarea>")
                continue
            }
            // Convertir argumento a entero
            id, err := strconv.Atoi(args[0])
            if err != nil {
                fmt.Println("ID inválido, debe ser un número")
                continue
            }
            // Crear un nuevo slice para las tareas restantes
            nuevaLista := tareas[:0] // reutiliza el array subyacente
            eliminado := false
            for _, t := range tareas {
                if t.ID == id {
                    eliminado = true
                    continue // omitimos la tarea a borrar
                }
                nuevaLista = append(nuevaLista, t)
            }
            tareas = nuevaLista // reasignar el slice filtrado
            if eliminado {
                fmt.Printf("🗑️  Tarea #%d eliminada\n", id)
            } else {
                fmt.Printf("Tarea #%d no encontrada\n", id)
            }

        case "help":
            // Mostrar ayuda con todos los comandos disponibles
            fmt.Println("Comandos disponibles:")
            fmt.Println("  add <texto>    – añade una nueva tarea")
            fmt.Println("  list           – lista todas las tareas")
            fmt.Println("  done <ID>      – marca una tarea como hecha")
            fmt.Println("  delete <ID>    – elimina una tarea")
            fmt.Println("  help           – muestra esta ayuda")
            fmt.Println("  exit           – sale del programa")

        case "exit":
            // Salir del programa
            fmt.Println("👋 ¡Hasta luego!")
            return

        default:
            // Comando no reconocido
            fmt.Printf("❓ Comando desconocido: %s (usa “help”)\n", cmd)
        }
    }
}
