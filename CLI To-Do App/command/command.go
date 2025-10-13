package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"todo-app/tasks"
)

// Custom errors
var (
	ErrorEmptyArray = errors.New("La lista de tareas no puede ser nula")
)

// Commands struct
type Commands struct {
	help       bool
	listTasks  bool
	addTask    bool
	updateTask int
	deleteTask int

	// Task attributes
	taskTitle       string
	taskDescription string
	pendingTask     int
	inProgressTask  int
	completedTask   int
}

// Commands constructor
func NewCommands() *Commands {
	commands := Commands{}

	flag.BoolVar(&commands.help, "help", false, "Listar comandos")
	flag.BoolVar(&commands.listTasks, "list", false, "Listar tareas")
	flag.BoolVar(&commands.addTask, "add", false, "Crear una tarea")
	flag.IntVar(&commands.updateTask, "update", 0, "Actualizar una tarea")
	flag.IntVar(&commands.deleteTask, "delete", 0, "Eliminar una tarea")

	// Task attributes
	flag.StringVar(&commands.taskTitle, "title", "", "Modificar el título de una tarea (al crear/actualizar)")
	flag.StringVar(&commands.taskDescription, "description", "", "Modificar la descripción de una tarea (al crear/actualizar)")
	flag.IntVar(&commands.pendingTask, "pending", 0, "Marcar una tarea como pendiente")
	flag.IntVar(&commands.inProgressTask, "in-progress", 0, "Marcar una tarea como en progreso")
	flag.IntVar(&commands.completedTask, "completed", 0, "Marcar una tarea como completada")

	flag.Parse()

	return &commands
}

// List commands function
func ListCommands() {
	fmt.Println("  -list")
	fmt.Println("      Listar todas las tareas")
	fmt.Println()
	fmt.Println("  -add -title \"Título\" -description \"Descripción\"")
	fmt.Println("      Crear una nueva tarea (es obligatorio proporcionar el título)")
	fmt.Println()
	fmt.Println("  -update <ID> -title \"Nuevo título\" -description \"Nueva descripción\"")
	fmt.Println("      Actualizar una tarea existente (es obligatorio proporcionar al menos un atributo)")
	fmt.Println()
	fmt.Println("  -delete <ID>")
	fmt.Println("      Eliminar una tarea existente")
	fmt.Println()
	fmt.Println("  -pending <ID>")
	fmt.Println("      Marcar una tarea existente como: Pendiente")
	fmt.Println()
	fmt.Println("  -in-progress <ID>")
	fmt.Println("      Marcar una tarea existente como: En Progreso")
	fmt.Println()
	fmt.Println("  -completed <ID>")
	fmt.Println("      Marcar una tarea existente como: Completada")
	fmt.Println()
	fmt.Println("  -help")
	fmt.Println("      Listar los comandos válidos")
}

// Execute commands method
func (c *Commands) Execute(tasksList *tasks.Tasks) error {
	if tasksList == nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", ErrorEmptyArray)
		return ErrorEmptyArray
	}

	switch {
		// List commands
		case c.help:
			ListCommands()
			return nil

		// Get tasks
		case c.listTasks:
			tasksList.GetTasks()
			return nil

		// Create task
		case c.addTask:
			if c.taskTitle == "" {
				return fmt.Errorf("El título es obligatorio. Use: -add -title \"Título de la tarea\"")
			}

			if e := tasksList.AddTask(c.taskTitle, c.taskDescription); e != nil {
				return fmt.Errorf("Error al crear tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea creada: \"%s\"", c.taskTitle)

			return nil

		// Update task
		case c.updateTask != 0:
			if c.taskTitle == "" && c.taskDescription == "" {
				return fmt.Errorf("Al menos un atributo debe ser proporcionado para actualizar la tarea. Use: -update <ID> -title \"Título de la tarea\" -description \"Descripción de la tarea\"")
			}

			if e := tasksList.UpdateTask(c.updateTask, c.taskTitle, c.taskDescription); e != nil {
				return fmt.Errorf("Error al actualizar tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea %d actualizada", c.updateTask)

			return nil

		// Delete task
		case c.deleteTask != 0:
			if e := tasksList.DeleteTask(c.deleteTask); e != nil {
				return fmt.Errorf("Error al eliminar tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea %d eliminada", c.deleteTask)

			return nil

		// Change task status (Pending)
		case c.pendingTask != 0:
			if e := tasksList.ChangeTaskStatus(c.pendingTask, tasks.StatusPending); e != nil {
				return fmt.Errorf("Error al cambiar estado de la tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea %d marcada como \"Pendiente\"", c.pendingTask)

			return nil

		// Change task status (In Progress)
		case c.inProgressTask != 0:
			if e := tasksList.ChangeTaskStatus(c.inProgressTask, tasks.StatusInProgress); e != nil {
				return fmt.Errorf("Error al cambiar estado de la tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea %d marcada como \"En Progreso\"", c.inProgressTask)

			return nil

		// Change task status (Completed)
		case c.completedTask != 0:
			if e := tasksList.ChangeTaskStatus(c.completedTask, tasks.StatusCompleted); e != nil {
				return fmt.Errorf("Error al cambiar estado de la tarea -> %w", e)
			}

			fmt.Fprintf(os.Stdout, "Tarea %d marcada como \"Completada\"", c.completedTask)

			return nil

		// Default
		default:
			fmt.Println("Comando no reconocido\nUse -help para ver la lista de comandos válidos")
			return nil
	}
}
