package models

import "github.com/google/uuid"

type Movie struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Director string    `json:"director"`
}

func (m *Movie) IsEmptyOrNil() bool {
	return m == nil || *m == Movie{}
}
