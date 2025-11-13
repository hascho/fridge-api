package item

import "errors"

type Service interface {
	Create(item *Item) error
	GetAll() ([]Item, error)
	GetById(id uint) (*Item, error)
	Update(item *Item) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(item *Item) error {
	if item.Name == "" {
		return errors.New("item name is required")
	}
	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	return s.repo.Create(item)
}

func (s *service) GetAll() ([]Item, error) {
	return s.repo.GetAll()
}

func (s *service) GetById(id uint) (*Item, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetById(id)
}

func (s *service) Update(item *Item) error {
	if item.ID == 0 {
		return errors.New("item ID is required for update")
	}
	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	return s.repo.Update(item)
}

func (s *service) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(id)
}
