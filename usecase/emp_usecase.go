package usecase

import (
	"go-laundry/model"
	"go-laundry/repository"
)

type EmpUsecase interface {
	Register(newEmp *model.Employee) error
	FindAll() ([]model.Employee, error)
	FindOne(id string) (*model.Employee, error)
	Edit(updatedEmp *model.Employee) error
	Unreg(id string) error
}

type empUsecase struct {
	repo repository.EmpRepository
}

func (u *empUsecase) Register(newEmp *model.Employee) error {
	return u.repo.Create(newEmp)
}

func (u *empUsecase) FindAll() ([]model.Employee, error) {
	return u.repo.FindAll()
}

func (u *empUsecase) FindOne(id string) (*model.Employee, error) {
	return u.repo.FindOne(id)
}

func (u *empUsecase) Edit(updatedEmp *model.Employee) error {
	return u.repo.Update(updatedEmp)
}

func (u *empUsecase) Unreg(id string) error {
	return u.repo.Delete(id)
}

func NewEmpUsecase(repo repository.EmpRepository) EmpUsecase {
	return &empUsecase{
		repo,
	}
}
