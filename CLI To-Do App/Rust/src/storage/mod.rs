use serde::de::DeserializeOwned;
use serde::Serialize;
use std::fmt;
use std::fs;
use std::io;

// Storage struct
#[derive(Debug, Clone)]
pub struct Storage<T> {
    pub file_name: String,
    _marker: std::marker::PhantomData<T>,
}

// Storage errors shown to the user when uploading or downloading data
#[derive(Debug)]
pub enum StorageError {
    EmptyFileName,
    Io(io::Error),
    Serde(serde_json::Error),
}

// Implement fmt::Display for StorageError
impl fmt::Display for StorageError { 
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            StorageError::EmptyFileName => write!(f, "El nombre del archivo no puede estar vacío"),
            StorageError::Io(e) => write!(f, "{}", e),
            StorageError::Serde(e) => write!(f, "{}", e),
        }
    }
}

// Storage implementation (T can be any type)
impl<T> Storage<T> {
    pub fn new(file_name: String) -> Self {
        Self { file_name, _marker: std::marker::PhantomData }
    }
}

// Storage implementation (T can be serialized/deserialized)
impl<T> Storage<T>
where
    T: Serialize + DeserializeOwned,
{
    // Upload tasks data to JSON file
    pub fn upload_data(&self, data: &T) -> Result<(), StorageError> {
        if self.file_name.is_empty() { 
            return Err(StorageError::EmptyFileName);
        }

        let content = serde_json::to_string_pretty(data).map_err(StorageError::Serde)?;
        fs::write(&self.file_name, content).map_err(StorageError::Io)?; 
        Ok(())
    }

    // Download tasks data from JSON file
    pub fn download_data(&self) -> Result<Option<T>, StorageError> { 
        if self.file_name.is_empty() {
            return Err(StorageError::EmptyFileName);
        }

        let content = match fs::read_to_string(&self.file_name) { 
            Ok(c) => c,
            Err(e) if e.kind() == io::ErrorKind::NotFound => return Ok(None),
            Err(e) => return Err(StorageError::Io(e)),
        };

        let data: T = serde_json::from_str(&content).map_err(StorageError::Serde)?;
        Ok(Some(data))
    }
}

// Test configuration
#[cfg(test)]
mod tests;