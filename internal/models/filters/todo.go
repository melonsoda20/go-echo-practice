package filters

import (
	"time"
)

type CreateToDoFilterDTO struct {
	Name         string
	Age          int
	Created_By   string
	Created_Date time.Time
}

func GenerateCreateToDoDTO(name string, age int, created_by string) CreateToDoFilterDTO {
	c := CreateToDoFilterDTO{}
	c.Name = name
	c.Age = age
	c.Created_By = created_by
	c.Created_Date = time.Now().UTC()
	return c
}
