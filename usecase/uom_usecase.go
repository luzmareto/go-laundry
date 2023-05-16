package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/repository"
)

type UomUsecase interface {
	Register(newUom *model.Uom) error
	FindAll() ([]model.Uom, error)
	FindOne(id string) (*model.Uom, error)
	Edit(updatedUom *model.Uom) error
	Unreg(id string) error
}

type uomUsecase struct {
	repo repository.UomRepository
}

func (u *uomUsecase) Register(newUom *model.Uom) error {
	return u.repo.Create(newUom)
}

func (u *uomUsecase) FindAll() ([]model.Uom, error) {
	return u.repo.FindAll()
}

func (u *uomUsecase) FindOne(id string) (*model.Uom, error) {
	return u.repo.FindOne(id)
}

func (u *uomUsecase) Edit(updatedUom *model.Uom) error {
	return u.repo.Update(updatedUom)
}

func (u *uomUsecase) Unreg(id string) error {
	return u.repo.Delete(id)
}

func NewUomUsecase(repo repository.UomRepository) UomUsecase {
	return &uomUsecase{
		repo,
	}
}
