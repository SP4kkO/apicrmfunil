package anotacao

import (
	"time"

	"gorm.io/gorm"
)

// Anotacao representa a estrutura de uma anotação associada a uma empresa.
type Anotacao struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Data      time.Time      `json:"data"`
	Assunto   string         `json:"assunto"`
	Anotacao  string         `json:"anotacao"`
	EmpresaID int            `json:"empresa_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
