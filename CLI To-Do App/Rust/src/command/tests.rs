use super::{Commands, CommandError};
use crate::tasks::{Tasks, TaskStatus};
use clap::Parser;

// Helper for stdout/stderr supression during tests
fn suppress_output<T>(f: impl FnOnce() -> T) -> T { f() }

#[test]
// Test: --help command
fn test_execute_help() {
    suppress_output(|| {
        let commands = Commands::parse_from(["test-bin", "--help"]);
        let mut tasks = Tasks::default();
        assert!(commands.execute(&mut tasks).is_ok());
    });
}

#[test]
// Test: --list command
fn test_execute_list() {
    suppress_output(|| {
        let commands = Commands::parse_from(["test-bin", "--list"]);
        let mut tasks = Tasks::default();
        assert!(commands.execute(&mut tasks).is_ok());
    });
}

#[test]
// Test: --add command without --title
fn test_execute_add_task_without_title() {
    suppress_output(|| {
        let commands = Commands::parse_from(["test-bin", "--add", "--description", "Description for testing"]);
        let mut tasks = Tasks::default();
        let err = commands.execute(&mut tasks).unwrap_err();
        assert!(matches!(err, CommandError::InvalidArgs(_)));
    });
}

#[test]
// Test: --add command with --title
fn test_execute_add_task_with_title() {
    suppress_output(|| {
        let commands = Commands::parse_from([
            "test-bin", "--add", "--title", "Title for testing", "--description", "Description for testing"
        ]);
        let mut tasks = Tasks::default();
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks.len(), 1);
        assert_eq!(tasks[0].title, "Title for testing");
    });
}

#[test]
// Test: --update command without attributes
fn test_execute_update_task_without_attributes() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--update", "1"]);
        let err = commands.execute(&mut tasks).unwrap_err();
        assert!(matches!(err, CommandError::InvalidArgs(_)));
    });
}

#[test]
// Test: --update command with --title
fn test_execute_update_task_with_title() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--update", "1", "--title", "Title 2 for testing"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks[0].title, "Title 2 for testing");
    });
}

#[test]
// Test: --update command with --description
fn test_execute_update_task_with_description() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--update", "1", "--description", "Description 2 for testing"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks[0].description, "Description 2 for testing");
    });
}

#[test]
// Test: --update command on non-existent task
fn test_execute_update_non_existent_task() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        let commands = Commands::parse_from(["test-bin", "--update", "1", "--title", "X"]);
        let err = commands.execute(&mut tasks).unwrap_err();
        assert!(matches!(err, CommandError::TaskError(_)));
    });
}

#[test]
// Test: --delete command
fn test_execute_delete_task() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--delete", "1"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert!(!tasks[0].visible);
    });
}

#[test]
// Test: --delete command on non-existent task
fn test_execute_delete_non_existent_task() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        let commands = Commands::parse_from(["test-bin", "--delete", "1"]);
        let err = commands.execute(&mut tasks).unwrap_err();
        assert!(matches!(err, CommandError::TaskError(_)));
    });
}

#[test]
// Test: --pending command
fn test_execute_change_task_status_to_pending() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title".into(), "Desc".into()).unwrap();
        tasks.change_task_status(1, TaskStatus::Completed).unwrap();
        let commands = Commands::parse_from(["test-bin", "--pending", "1"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks[0].status, TaskStatus::Pending);
    });
}

#[test]
// Test: --in-progress command
fn test_execute_change_task_status_to_in_progress() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title".into(), "Desc".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--in-progress", "1"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks[0].status, TaskStatus::InProgress);
    });
}

#[test]
// Test: --completed command
fn test_execute_change_task_status_to_completed() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        tasks.add_task("Title".into(), "Desc".into()).unwrap();
        let commands = Commands::parse_from(["test-bin", "--completed", "1"]);
        assert!(commands.execute(&mut tasks).is_ok());
        assert_eq!(tasks[0].status, TaskStatus::Completed);
        assert!(tasks[0].completed_at >= tasks[0].created_at);
    });
}

#[test]
// Test: change status  (--pending, --in-progress, --completed) on non-existent task
fn test_execute_change_status_non_existent_task() {
    suppress_output(|| {
        let mut tasks = Tasks::default();
        let commands = Commands::parse_from(["test-bin", "--pending", "1"]);
        let err = commands.execute(&mut tasks).unwrap_err();
        assert!(matches!(err, CommandError::TaskError(_)));
    });
}

#[test]
// Test: execute no command
fn test_execute_no_command() {
    suppress_output(|| {
        let commands = Commands::parse_from(["test-bin"]);
        let mut tasks = Tasks::default();
        assert!(commands.execute(&mut tasks).is_ok());
    });
}

#[test]
// Test: execute with null tasks (with Option::None)
fn test_execute_with_nil_tasks() {
    let commands = Commands::parse_from(["test-bin", "--list"]);
    let mut tasks = Tasks::default();
    assert!(commands.execute_option(Some(&mut tasks)).is_ok());
    let err = commands.execute_option(None).unwrap_err();
    assert!(matches!(err, CommandError::InvalidArgs(_)));
}