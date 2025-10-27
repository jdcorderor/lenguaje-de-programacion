package main

import (
	"fmt"
	"os"
	"todo-app/command"
	"todo-app/storage"
	"todo-app/tasks"
)

func main() {
	// Initialize tasks
	tasksList := tasks.Tasks{}

	// Create storage
	storageInstance := storage.CreateStorage[tasks.Tasks]("tasks.json")

	// Load existing tasks
	if e := storageInstance.DownloadData(&tasksList); e != nil {
		fmt.Fprintf(os.Stderr, "Error al cargar las tareas: %v\n", e)
	}

	commands := command.NewCommands()

	// Execute commands
	if e := commands.Execute(&tasksList); e != nil {
		fmt.Fprintf(os.Stderr, "Error al ejecutar los comandos: %v\n", e)
	}

	// Save tasks
	if e := storageInstance.UploadData(tasksList); e != nil {
		fmt.Fprintf(os.Stderr, "Error al guardar las tareas: %v\n", e)
	}
}