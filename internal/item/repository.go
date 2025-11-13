package item

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(item *Item) error
	GetAll() ([]Item, error)
	GetById(id uint) (*Item, error)
	Update(item *Item) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *Item) error {
	return r.db.Create(item).Error
}

func (r *repository) GetAll() ([]Item, error) {
	var items []Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *repository) GetById(id uint) (*Item, error) {
	var item Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *repository) Update(item *Item) error {
	return r.db.Save(item).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Item{}, id).Error
}
