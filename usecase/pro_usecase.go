package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/repository"
)

type ProUsecase interface {
	Register(newPro *model.Product) error
	FindAll() ([]model.Product, error)
	FindOne(id string) (*model.Product, error)
	Edit(updatedPro *model.Product) error
	Unreg(id string) error
}

type proUsecase struct {
	repo repository.ProRepository
}

func (u *proUsecase) Register(newPro *model.Product) error {
	return u.repo.Create(newPro)
}

func (u *proUsecase) FindAll() ([]model.Product, error) {
	return u.repo.FindAll()
}

func (u *proUsecase) FindOne(id string) (*model.Product, error) {
	return u.repo.FindOne(id)
}

func (u *proUsecase) Edit(updatedPro *model.Product) error {
	return u.repo.Update(updatedPro)
}

func (u *proUsecase) Unreg(id string) error {
	return u.repo.Delete(id)
}

func NewProUsecase(repo repository.ProRepository) ProUsecase {
	return &proUsecase{
		repo,
	}
}
