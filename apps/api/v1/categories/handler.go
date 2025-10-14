package categories

import (
	"net/http"
	"strconv"

	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
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
	var input validations.CategoryCreate
	err := c.ShouldBindBodyWithJSON((&input))

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	category, err := h.service.Create(input)

	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to create category", nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusCreated,"success","Category created successfully", category, nil, nil)
}

// CreateSubCategory implements Handler.
func (h *handler) CreateSubCategory(c *gin.Context) {
	parentIDStr := c.Param("parent_id")
	parentID, err := utils.ToUint(parentIDStr)

	if err !=nil{
		utils.JSON(c, 400, "error","Invalid parent ID", nil, err.Error(), nil)
		return
	}

	var input validations.SubCategoryCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	category, err := h.service.CreateSubCategory(uint(parentID), input)

	if err != nil{
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to create sub category", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusCreated, "success", "Subcategory created successfully", category, err.Error(), nil)
}

// Delete implements Handler.
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.ToUint(idStr)

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error","invalid category ID",nil, err.Error(),nil)
		return
	}

	if err := h.service.Delete(uint(id));err != nil{
		utils.JSON(c, http.StatusInternalServerError, "error","Failed to delete category", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success","Category deleted successfully", nil, err.Error(), nil)
}

// GetAll implements Handler.
func (h *handler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	withChildrenStr := c.DefaultQuery("with_children", "false")
	q := c.DefaultQuery("q", "")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	withChildren := withChildrenStr == "true"

		categories, total, err := h.service.GetAll(withChildren, page, limit, q)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to fetch categories", nil, err.Error(), nil)
		return
	}

	meta := gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	utils.JSON(c, http.StatusOK, "success", "Categories fetched successfully", categories, nil, meta)
}

// GetByID implements Handler.
func (h *handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.ToUint(idStr)

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error","Invalid category id",nil, err.Error(), nil)
		return
	}

	withChildrenStr := c.DefaultQuery("with_children","false")
	withChildren := withChildrenStr == "true"

	category, err :=h.service.GetByID(uint(id), withChildren)

	if err != nil{
		utils.JSON(c, http.StatusNotFound,"error","Cetgory not found", nil, err.Error(), nil)
	}

	utils.JSON(c, http.StatusOK, "success", "Category fetched successfully", category, nil, nil)
}

// GetChildren implements Handler.
func (h *handler) GetChildren(c *gin.Context) {
	parentIDStr := c.Param("parent_id")

	parentID, err := utils.ToUint(parentIDStr)

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error","ivalid parent ID", nil,err.Error(), nil)
		return
	}

	children, err := h.service.GetChildren(uint(parentID))

	if err != nil{
		utils.JSON(c, http.StatusInternalServerError, "error", "Errror getting subCategories", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Sucessfully fetched subCategories", children, nil, nil)
}

// GetParent implements Handler.
func (h *handler) GetParent(c *gin.Context) {
	childIDStr := c.Param("child_id")
	childID, err := strconv.ParseUint(childIDStr, 10, 64)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid child_id", nil, err.Error(), nil)
		return
	}

	parent, err := h.service.GetParent(uint(childID))
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to fetch parent", nil, err.Error(), nil)
		return
	}

	if parent == nil {
		utils.JSON(c, http.StatusOK, "success", "This category has no parent", nil, nil, nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Parent fetched successfully", parent, nil, nil)
}

// Update implements Handler.
func (h *handler) Update(c *gin.Context) {
idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid category ID", nil, err.Error(), nil)
		return
	}

	var input validations.CategoryUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	category, err := h.service.Update(uint(id), input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to update category", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Category updated successfully", category, nil, nil)
}

