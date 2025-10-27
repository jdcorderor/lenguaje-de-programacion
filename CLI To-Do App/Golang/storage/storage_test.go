package storage

import (
	"os"
	"testing"
	"todo-app/tasks"
	"todo-app/testutil"
)

// Test: Create storage function
func TestCreateStorage(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	storage := CreateStorage[tasks.Tasks]("testing.json")
	
	if storage == nil {
		t.Error("Expected storage to be created, but got nil")
	}
	
	if storage.FileName != "testing.json" {
		t.Errorf("Expected filename \"testing.json\", but got: '%s'", storage.FileName)
	}
}

// Test: Upload data function with empty filename
func TestUploadDataWithEmptyFilename(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	storage := &Storage[tasks.Tasks]{FileName: ""}

	tasksList := tasks.Tasks{}
	
	e := storage.UploadData(tasksList)

	if e != ErrorEmptyFileName {
		t.Errorf("Expected ErrorEmptyFileName, but got: %v", e)
	}
}

// Test: Upload data function with valid data
func TestUploadDataWithValidData(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	filename := "testing.json"

	defer os.Remove(filename)
	
	storage := CreateStorage[tasks.Tasks](filename)

	tasksList := tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")
	
	e := storage.UploadData(tasksList)

	if e != nil {
		t.Errorf("Unexpected error: %v", e)
	}

	// ----------------------------------------------------------
	
	// Verify if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Expected file to exist, but it doesn't")
	}
}

// Test: Download data function with empty filename
func TestDownloadDataWithEmptyFilename(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	storage := &Storage[tasks.Tasks]{FileName: ""}

	tasksList := tasks.Tasks{}
	
	e := storage.DownloadData(&tasksList)

	if e != ErrorEmptyFileName {
		t.Errorf("Expected ErrorEmptyFileName, but got: %v", e)
	}
}

// Test: Download data function with nil pointer
func TestDownloadDataWithNilPointer(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	storage := CreateStorage[tasks.Tasks]("testing.json")
	
	e := storage.DownloadData(nil)

	if e != ErrorNilReference {
		t.Errorf("Expected ErrorNilReference, but got: %v", e)
	}
}

// Test: Download data function with valid file
func TestDownloadDataWithValidFile(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	filename := "testing.json"

	defer os.Remove(filename)
	
	storage := CreateStorage[tasks.Tasks](filename)

	// ----------------------------------------------------------

	tasksList := tasks.Tasks{}

	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")
	tasksList.AddTask("Title 2 for testing", "Description 2 for testing")
	
	e := storage.UploadData(tasksList)

	if e != nil {
		t.Errorf("Unexpected error (while uploading): %v", e)
	}

	// ----------------------------------------------------------
	
	downloaded := tasks.Tasks{}

	err := storage.DownloadData(&downloaded)

	if err != nil {
		t.Errorf("Unexpected error (while downloading): %v", err)
	}

	// ----------------------------------------------------------
	
	if len(downloaded) != 2 {
		t.Errorf("Expected 2 tasks, but got: %d", len(downloaded))
	}
	
	if downloaded[0].Title != "Title 1 for testing" {
		t.Errorf("Expected title \"Title 1 for testing\", but got: '%s'", downloaded[0].Title)
	}
	
	if downloaded[1].Title != "Title 2 for testing" {
		t.Errorf("Expected title \"Title 2 for testing\", but got '%s'", downloaded[1].Title)
	}
}

// Test: Upload data and download data functions integration
func TestUploadDownloadIntegration(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	filename := "testing.json"

	defer os.Remove(filename)
	
	storage := CreateStorage[tasks.Tasks](filename)

	// ----------------------------------------------------------
	
	tasksList := tasks.Tasks{}

	tasksList.AddTask("Task 1 for testing", "Description 1 for testing")
	tasksList.ChangeTaskStatus(1, tasks.StatusCompleted)
	
	e := storage.UploadData(tasksList)

	if e != nil {
		t.Errorf("Unexpected error (while uploading): %v", e)
	}

	// ----------------------------------------------------------
	
	downloaded := tasks.Tasks{}

	e = storage.DownloadData(&downloaded)

	if e != nil {
		t.Errorf("Unexpected error (while downloading): %v", e)
	}

	// ----------------------------------------------------------
	
	if len(downloaded) != 1 {
		t.Errorf("Expected 1 task, but got: %d", len(downloaded))
	}
	
	if downloaded[0].Title != "Task 1 for testing" {
		t.Errorf("Expected title \"Task 1 for testing\", but got: '%s'", downloaded[0].Title)
	}
	
	if downloaded[0].Description != "Description 1 for testing" {
		t.Errorf("Expected description \"Description 1 for testing\", but got: '%s'", downloaded[0].Description)
	}
	
	if downloaded[0].Status != tasks.StatusCompleted {
		t.Errorf("Expected status \"Completada\", but got: %v", downloaded[0].Status.Stringify())
	}
	
	if downloaded[0].Visible != true {
		t.Error("Expected task to be visible, but it's not")
	}
}

// Test: Upload data function with complex tasks
func TestUploadDataWithComplexTasks(t *testing.T) {
	defer testutil.SuppressOutput(t)()
	
	filename := "testing.json"

	defer os.Remove(filename)
	
	storage := CreateStorage[tasks.Tasks](filename)

	// ----------------------------------------------------------

	tasksList := tasks.Tasks{}
	
	tasksList.AddTask("Title 1 for testing", "Description 1 for testing")
	tasksList.AddTask("Title 2 for testing", "Description 2 for testing")
	tasksList.AddTask("Title 3 for testing", "Description 3 for testing")
	tasksList.AddTask("Title 4 for testing", "Description 4 for testing")
	
	tasksList.ChangeTaskStatus(2, tasks.StatusInProgress)
	tasksList.ChangeTaskStatus(3, tasks.StatusCompleted)
	tasksList.DeleteTask(4)
	
	e := storage.UploadData(tasksList)
	
	if e != nil {
		t.Errorf("Unexpected error (while uploading): %v", e)
	}

	// ----------------------------------------------------------
	
	downloaded := tasks.Tasks{}

	e = storage.DownloadData(&downloaded)

	if e != nil {
		t.Errorf("Unexpected error (while downloading): %v", e)
	}

	// ----------------------------------------------------------
	
	if len(downloaded) != 4 {
		t.Errorf("Expected 4 tasks, but got: %d", len(downloaded))
	}
	
	if downloaded[0].Status != tasks.StatusPending {
		t.Errorf("Expected task 1 to be \"Pendiente\", but got: %v", downloaded[0].Status.Stringify())
	}
	
	if downloaded[1].Status != tasks.StatusInProgress {
		t.Errorf("Expected task 2 to be \"En progreso\", but got: %v", downloaded[1].Status.Stringify())
	}
	
	if downloaded[2].Status != tasks.StatusCompleted {
		t.Errorf("Expected task 3 to be \"Completada\", but got: %v", downloaded[2].Status.Stringify())
	}
	
	if downloaded[3].Visible != false {
		t.Error("Expected task 4 to be not visible, but it is")
	}
}
