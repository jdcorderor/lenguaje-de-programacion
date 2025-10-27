use super::Tasks;
use super::TaskStatus;
use std::{thread, time::Duration};

#[test]
// Test: TaskStatus stringify method
fn test_task_status_stringify() {
    assert_eq!(TaskStatus::Pending.stringify(), "Pendiente");
    assert_eq!(TaskStatus::InProgress.stringify(), "En progreso");
    assert_eq!(TaskStatus::Completed.stringify(), "Completada");
}

#[test]
// Test: Validate Task ID (valid ID)
fn test_validate_task_id_valid() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.add_task("Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    assert!(tasks.validate_task_id(1).is_ok());
    assert!(tasks.validate_task_id(2).is_ok());
}

#[test]
// Test: Validate Task ID (invalid ID)
fn test_validate_task_id_invalid() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    assert!(tasks.validate_task_id(0).is_err());
    assert!(tasks.validate_task_id(5).is_err());
}

#[test]
// Test: Validate Task ID (deleted task)
fn test_validate_task_id_invisible() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.delete_task(1).unwrap();
    assert!(tasks.validate_task_id(1).is_err());
}

#[test]
// Test: Get tasks method
fn test_get_tasks() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.add_task("Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    tasks.get_tasks();
}

#[test]
// Test: Add task method (with title and description)
fn test_add_task_valid() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    assert_eq!(tasks.len(), 1);
    let task = &tasks[0];
    assert_eq!(task.title, "Title 1 for testing");
    assert_eq!(task.description, "Description 1 for testing");
    assert_eq!(task.status, TaskStatus::Pending);
    assert!(task.visible);
    assert_eq!(task.id, 1);
}

#[test]
// Test: Add task method (with empty title)
fn test_add_task_empty_title() {
    let mut tasks = Tasks::default();
    assert!(tasks.add_task("".into(), "Description for testing".into()).is_err());
}

#[test]
// Test: Update task (with title)
fn test_update_task_title() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.update_task(1, "Title 2 for testing".into(), "".into()).unwrap();
    assert_eq!(tasks[0].title, "Title 2 for testing");
    assert_eq!(tasks[0].description, "Description 1 for testing");
}

#[test]
// Test: Update task (with description)
fn test_update_task_description() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.update_task(1, "".into(), "Description 2 for testing".into()).unwrap();
    assert_eq!(tasks[0].title, "Title 1 for testing");
    assert_eq!(tasks[0].description, "Description 2 for testing");
}

#[test]
// Test: Update task (with both title and description)
fn test_update_task_both() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.update_task(1, "Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    assert_eq!(tasks[0].title, "Title 2 for testing");
    assert_eq!(tasks[0].description, "Description 2 for testing");
}

#[test]
// Test: Update task (with empty fields)
fn test_update_task_empty_fields() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    assert!(tasks.update_task(1, "".into(), "".into()).is_err());
}

#[test]
// Test: Update task (with invalid ID)
fn test_update_task_invalid_id() {
    let mut tasks = Tasks::default();
    assert!(tasks.update_task(1, "Title 1 for testing".into(), "Description 1 for testing".into()).is_err());
}

#[test]
// Test: Update task (modifies updated_at timestamp)
fn test_update_task_updates_timestamp() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    let created_at = tasks[0].updated_at;
    thread::sleep(Duration::from_millis(10));
    tasks.update_task(1, "Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    assert!(tasks[0].updated_at > created_at);
}

#[test]
// Test: Change task status to Pending
fn test_change_task_status_to_pending() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.change_task_status(1, TaskStatus::Completed).unwrap();
    tasks.change_task_status(1, TaskStatus::Pending).unwrap();
    assert_eq!(tasks[0].status, TaskStatus::Pending);
}

#[test]
// Test: Change task status to In Progress
fn test_change_task_status_to_in_progress() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.change_task_status(1, TaskStatus::InProgress).unwrap();
    assert_eq!(tasks[0].status, TaskStatus::InProgress);
}

#[test]
// Test: Change task status to Completed
fn test_change_task_status_to_completed() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.change_task_status(1, TaskStatus::Completed).unwrap();
    assert_eq!(tasks[0].status, TaskStatus::Completed);
    assert!(tasks[0].completed_at >= tasks[0].created_at);
}

#[test]
// Test: Change task status (with invalid ID)
fn test_change_task_status_invalid_id() {
    let mut tasks = Tasks::default();
    assert!(tasks.change_task_status(1, TaskStatus::Completed).is_err());
}

#[test]
// Test: Change task status (updates updated_at timestamp)
fn test_change_task_status_updates_timestamp() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    let created_at = tasks[0].updated_at;
    thread::sleep(Duration::from_millis(10));
    tasks.change_task_status(1, TaskStatus::Completed).unwrap();
    assert!(tasks[0].updated_at > created_at);
}

#[test]
// Test: Delete task
fn test_delete_task() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.delete_task(1).unwrap();
    assert!(!tasks[0].visible);
}

#[test]
// Test: Delete task (with invalid ID)
fn test_delete_task_invalid_id() {
    let mut tasks = Tasks::default();
    assert!(tasks.delete_task(1).is_err());
}

#[test]
// Test: Delete task (updates updated_at timestamp)
fn test_delete_task_updates_timestamp() {
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    let created_at = tasks[0].updated_at;
    thread::sleep(Duration::from_millis(10));
    tasks.delete_task(1).unwrap();
    assert!(tasks[0].updated_at > created_at);
}
