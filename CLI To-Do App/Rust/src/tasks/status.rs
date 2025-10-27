use chrono::{DateTime, FixedOffset, Local};
use chrono::Offset;

use super::types::{TaskError, TaskStatus, Tasks};

// Get local datetime with timezone
fn now_fixed() -> DateTime<FixedOffset> {
    let local_now = Local::now();
    let offset = local_now.offset().fix();
    local_now.with_timezone(&offset)
}

// Tasks implementation
impl Tasks {
    // Change TaskStatus method
    pub fn change_task_status(&mut self, id: usize, new_status: TaskStatus) -> Result<(), TaskError> {
        self.validate_task_id(id)?;

        let idx = id - 1;

        match new_status {
            TaskStatus::Pending | TaskStatus::InProgress | TaskStatus::Completed => {}
        }

        self[idx].status = new_status;

        if matches!(new_status, TaskStatus::Completed) {
            self[idx].completed_at = now_fixed();
        } else {
            self[idx].completed_at = super::types::zero_datetime();
        }

        self[idx].updated_at = now_fixed();

        return Ok(());
    }
}
