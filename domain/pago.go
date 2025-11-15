package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pago struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CodPago      string         `json:"code"`
	FechaPago string      `json:"payment_date"`
	DetallePago	string     `json:"detail"`
	MetodoPago   string         `json:"method"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type PagoRepository interface {
	Create(c context.Context, pago Pago) error
	Fetch(c context.Context) ([]Pago, error)
	FetchById(c context.Context, id int) (Pago, error)
	Update(c context.Context, updatedPago Pago) error
	Delete(c context.Context, id int) error
}
