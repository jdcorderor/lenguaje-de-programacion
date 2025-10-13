package tasks

import "time"

// Change task status method
func (t *Tasks) ChangeTaskStatus(id int, newStatus TaskStatus) error {
	if e := t.ValidateTaskID(id); e != nil {
		return e
	}

	if newStatus < StatusPending || newStatus > StatusCompleted {
		return ErrorInvalidStatus
	}

	(*t)[id-1].Status = newStatus

	if newStatus == StatusCompleted {
		(*t)[id-1].CompletedAt = time.Now()
	} else {
		(*t)[id-1].CompletedAt = time.Time{}
	}

	(*t)[id-1].UpdatedAt = time.Now()

	return nil
}
