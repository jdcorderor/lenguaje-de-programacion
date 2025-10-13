package tasks

import "time"

// Update task method
func (t *Tasks) UpdateTask(id int, title string, description string) error {
	if e := t.ValidateTaskID(id); e != nil {
		return e
	}

	if title == "" && description == "" {
		return ErrorEmptyFields
	}

	if title != "" {
		(*t)[id-1].Title = title
	}

	if description != "" {
		(*t)[id-1].Description = description
	}

	(*t)[id-1].UpdatedAt = time.Now()

	return nil
}
