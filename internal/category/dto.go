package category

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,max=50"`
}

type UpdateCategoryRequest struct {
	Name *string `json:"name,omitempty" binding:"omitempty,max=50"`
}
