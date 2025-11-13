package item

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(item *Item) error
	GetAll(filters ItemFilters) ([]Item, error)
	GetById(id uint) (*Item, error)
	Update(item *Item) error
	Delete(id uint) error
}

type ItemFilters struct {
	CategoryID     *uint
	Expired        *bool
	ExpiringWithin *int // days
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

func (r *repository) GetAll(filters ItemFilters) ([]Item, error) {
	var items []Item
	query := r.db.Model(&Item{})

	// filter by category
	if filters.CategoryID != nil {
		query = query.Where("category_id = ?", *filters.CategoryID)
	}

	// filter by expired
	if filters.Expired != nil {
		now := time.Now()
		if *filters.Expired {
			query = query.Where("expiry_date < ?", now)
		} else {
			query = query.Where("expiry_date >= ?", now)
		}
	}

	// filter by expiring within N days
	if filters.ExpiringWithin != nil {
		cutoff := time.Now().AddDate(0, 0, *filters.ExpiringWithin)
		query = query.Where("expiry_date BETWEEN ? AND ?", time.Now(), cutoff)
	}

	if err := query.Find(&items).Error; err != nil {
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
