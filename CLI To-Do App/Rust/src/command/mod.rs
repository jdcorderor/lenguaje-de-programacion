use clap::{ArgAction, Parser};
use std::fmt;

use crate::tasks::{TaskStatus, Tasks};

// Errors shown to the user when handling commands
#[derive(Debug)]
pub enum CommandError {
    InvalidArgs(String),
    TaskError(String),
}

// Implement fmt::Display for CommandError
impl fmt::Display for CommandError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            CommandError::InvalidArgs(s) => write!(f, "{}", s),
            CommandError::TaskError(s) => write!(f, "{}", s),
        }
    }
}

// Define CLI interface, using clap for flags and arguments parsing
#[derive(Parser, Debug, Clone)]
#[command(author, version, about = "CLI To-Do App", long_about = None, disable_help_flag = true)]
pub struct Commands {
    // List commands
    #[arg(long, action = ArgAction::SetTrue)]
    help: bool,

    // List tasks
    #[arg(long, action = ArgAction::SetTrue)]
    list: bool,

    // Add task
    #[arg(long, action = ArgAction::SetTrue)]
    add: bool,

    // Update task (ID)
    #[arg(long, default_value_t = 0)]
    update: i32,

    // Delete task (ID)
    #[arg(long, default_value_t = 0)]
    delete: i32,

    // Task attributes: title
    #[arg(long, default_value = "")]
    title: String,

    // Task attributes: description
    #[arg(long, default_value = "")]
    description: String,

    // Mark task as pending (ID)
    #[arg(long = "pending", default_value_t = 0)]
    pending: i32,

    // Mark task as in progress (ID)
    #[arg(long = "in-progress", default_value_t = 0)]
    in_progress: i32,

    // Mark task as completed (ID)
    #[arg(long = "completed", default_value_t = 0)]
    completed: i32,
}

// Commands implementation
impl Commands {
    // Parse arguments from environment
    pub fn parse_from_env() -> Self {
        <Self as Parser>::parse()
    }

    // Execute commands on the task list
    pub fn execute(&self, tasks_list: &mut Tasks) -> Result<(), CommandError> {
        if self.help {
            Self::list_commands();
            return Ok(());
        }

        if self.list {
            tasks_list.get_tasks();
            return Ok(());
        }

        if self.add {
            if self.title.trim().is_empty() {
                return Err(CommandError::InvalidArgs(
                    "El título es obligatorio. Use: --add --title \"Título de la tarea\"".to_string(),
                ));
            }
            
            if let Err(e) = tasks_list.add_task(self.title.clone(), self.description.clone()) {
                return Err(CommandError::TaskError(format!("Error al crear tarea -> {}", e)));
            }

            println!("Tarea creada: \"{}\"", self.title);
            return Ok(());
        }

        if self.update != 0 {
            if self.title.trim().is_empty() && self.description.trim().is_empty() {
                return Err(CommandError::InvalidArgs("Al menos un atributo debe ser proporcionado para actualizar la tarea. Use: --update <ID> --title \"Título de la tarea\" --description \"Descripción de la tarea\"".to_string()));
            }

            if let Err(e) = tasks_list.update_task(self.update as usize, self.title.clone(), self.description.clone()) {
                return Err(CommandError::TaskError(format!("Error al actualizar tarea -> {}", e)));
            }

            println!("Tarea {} actualizada", self.update);
            return Ok(());
        }

        if self.delete != 0 {
            if let Err(e) = tasks_list.delete_task(self.delete as usize) {
                return Err(CommandError::TaskError(format!("Error al eliminar tarea -> {}", e)));
            }

            println!("Tarea {} eliminada", self.delete);
            return Ok(());
        }

        if self.pending != 0 {
            if let Err(e) = tasks_list.change_task_status(self.pending as usize, TaskStatus::Pending) {
                return Err(CommandError::TaskError(format!("Error al cambiar estado de la tarea -> {}", e)));
            }
            
            println!("Tarea {} marcada como \"Pendiente\"", self.pending);
            return Ok(());
        }

        if self.in_progress != 0 {
            if let Err(e) = tasks_list.change_task_status(self.in_progress as usize, TaskStatus::InProgress) {
                return Err(CommandError::TaskError(format!("Error al cambiar estado de la tarea -> {}", e)));
            }
            
            println!("Tarea {} marcada como \"En Progreso\"", self.in_progress);
            return Ok(());
        }

        if self.completed != 0 {
            if let Err(e) = tasks_list.change_task_status(self.completed as usize, TaskStatus::Completed) {
                return Err(CommandError::TaskError(format!("Error al cambiar estado de la tarea -> {}", e)));
            }
            
            println!("Tarea {} marcada como \"Completada\"", self.completed);
            return Ok(());
        }

        println!("Comando no reconocido\nUse --help para ver la lista de comandos válidos");
        Ok(())
    }

    // Print list of supported commands
    pub fn list_commands() {
        println!("  --list");
        println!("      Listar todas las tareas\n");
        println!("  --add --title \"Título\" --description \"Descripción\"");
        println!("      Crear una nueva tarea (es obligatorio proporcionar el título)\n");
        println!("  --update <ID> --title \"Nuevo título\" --description \"Nueva descripción\"");
        println!("      Actualizar una tarea existente (es obligatorio proporcionar al menos un atributo)\n");
        println!("  --delete <ID>");
        println!("      Eliminar una tarea existente\n");
        println!("  --pending <ID>");
        println!("      Marcar una tarea existente como: Pendiente\n");
        println!("  --in-progress <ID>");
        println!("      Marcar una tarea existente como: En Progreso\n");
        println!("  --completed <ID>");
        println!("      Marcar una tarea existente como: Completada\n");
        println!("  --help");
        println!("      Listar los comandos válidos");
    }
}

// Alternative function for testing: simulates null task list using Option
#[cfg(test)]
impl Commands {
    pub fn execute_option(&self, tasks_opt: Option<&mut Tasks>) -> Result<(), CommandError> {
        match tasks_opt {
            Some(tasks) => self.execute(tasks),
            None => Err(CommandError::InvalidArgs(
                "La lista de tareas no puede ser nula".to_string(),
            )),
        }
    }
}

// Test configuration
#[cfg(test)]
mod tests;