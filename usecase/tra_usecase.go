package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/repository"
)

type TraUsecase interface {
	Register(newPro *model.Bill) error
	FindOne(id string) (*model.Bill, error)
}

type traUsecase struct {
	repo repository.TraRepository
}

func (u *traUsecase) Register(newTra *model.Bill) error {
	return u.repo.Create(newTra)
}

func (u *traUsecase) FindOne(id string) (*model.Bill, error) {
	return u.repo.FindOne(id)
}

func NewTraUsecase(repo repository.TraRepository) TraUsecase {
	return &traUsecase{
		repo,
	}
}
