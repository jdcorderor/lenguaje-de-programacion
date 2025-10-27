package tasks

// Validate task ID method
func (t *Tasks) ValidateTaskID(id int) error {
	if id < 1 || id > len(*t) {
		return ErrorInvalidTaskID
	}

	if !(*t)[id-1].Visible {
		return ErrorTaskNotFound
	}

	return nil
}
