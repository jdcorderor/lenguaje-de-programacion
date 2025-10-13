package tasks

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strconv"
)

// Get tasks method
func (t *Tasks) GetTasks() {
	tasksTable := table.NewWriter()
	tasksTable.SetOutputMirror(os.Stdout)
	tasksTable.AppendHeader(table.Row{"ID", "Título", "Descripción", "Creado en", "Actualizado en", "Estado", "Completado en"})

	for _, task := range *t {
		completedAt := ""

		if task.Status == StatusCompleted {
			completedAt = task.CompletedAt.Format("01-02-2006 15:04:05")
		}

		if task.Visible {
			tasksTable.AppendRow(table.Row{
				strconv.Itoa(task.ID),
				task.Title,
				task.Description,
				task.CreatedAt.Format("01-02-2006 15:04:05"),
				task.UpdatedAt.Format("01-02-2006 15:04:05"),
				task.Status.Stringify(),
				completedAt,
			})
		}
	}

	tasksTable.Render()
}
