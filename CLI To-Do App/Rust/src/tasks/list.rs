use comfy_table::{Cell, Table};

use super::types::{TaskStatus, Tasks};

// Tasks implementation
impl Tasks {
    // GET method
    pub fn get_tasks(&self) {
        let mut table = Table::new();
        
        table.set_header(vec![
            Cell::new("ID"),
            Cell::new("Título"),
            Cell::new("Descripción"),
            Cell::new("Creado en"),
            Cell::new("Actualizado en"),
            Cell::new("Estado"),
            Cell::new("Completado en"),
        ]);

        for task in self.iter() {
            if !task.visible {
                continue;
            }

            let completed_at = if task.status == TaskStatus::Completed {
                task.completed_at.format("%d-%m-%Y %H:%M:%S").to_string()
            } else {
                String::new()
            };

            table.add_row(vec![
                Cell::new(task.id.to_string()),
                Cell::new(&task.title),
                Cell::new(&task.description),
                Cell::new(task.created_at.format("%d-%m-%Y %H:%M:%S").to_string()),
                Cell::new(task.updated_at.format("%d-%m-%Y %H:%M:%S").to_string()),
                Cell::new(task.status.stringify()),
                Cell::new(completed_at),
            ]);
        }

        println!("{}", table);
    }
}
