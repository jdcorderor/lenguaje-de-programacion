use super::types::{TaskError, Tasks};

// Tasks implementation
impl Tasks {
    // Validate Task ID method
    pub fn validate_task_id(&self, id: usize) -> Result<(), TaskError> {
        if id < 1 || id > self.len() {
            return Err(TaskError::InvalidTaskId);
        }

        if !self[id - 1].visible {
            return Err(TaskError::TaskNotFound);
        }

        Ok(())
    }
}
