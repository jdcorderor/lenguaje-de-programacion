use super::{Storage, StorageError};
use crate::tasks::{Tasks, TaskStatus};
use std::fs;

// Create a temporary file
fn temp_file(name: &str) -> String {
    let temp_dir = std::env::temp_dir();
    let unique = format!(
        "{}_{}_{}.json",
        name,
        std::process::id(),
        std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .unwrap()
            .as_nanos()
    );
    temp_dir.join(unique).to_string_lossy().into_owned()
}

#[test]
// Test: create storage
fn test_create_storage() {
    let s: Storage<Tasks> = Storage::new("testing.json".to_string());
    assert_eq!(s.file_name, "testing.json");
}

#[test]
// Test: upload data with empty filename
fn test_upload_data_with_empty_filename() {
    let s: Storage<Tasks> = Storage::new("".to_string());
    let tasks = Tasks::default();
    let e = s.upload_data(&tasks).unwrap_err();
    assert!(matches!(e, StorageError::EmptyFileName));
}

#[test]
// Test: upload data
fn test_upload_data_with_valid_data() {
    let filename = temp_file("upload_valid");
    let s: Storage<Tasks> = Storage::new(filename.clone());
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    s.upload_data(&tasks).unwrap();
    assert!(fs::metadata(&filename).is_ok());
    let _ = fs::remove_file(filename);
}

#[test]
// Test: download data with empty filename
fn test_download_data_with_empty_filename() {
    let s: Storage<Tasks> = Storage::new("".to_string());
    let err = s.download_data().unwrap_err();
    assert!(matches!(err, StorageError::EmptyFileName));
}

#[test]
// Test: download data
fn test_download_data_with_valid_file() {
    let filename = temp_file("download_valid");
    let s: Storage<Tasks> = Storage::new(filename.clone());
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.add_task("Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    s.upload_data(&tasks).unwrap();
    let downloaded = s.download_data().unwrap().unwrap();
    assert_eq!(downloaded.len(), 2);
    assert_eq!(downloaded[0].title, "Title 1 for testing");
    assert_eq!(downloaded[1].title, "Title 2 for testing");
    let _ = fs::remove_file(filename);
}

#[test]
// Test: upload and download data
fn test_upload_download_integration() {
    let filename = temp_file("integration");
    let s: Storage<Tasks> = Storage::new(filename.clone());
    let mut tasks = Tasks::default();
    tasks.add_task("Task 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.change_task_status(1, TaskStatus::Completed).unwrap();
    s.upload_data(&tasks).unwrap();
    let downloaded = s.download_data().unwrap().unwrap(); 
    assert_eq!(downloaded.len(), 1);
    assert_eq!(downloaded[0].title, "Task 1 for testing");
    assert_eq!(downloaded[0].description, "Description 1 for testing");
    assert_eq!(downloaded[0].status, TaskStatus::Completed);
    assert!(downloaded[0].visible);
    let _ = fs::remove_file(filename);
}

#[test]
// Test: upload data using complex tasks
fn test_upload_data_with_complex_tasks() {
    let filename = temp_file("complex");
    let s: Storage<Tasks> = Storage::new(filename.clone());
    let mut tasks = Tasks::default();
    tasks.add_task("Title 1 for testing".into(), "Description 1 for testing".into()).unwrap();
    tasks.add_task("Title 2 for testing".into(), "Description 2 for testing".into()).unwrap();
    tasks.add_task("Title 3 for testing".into(), "Description 3 for testing".into()).unwrap();
    tasks.add_task("Title 4 for testing".into(), "Description 4 for testing".into()).unwrap();
    tasks.change_task_status(2, TaskStatus::InProgress).unwrap();
    tasks.change_task_status(3, TaskStatus::Completed).unwrap();
    tasks.delete_task(4).unwrap();
    s.upload_data(&tasks).unwrap();
    let downloaded = s.download_data().unwrap().unwrap();
    assert_eq!(downloaded.len(), 4);
    assert_eq!(downloaded[0].status, TaskStatus::Pending);
    assert_eq!(downloaded[1].status, TaskStatus::InProgress);
    assert_eq!(downloaded[2].status, TaskStatus::Completed);
    assert_eq!(downloaded[3].visible, false);
    let _ = fs::remove_file(filename);
}

#[test]
// Test: download data with invalid status value
fn test_download_data_with_invalid_status_value() {
    use std::fs;
    use std::time::{SystemTime, UNIX_EPOCH};
    let dir = std::env::temp_dir();
    let filename = dir
        .join(format!(
            "invalid_status_{}_{}.json",
            std::process::id(),
            SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos()
        ))
        .to_string_lossy()
        .into_owned();
    let json = r#"[{
        "ID": 1,
        "Title": "X",
        "Description": "Y",
        "CreatedAt": "2024-01-01T00:00:00+00:00",
        "UpdatedAt": "2024-01-01T00:00:00+00:00",
        "CompletedAt": "0001-01-01T00:00:00+00:00",
        "Status": 99,
        "Visible": true
    }]"#;
    fs::write(&filename, json).unwrap();
    let s: super::Storage<crate::tasks::Tasks> = super::Storage::new(filename.clone());
    let err = s.download_data().unwrap_err();
    assert!(matches!(err, super::StorageError::Serde(_)));
    let _ = fs::remove_file(filename);
}
