package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ProductoUseCase struct{}

func (eu *ProductoUseCase) Create(c context.Context, producto domain.Producto) error {
	db := bootstrap.DB
	err := db.Create(&producto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ProductoUseCase) Fetch(c context.Context) ([]domain.Producto, error) {
	db := bootstrap.DB
	entity := []domain.Producto{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ProductoUseCase) FetchById(c context.Context, id int) (domain.Producto, error) {
	db := bootstrap.DB
	pedido := domain.Producto{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Producto{}, err.Error
	}
	return pedido, nil
}

func (eu *ProductoUseCase) Update(c context.Context, updatedProducto domain.Producto) error {
	db := bootstrap.DB
	if err := db.Model(&updatedProducto).
		Omit("deleted_at", "created_at").
		Updates(updatedProducto).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ProductoUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Producto{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
