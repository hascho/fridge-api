package category

import "errors"

type Service interface {
	Create(category *Category) error
	GetAll() ([]Category, error)
	GetById(id uint) (*Category, error)
	Update(category *Category) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(category *Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	return s.repo.Create(category)
}

func (s *service) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *service) GetById(id uint) (*Category, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetById(id)
}

func (s *service) Update(category *Category) error {
	if category.ID == 0 {
		return errors.New("category ID is required for update")
	}
	if category.Name == "" {
		return errors.New("category name is required")
	}
	return s.repo.Update(category)
}

func (s *service) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(id)
}
