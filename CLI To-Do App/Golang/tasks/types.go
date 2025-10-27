package tasks

import (
	"errors"
	"time"
)

// Custom errors
var (
	ErrorInvalidTaskID = errors.New("ID de tarea inválido")
	ErrorTaskNotFound  = errors.New("Tarea no encontrada")
	ErrorEmptyTitle    = errors.New("El título de la tarea no puede estar vacío")
	ErrorEmptyFields   = errors.New("Al menos un campo debe ser proporcionado para actualizar la tarea")
	ErrorInvalidStatus = errors.New("Estado de tarea inválido")
)

// Type for task status
type TaskStatus int

// Enum for task status
const (
	StatusPending TaskStatus = iota
	StatusInProgress
	StatusCompleted
)

// Stringify method for task status
func (s TaskStatus) Stringify() string {
	switch s {
	case StatusPending:
		return "Pendiente"
	case StatusInProgress:
		return "En progreso"
	case StatusCompleted:
		return "Completada"
	default:
		return "Desconocido"
	}
}

// Task struct
type Task struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
	Status      TaskStatus
	Visible     bool
}

// Tasks array
type Tasks []Task
