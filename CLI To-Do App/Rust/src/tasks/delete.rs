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
    // DELETE method
    pub fn delete_task(&mut self, id: usize) -> Result<(), TaskError> {
        self.validate_task_id(id)?;

        let idx = id - 1;

        self[idx].visible = false;

        self[idx].updated_at = now_fixed();

        Ok(())
    }
}
