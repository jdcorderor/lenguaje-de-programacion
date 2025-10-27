use chrono::{DateTime, FixedOffset, Local};
use chrono::Offset;

use super::types::{zero_datetime, Task, TaskError, Tasks, TaskStatus};

// Get local datetime with timezone
fn now_fixed() -> DateTime<FixedOffset> {
    let local_now = Local::now();
    let offset = local_now.offset().fix();
    local_now.with_timezone(&offset)
}

// Tasks implementation
impl Tasks {
    // POST method
    pub fn add_task(&mut self, title: String, description: String) -> Result<(), TaskError> {
        if title.trim().is_empty() {
            return Err(TaskError::EmptyTitle);
        }

        let now = now_fixed();

        let task = Task {
            id: self.len() + 1,
            title,
            description,
            created_at: now,
            updated_at: now, 
            completed_at: zero_datetime(),
            status: TaskStatus::Pending,
            visible: true,
        };

        self.push(task);
        
        Ok(())
    }
}
