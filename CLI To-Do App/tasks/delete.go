package tasks

import "time"

// Delete task method
func (t *Tasks) DeleteTask(id int) error {
	if e := t.ValidateTaskID(id); e != nil {
		return e
	}

	(*t)[id-1].Visible = false
	(*t)[id-1].UpdatedAt = time.Now()

	return nil
}
