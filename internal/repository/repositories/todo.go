package repositories

import (
	dtos "ToDoApp/internal/models/filters"
)

type ToDoRepository interface {
	Create(*dtos.CreateToDoFilterDTO) error
	// Delete(id uuid.UUID) error
	// Get(id uuid.UUID) (*dtos.ToDoDTO, error)
	// Update(*dtos.ToDoDTO) (*dtos.ToDoDTO, error)
}
