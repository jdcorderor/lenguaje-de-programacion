## Comandos

### Listar comandos
```bash
go run . -help
```

### Listar tareas
```bash
go run . -list
```

### Crear tarea
```bash
# Crear (con título)
go run . -add -title "Título"

# Crear (con título y descripción)
go run . -add -title "Título" -description "Descripción"
```

### Actualizar tarea
```bash
# Actualizar (con título)
go run . -update 1 -title "Nuevo título"

# Actualizar (con descripción)
go run . -update 1 -description "Nueva descripción"

# Actualizar (con título y descripción)
go run . -update 1 -title "Nuevo título" -description "Nueva descripción"
```

### Eliminar tarea
```bash
go run . -delete 1
```

### Cambiar estado de tarea

```bash
# Marcar como "pendiente"
go run . -pending 1

# Marcar como "en progreso"
go run . -in-progress 1

# Marcar como "completada"
go run . -completed 1
```