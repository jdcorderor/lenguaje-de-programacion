package tasks

import (
	"testing"
	"time"
	"todo-app/testutil"
)

// Test: Stringify method for TaskStatus
func TestTaskStatusStringify(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tests := []struct {
		status   TaskStatus
		expected string
	}{
		{StatusPending, "Pendiente"},
		{StatusInProgress, "En progreso"},
		{StatusCompleted, "Completada"},
		{TaskStatus(99), "Desconocido"},
	}
	
	for _, test := range tests {
		result := test.status.Stringify()

		if result != test.expected {
			t.Errorf("Expected '%s', but got: '%s'", test.expected, result)
		}
	}
}

// Test: ValidateTaskID method with valid ID
func TestValidateTaskIDValid(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}

	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	tasks.AddTask("Title 2 for testing", "Description 2 for testing")
	
	e := tasks.ValidateTaskID(1)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	e = tasks.ValidateTaskID(2)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
}

// Test: ValidateTaskID method with invalid ID (out of range)
func TestValidateTaskIDInvalid(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}

	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.ValidateTaskID(0)

	if e != ErrorInvalidTaskID {
		t.Errorf("Expected ErrorInvalidTaskID, but got: %v", e)
	}
	
	e = tasks.ValidateTaskID(5)

	if e != ErrorInvalidTaskID {
		t.Errorf("Expected ErrorInvalidTaskID, but got: %v", e)
	}
}

// Test: ValidateTaskID method with invisible task
func TestValidateTaskIDInvisible(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}

	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	tasks.DeleteTask(1)
	
	e := tasks.ValidateTaskID(1)

	if e != ErrorTaskNotFound {
		t.Errorf("Expected ErrorTaskNotFound, but got: %v", e)
	}
}

// Test: GetTasks method
func TestGetTasks(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	tasks.AddTask("Title 2 for testing", "Description 2 for testing")
	
	tasks.GetTasks()
}

// Test: AddTask method with valid data
func TestAddTaskValid(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	e := tasks.AddTask("Title 1 for testing", "Description 1 for testing")

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if len(*tasks) != 1 {
		t.Errorf("Expected 1 task, but got: %d", len(*tasks))
	}
	
	task := (*tasks)[0]

	if task.Title != "Title 1 for testing" {
		t.Errorf("Expected title \"Title 1 for testing\", but got: '%s'", task.Title)
	}
	
	if task.Description != "Description 1 for testing" {
		t.Errorf("Expected description \"Description 1 for testing\", but got: '%s'", task.Description)
	}
	
	if task.Status != StatusPending {
		t.Errorf("Expected status \"Pendiente\", but got: %v", task.Status.Stringify())
	}
	
	if !task.Visible {
		t.Error("Expected task to be visible, but it's not")
	}
	
	if task.ID != 1 {
		t.Errorf("Expected ID 1, but got: %d", task.ID)
	}
}

// Test: AddTask method with empty title
func TestAddTaskEmptyTitle(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	e := tasks.AddTask("", "Description for testing")

	if e != ErrorEmptyTitle {
		t.Errorf("Expected ErrorEmptyTitle, but got: %v", e)
	}
}

// Test: UpdateTask method with title
func TestUpdateTaskTitle(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.UpdateTask(1, "Title 2 for testing", "")

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Title != "Title 2 for testing" {
		t.Errorf("Expected title \"Title 2 for testing\", but got: '%s'", (*tasks)[0].Title)
	}
	
	if (*tasks)[0].Description != "Description 1 for testing" {
		t.Errorf("Expected description \"Description 1 for testing\", but got: '%s'", (*tasks)[0].Description)
	}
}

// Test: UpdateTask method with description
func TestUpdateTaskDescription(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.UpdateTask(1, "", "Description 2 for testing")

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Title != "Title 1 for testing" {
		t.Errorf("Expected title \"Title 1 for testing\", but got: '%s'", (*tasks)[0].Title)
	}
	
	if (*tasks)[0].Description != "Description 2 for testing" {
		t.Errorf("Expected description \"Description 2 for testing\", but got: '%s'", (*tasks)[0].Description)
	}
}

// Test: UpdateTask method with both title and description
func TestUpdateTaskBoth(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.UpdateTask(1, "Title 2 for testing", "Description 2 for testing")
	
	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Title != "Title 2 for testing" {
		t.Errorf("Expected title \"Title 2 for testing\", but got: '%s'", (*tasks)[0].Title)
	}
	
	if (*tasks)[0].Description != "Description 2 for testing" {
		t.Errorf("Expected description \"Description 2 for testing\", but got: '%s'", (*tasks)[0].Description)
	}
}

// Test: UpdateTask method with empty fields
func TestUpdateTaskEmptyFields(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.UpdateTask(1, "", "")
	
	if e != ErrorEmptyFields {
		t.Errorf("Expected ErrorEmptyFields, but got: %v", e)
	}
}

// Test: UpdateTask method with invalid ID
func TestUpdateTaskInvalidID(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	e := tasks.UpdateTask(1, "Title 1 for testing", "Description 1 for testing")
	
	if e != ErrorInvalidTaskID {
		t.Errorf("Expected ErrorInvalidTaskID, but got: %v", e)
	}
}

// Test: UpdateTask method updates UpdatedAt
func TestUpdateTaskUpdatesTimestamp(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}

	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	createdAt := (*tasks)[0].UpdatedAt

	time.Sleep(10 * time.Millisecond)
	
	tasks.UpdateTask(1, "Title 2 for testing", "Description 2 for testing")
	
	if !(*tasks)[0].UpdatedAt.After(createdAt) {
		t.Error("Expected UpdatedAt to be updated, but it's not")
	}
}

// Test: ChangeTaskStatus method to Pending
func TestChangeTaskStatusToPending(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	tasks.ChangeTaskStatus(1, StatusCompleted)
	
	e := tasks.ChangeTaskStatus(1, StatusPending)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Status != StatusPending {
		t.Errorf("Expected status \"Pendiente\", but got: %v", (*tasks)[0].Status.Stringify())
	}
	
	if !(*tasks)[0].CompletedAt.IsZero() {
		t.Error("Expected CompletedAt to be cleared, but it's not")
	}
}

// Test: ChangeTaskStatus method to InProgress
func TestChangeTaskStatusToInProgress(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.ChangeTaskStatus(1, StatusInProgress)
	
	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Status != StatusInProgress {
		t.Errorf("Expected status \"En progreso\", but got: %v", (*tasks)[0].Status.Stringify())
	}
}

// Test: ChangeTaskStatus method to Completed
func TestChangeTaskStatusToCompleted(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.ChangeTaskStatus(1, StatusCompleted)
	
	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Status != StatusCompleted {
		t.Errorf("Expected status \"Completado\", but got: %v", (*tasks)[0].Status.Stringify())
	}
	
	if (*tasks)[0].CompletedAt.IsZero() {
		t.Error("Expected CompletedAt to be set, but it's not")
	}
}

// Test: ChangeTaskStatus method with invalid status
func TestChangeTaskStatusInvalidStatus(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}

	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.ChangeTaskStatus(1, TaskStatus(99))
	
	if e != ErrorInvalidStatus {
		t.Errorf("Expected ErrorInvalidStatus, but got: %v", e)
	}
}

// Test: ChangeTaskStatus method with invalid ID
func TestChangeTaskStatusInvalidID(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	e := tasks.ChangeTaskStatus(1, StatusCompleted)
	
	if e != ErrorInvalidTaskID {
		t.Errorf("Expected ErrorInvalidTaskID, but got: %v", e)
	}
}

// Test: ChangeTaskStatus method updates UpdatedAt
func TestChangeTaskStatusUpdatesTimestamp(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	createdAt := (*tasks)[0].UpdatedAt

	time.Sleep(10 * time.Millisecond)
	
	tasks.ChangeTaskStatus(1, StatusCompleted)
	
	if !(*tasks)[0].UpdatedAt.After(createdAt) {
		t.Error("Expected UpdatedAt to be updated, but it's not")
	}
}

// Test: DeleteTask method
func TestDeleteTask(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := tasks.DeleteTask(1)
	
	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
	
	if (*tasks)[0].Visible {
		t.Error("Expected task to be not visible, but it is")
	}
}

// Test: DeleteTask method with invalid ID
func TestDeleteTaskInvalidID(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	e := tasks.DeleteTask(1)
	
	if e != ErrorInvalidTaskID {
		t.Errorf("Expected ErrorInvalidTaskID, but got: %v", e)
	}
}

// Test: DeleteTask method updates UpdatedAt
func TestDeleteTaskUpdatesTimestamp(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasks := &Tasks{}
	
	tasks.AddTask("Title 1 for testing", "Description 1 for testing")
	
	createdAt := (*tasks)[0].UpdatedAt
	
	time.Sleep(10 * time.Millisecond)
	
	tasks.DeleteTask(1)
	
	if !(*tasks)[0].UpdatedAt.After(createdAt) {
		t.Error("Expected UpdatedAt to be updated, but it's not")
	}
}
