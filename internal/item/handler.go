package item

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("", h.Create)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

// Create godoc
// @Summary Create a new item
// @Description Create an item
// @Tags items
// @Accept json
// @Produce json
// @Param item body CreateItemRequest true "Item to create"
// @Success 201 {object} item.Item
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items [post]
func (h *Handler) Create(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var expiry *time.Time
	if req.ExpiryDate != "" {
		parsed, err := time.Parse("2006-01-02", req.ExpiryDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiry_date format, expected YYYY-MM-DD"})
			return
		}
		expiry = &parsed
	}

	item := Item{
		Name:       req.Name,
		Quantity:   req.Quantity,
		Unit:       req.Unit,
		ExpiryDate: expiry,
		CategoryID: req.CategoryID,
		Notes:      req.Notes,
	}

	if err := h.service.Create(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetAll godoc
// @Summary Get all items
// @Description Retrieve a list of all items, optionally filtered by category, expiry, or expiring soon
// @Tags items
// @Produce json
// @Param category_id query int false "Filter by category ID"
// @Param expired query bool false "Filter expired items only"
// @Param expiring_within query int false "Filter items expiring within N days"
// @Success 200 {array} item.Item
// @Failure 500 {object} map[string]string
// @Router /items [get]
func (h *Handler) GetAll(c *gin.Context) {
	var filters ItemFilters

	if categoryIDStr := c.Query("category_id"); categoryIDStr != "" {
		if id, err := strconv.Atoi(categoryIDStr); err == nil {
			idUint := uint(id)
			filters.CategoryID = &idUint
		}
	}

	if expiredStr := c.Query("expired"); expiredStr != "" {
		expired := expiredStr == "true"
		filters.Expired = &expired
	}

	if expiringWithinStr := c.Query("expiring_within"); expiringWithinStr != "" {
		if days, err := strconv.Atoi(expiringWithinStr); err == nil {
			filters.ExpiringWithin = &days
		}
	}

	items, err := h.service.GetAll(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetByID godoc
// @Summary Get an item by ID
// @Description Retrieve a single item by its ID
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} item.Item
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Update godoc
// @Summary Update an item
// @Description Update the attributes of an item by ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body UpdateItemRequest true "Item fields to update"
// @Success 200 {object} item.Item
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	if req.Unit != nil {
		item.Unit = *req.Unit
	}
	if req.Notes != nil {
		item.Notes = *req.Notes
	}
	if req.CategoryID != nil {
		item.CategoryID = req.CategoryID
	}
	if req.ExpiryDate != nil && *req.ExpiryDate != "" {
		parsed, err := time.Parse("2006-01-02", *req.ExpiryDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiry_date format, expected YYYY-MM-DD"})
			return
		}
		item.ExpiryDate = &parsed
	}

	if err := h.service.Update(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Delete godoc
// @Summary Delete an item
// @Description Delete an item by its ID
// @Tags items
// @Param id path int true "Item ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
