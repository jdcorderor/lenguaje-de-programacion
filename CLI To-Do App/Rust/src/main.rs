mod command;
mod storage;
mod tasks;

use crate::command::Commands;
use crate::storage::Storage;
use crate::tasks::Tasks;
use std::io::{self, Write};

fn main() {
    // Initialize tasks list
    let mut tasks_list: Tasks = Tasks::default();

    // Initialize storage using JSON file
    let storage: Storage<Tasks> = Storage::new("tasks.json".to_string());

    // Load existing tasks
    match storage.download_data() {
        Ok(Some(data)) => tasks_list = data,
        Ok(None) => {}
        Err(e) => {
            let _ = writeln!(io::stderr(), "Error al cargar las tareas: {}", e);
        }
    }

    // Parse commands
    let commands = Commands::parse_from_env();

    // Execute commands
    if let Err(e) = commands.execute(&mut tasks_list) {
        let _ = writeln!(io::stderr(), "Error al ejecutar los comandos: {}", e);
    }

    // Save tasks to storage
    if let Err(e) = storage.upload_data(&tasks_list) {
        let _ = writeln!(io::stderr(), "Error al guardar las tareas: {}", e);
    }
}
