package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Custom errors
var (
	ErrorEmptyFileName = errors.New("El nombre del archivo no puede estar vacío")
	ErrorNilReference  = errors.New("El puntero de datos no puede ser nulo")
)

// Storage struct
type Storage[T any] struct {
	FileName string
}

// Storage constructor
func CreateStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{FileName: filename}
}

// Upload data method
func (s *Storage[T]) UploadData(data T) error {
	if s.FileName == "" {
		return ErrorEmptyFileName
	}

	filedata, e := json.MarshalIndent(data, "", "    ")

	if e != nil {
		return fmt.Errorf("Error al serializar datos -> %w", e)
	}

	if e := os.WriteFile(s.FileName, filedata, 0644); e != nil {
		return fmt.Errorf("Error al escribir el archivo '%s' -> %w", s.FileName, e)
	}

	return nil
}

// Download data method
func (s *Storage[T]) DownloadData(data *T) error {
	if s.FileName == "" {
		return ErrorEmptyFileName
	}

	if data == nil {
		return ErrorNilReference
	}

	filedata, e := os.ReadFile(s.FileName)

	if e != nil {
		if os.IsNotExist(e) {
			return nil
		}

		return fmt.Errorf("Error al leer el archivo '%s' -> %w", s.FileName, e)
	}

	if e := json.Unmarshal(filedata, data); e != nil {
		return fmt.Errorf("Error al deserializar datos -> %w", e)
	}

	return nil
}
