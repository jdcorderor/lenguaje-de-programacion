use chrono::{DateTime, FixedOffset, TimeZone};
use serde::{Deserialize, Serialize};
use serde_repr::{Deserialize_repr, Serialize_repr};
use std::fmt;
use std::ops::{Deref, DerefMut};

// Errors shown to the user while executing CRUD functions
#[derive(Debug, Clone)]
pub enum TaskError {
    InvalidTaskId,
    TaskNotFound,
    EmptyTitle,
    EmptyFields,
}

// Implement fmt::Display for TaskError
impl fmt::Display for TaskError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            TaskError::InvalidTaskId => write!(f, "ID de tarea inválido"),
            TaskError::TaskNotFound => write!(f, "Tarea no encontrada"),
            TaskError::EmptyTitle => write!(f, "El título de la tarea no puede estar vacío"),
            TaskError::EmptyFields => write!(f, "Al menos un campo debe ser proporcionado para actualizar la tarea"),
        }
    }
}

// TaskStatus enum, serialized as an integer (using Serialize_repr/Deserialize_repr)
#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize_repr, Deserialize_repr)]
#[repr(u8)]
pub enum TaskStatus {
    Pending = 0,
    InProgress = 1,
    Completed = 2,
}

// TaskStatus implementation
impl TaskStatus {
    // TaskStatus stringify method
    pub fn stringify(self) -> &'static str {
        match self {
            TaskStatus::Pending => "Pendiente",
            TaskStatus::InProgress => "En progreso",
            TaskStatus::Completed => "Completada",
        }
    }
}

// Zero datetime representation for CompletedAt attribute (when a task has not been completed)
pub fn zero_datetime() -> DateTime<FixedOffset> {
    FixedOffset::east_opt(0)
        .unwrap()
        .with_ymd_and_hms(1, 1, 1, 0, 0, 0)
        .unwrap()
}

// Task struct
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Task {
    #[serde(rename = "ID")]
    pub id: usize,

    #[serde(rename = "Title")]
    pub title: String,

    #[serde(rename = "Description")]
    pub description: String,

    #[serde(rename = "CreatedAt")]
    pub created_at: DateTime<FixedOffset>,

    #[serde(rename = "UpdatedAt")]
    pub updated_at: DateTime<FixedOffset>,

    #[serde(rename = "CompletedAt")]
    pub completed_at: DateTime<FixedOffset>, 

    #[serde(rename = "Status")]
    pub status: TaskStatus,

    #[serde(rename = "Visible")]
    pub visible: bool,
}

// Implement Default for Task
impl Default for Task {
    fn default() -> Self {
        Self {
            id: 0,
            title: String::new(),
            description: String::new(),
            created_at: FixedOffset::east_opt(0).unwrap().from_local_datetime(&chrono::NaiveDate::from_ymd_opt(1970,1,1).unwrap().and_hms_opt(0,0,0).unwrap()).unwrap(),
            updated_at: FixedOffset::east_opt(0).unwrap().from_local_datetime(&chrono::NaiveDate::from_ymd_opt(1970,1,1).unwrap().and_hms_opt(0,0,0).unwrap()).unwrap(),
            completed_at: zero_datetime(),
            status: TaskStatus::Pending,
            visible: true,
        }
    }
}

// Tasks struct
#[derive(Debug, Clone, Serialize, Deserialize, Default)]
pub struct Tasks(pub Vec<Task>);

// Implement Deref for Tasks (allows read-only access)
impl Deref for Tasks {
    type Target = Vec<Task>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

// Implement DerefMut for Tasks (allows mutable access)
impl DerefMut for Tasks {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}
