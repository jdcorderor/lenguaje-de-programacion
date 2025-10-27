## Comandos

### Listar comandos
```bash
cargo run -- --help
```

### Listar tareas
```bash
cargo run -- --list
```

### Crear tarea
```bash
# Crear (con título)
cargo run -- --add --title "Título"

# Crear (con título y descripción)
cargo run -- --add --title "Título" --description "Descripción"
```

### Actualizar tarea
```bash
# Actualizar (con título)
cargo run -- --update 1 --title "Nuevo título"

# Actualizar (con descripción)
cargo run -- --update 1 --description "Nueva descripción"

# Actualizar (con título y descripción)
cargo run -- --update 1 --title "Nuevo título" --description "Nueva descripción"
```

### Eliminar tarea
```bash
cargo run -- --delete 1
```

### Cambiar estado de tarea

```bash
# Marcar como "pendiente"
cargo run -- --pending 1

# Marcar como "en progreso"
cargo run -- --in-progress 1

# Marcar como "completada"
cargo run -- --completed 1
```

### Ejecutar tests
```bash
cargo test
```