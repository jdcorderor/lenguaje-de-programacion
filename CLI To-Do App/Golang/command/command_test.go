package command

import (
	"testing"
	"todo-app/tasks"
	"todo-app/testutil"
)

// Test: Execute with nil tasks
func TestExecuteWithNilTasks(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{
		listTasks: true,
	}

	e := commands.Execute(nil)

	if e != ErrorEmptyArray {
		t.Errorf("Expected ErrorEmptyArray, but got an unexpected error: %v", e)
	}
}

// Test: Execute with help command
func TestExecuteHelp(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{
		help: true,
	}

	tasksList := &tasks.Tasks{}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
}

// Test: Execute with list command
func TestExecuteList(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{
		listTasks: true,
	}

	tasksList := &tasks.Tasks{}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
}

// Test: Execute add task without title
func TestExecuteAddTaskWithoutTitle(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{
		addTask:         true,
		taskTitle:       "",
		taskDescription: "Description for testing",
	}

	tasksList := &tasks.Tasks{}

	e := commands.Execute(tasksList)

	if e == nil {
		t.Error("Error expected: Can't add task without title")
	}
}

// Test: Execute add task with title
func TestExecuteAddTaskWithTitle(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{
		addTask:         true,
		taskTitle:       "Title for testing",
		taskDescription: "Description for testing",
	}

	tasksList := &tasks.Tasks{}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if len(*tasksList) != 1 {
		t.Errorf("Expected 1 task, but got: %d", len(*tasksList))
	}

	if (*tasksList)[0].Title != "Title for testing" {
		t.Errorf("Expected title \"Title for testing\", but got: '%s'", (*tasksList)[0].Title)
	}
}

// Test: Execute update task without attributes
func TestExecuteUpdateTaskWithoutAttributes(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		updateTask:      1,
		taskTitle:       "",
		taskDescription: "",
	}

	e := commands.Execute(tasksList)

	if e == nil {
		t.Error("Error expected: Can't update task without attributes")
	}
}

// Test: Execute update task with title
func TestExecuteUpdateTaskWithTitle(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		updateTask: 1,
		taskTitle:  "Title 2 for testing",
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Title != "Title 2 for testing" {
		t.Errorf("Expected title \"Title 2 for testing\", but got: '%s'", (*tasksList)[0].Title)
	}
}

// Test: Execute update task with description
func TestExecuteUpdateTaskWithDescription(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		updateTask:      1,
		taskDescription: "Description 2 for testing",
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Description != "Description 2 for testing" {
		t.Errorf("Expected description \"Description 2 for testing\", but got: '%s'", (*tasksList)[0].Description)
	}
}

// Test: Execute update non-existent task
func TestExecuteUpdateNonExistentTask(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	commands := &Commands{
		updateTask: 1,
		taskTitle:  "Title 2 for testing",
	}

	e := commands.Execute(tasksList)

	if e == nil {
		t.Error("Error expected: Can't update non-existent task")
	}
}

// Test: Execute delete task
func TestExecuteDeleteTask(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		deleteTask: 1,
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Visible {
		t.Error("Expected task to be not visible after deletion")
	}
}

// Test: Execute delete non-existent task
func TestExecuteDeleteNonExistentTask(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	commands := &Commands{
		deleteTask: 1,
	}

	e := commands.Execute(tasksList)

	if e == nil {
		t.Error("Error expected: Can't delete non-existent task")
	}
}

// Test: Execute change task status to pending
func TestExecuteChangeTaskStatusToPending(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	tasksList.ChangeTaskStatus(1, tasks.StatusCompleted)

	commands := &Commands{
		pendingTask: 1,
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Status != tasks.StatusPending {
		t.Errorf("Expected status \"Pendiente\", but got: %v", (*tasksList)[0].Status.Stringify())
	}
}

// Test: Execute change task status to in progress
func TestExecuteChangeTaskStatusToInProgress(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		inProgressTask: 1,
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Status != tasks.StatusInProgress {
		t.Errorf("Expected status \"En progreso\", but got: %v", (*tasksList)[0].Status.Stringify())
	}
}

// Test: Execute change task status to completed
func TestExecuteChangeTaskStatusToCompleted(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}
	
	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")

	commands := &Commands{
		completedTask: 1,
	}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	if (*tasksList)[0].Status != tasks.StatusCompleted {
		t.Errorf("Expected status \"Completada\", but got: %v", (*tasksList)[0].Status.Stringify())
	}

	if (*tasksList)[0].CompletedAt.IsZero() {
		t.Error("Expected CompletedAt attribute to be set")
	}
}

// Test: Execute change status of non-existent task
func TestExecuteChangeStatusNonExistentTask(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	tasksList := &tasks.Tasks{}

	commands := &Commands{
		pendingTask: 1,
	}

	e := commands.Execute(tasksList)

	if e == nil {
		t.Error("Error expected: Can't change status of non-existent task")
	}
}

// Test: Execute with no command (default case)
func TestExecuteNoCommand(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	commands := &Commands{}

	tasksList := &tasks.Tasks{}

	e := commands.Execute(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}
}
