package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Producto struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CodProducto      string         `json:"code"`
	NombreProducto string      `json:"product_name"`
	Proveedor	string     `json:"provider"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ProductoRepository interface {
	Create(c context.Context, producto Producto) error
	Fetch(c context.Context) ([]Producto, error)
	FetchById(c context.Context, id int) (Producto, error)
	Update(c context.Context, updatedProducto Producto) error
	Delete(c context.Context, id int) error
}
