package item

type CreateItemRequest struct {
	Name       string  `json:"name" binding:"required,max=100"`
	Quantity   float64 `json:"quantity" binding:"required,gte=0"`
	Unit       string  `json:"unit" binding:"required,max=20"`
	ExpiryDate string  `json:"expiry_date,omitempty" example:"2025-12-31"`
	CategoryID *uint   `json:"category_id,omitempty"`
	Notes      string  `json:"notes,omitempty"`
}

type UpdateItemRequest struct {
	Name       *string  `json:"name,omitempty" binding:"omitempty,max=100"`
	Quantity   *float64 `json:"quantity,omitempty" binding:"omitempty,gte=0"`
	Unit       *string  `json:"unit,omitempty" binding:"omitempty,max=20"`
	ExpiryDate *string  `json:"expiry_date,omitempty"`
	CategoryID *uint    `json:"category_id,omitempty"`
	Notes      *string  `json:"notes,omitempty"`
}
