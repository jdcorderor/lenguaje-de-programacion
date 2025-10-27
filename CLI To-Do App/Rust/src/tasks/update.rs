use chrono::{DateTime, FixedOffset, Local};
use chrono::Offset;

use super::types::{TaskError, Tasks};

// Get local datetime with timezone
fn now_fixed() -> DateTime<FixedOffset> {
    let local_now = Local::now();
    let offset = local_now.offset().fix();
    local_now.with_timezone(&offset)
}

// Tasks implementation
impl Tasks {
    // PUT method
    pub fn update_task(&mut self, id: usize, title: String, description: String) -> Result<(), TaskError> {
        self.validate_task_id(id)?;

        if title.trim().is_empty() && description.trim().is_empty() {
            return Err(TaskError::EmptyFields);
        }

        let idx = id - 1;

        if !title.trim().is_empty() {
            self[idx].title = title;
        }

        if !description.trim().is_empty() {
            self[idx].description = description;
        }

        self[idx].updated_at = now_fixed(); 
        
        Ok(())
    }
}
