package tasks

import "time"

// Add task method
func (t *Tasks) AddTask(title string, description string) error {
	if title == "" {
		return ErrorEmptyTitle
	}

	task := Task{
		ID:          len(*t) + 1,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CompletedAt: time.Time{},
		Status:      StatusPending,
		Visible:     true,
	}

	*t = append(*t, task)

	return nil
}
