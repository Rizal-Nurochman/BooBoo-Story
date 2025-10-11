package categories

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	// Hierarchical endpoints
	GetChildren(c *gin.Context)
	GetParent(c *gin.Context)
	CreateSubCategory(c *gin.Context)
}

type handler struct {
	service Service
}


func NewHandler(service Service) Handler {
	return &handler{service: service}
}


// Create implements Handler.
func (h *handler) Create(c *gin.Context) {
	panic("unimplemented")
}

// CreateSubCategory implements Handler.
func (h *handler) CreateSubCategory(c *gin.Context) {
	panic("unimplemented")
}

// Delete implements Handler.
func (h *handler) Delete(c *gin.Context) {
	panic("unimplemented")
}

// GetAll implements Handler.
func (h *handler) GetAll(c *gin.Context) {
	panic("unimplemented")
}

// GetByID implements Handler.
func (h *handler) GetByID(c *gin.Context) {
	panic("unimplemented")
}

// GetChildren implements Handler.
func (h *handler) GetChildren(c *gin.Context) {
	panic("unimplemented")
}

// GetParent implements Handler.
func (h *handler) GetParent(c *gin.Context) {
	panic("unimplemented")
}

// Update implements Handler.
func (h *handler) Update(c *gin.Context) {
	panic("unimplemented")
}

