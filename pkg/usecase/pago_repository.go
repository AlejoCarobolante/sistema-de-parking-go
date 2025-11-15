package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type PagoUseCase struct{}

func (eu *PagoUseCase) Create(c context.Context, pago domain.Pago) error {
	db := bootstrap.DB
	err := db.Create(&pago)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *PagoUseCase) Fetch(c context.Context) ([]domain.Pago, error) {
	db := bootstrap.DB
	entity := []domain.Pago{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *PagoUseCase) FetchById(c context.Context, id int) (domain.Pago, error) {
	db := bootstrap.DB
	pedido := domain.Pago{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Pago{}, err.Error
	}
	return pedido, nil
}

func (eu *PagoUseCase) Update(c context.Context, updatedPago domain.Pago) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPago).
		Omit("deleted_at", "created_at").
		Updates(updatedPago).Error; err != nil {
		return err
	}
	return nil
}

func (eu *PagoUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Pago{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
