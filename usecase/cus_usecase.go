package usecase

import (
	"go-laundry/model"
	"go-laundry/repository"
)

type CusUsecase interface {
	Register(newCus *model.Customer) error
	FindAll() ([]model.Customer, error)
	FindOne(id string) (*model.Customer, error)
	Edit(updatedCus *model.Customer) error
	Unreg(id string) error
}

type cusUsecase struct {
	repo repository.CusRepository
}

func (u *cusUsecase) Register(newCus *model.Customer) error {
	return u.repo.Create(newCus)
}

func (u *cusUsecase) FindAll() ([]model.Customer, error) {
	return u.repo.FindAll()
}

func (u *cusUsecase) FindOne(id string) (*model.Customer, error) {
	return u.repo.FindOne(id)
}

func (u *cusUsecase) Edit(updatedCus *model.Customer) error {
	return u.repo.Update(updatedCus)
}

func (u *cusUsecase) Unreg(id string) error {
	return u.repo.Delete(id)
}

func NewCusUsecase(repo repository.CusRepository) CusUsecase {
	return &cusUsecase{
		repo,
	}
}
