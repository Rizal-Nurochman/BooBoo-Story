package creators

import (
	"net/http"

	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) Create(c *gin.Context) {
	var input validations.CreatorCreate
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, 400, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	newCreator, err := h.service.Create(input)
	if err != nil {
		utils.JSON(c, 500, "error", "Failed to create creator", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, 201, "success", "Creator created successfully", newCreator, nil, nil)
}

func (h *handler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	q := c.DefaultQuery("q", "")

	page, err := utils.ToUint(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := utils.ToUint(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	creators, total, err := h.service.GetAll(int(page), int(limit), q)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to get creators data", nil, err.Error(), nil)
		return
	}

	meta := gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	utils.JSON(c, http.StatusOK, "success", "Success to get creators data", creators, nil, meta)
}

func (h *handler) GetByID(c *gin.Context) {
	creatorIDStr := c.Param("id")
	creatorID, err := utils.ToUint(creatorIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID creator not valid", nil, err.Error(), nil)
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	q := c.DefaultQuery("q", "")
	includes := c.QueryArray("include")

	page, err := utils.ToUint(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := utils.ToUint(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	creator, err := h.service.GetByID(uint(creatorID), includes, int(page), int(limit), q)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Failed to get creator data", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Success to get creator data", creator, nil, nil)
}

func (h *handler) Update(c *gin.Context) {
	creatorIDStr := c.Param("id")
	creatorID, err := utils.ToUint(creatorIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID creator not valid", nil, err.Error(), nil)
		return
	}

	var input validations.CreatorUpdate
	err = c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	updatedCreator, err := h.service.Update(uint(creatorID), input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to update creator", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Creator updated successfully", updatedCreator, nil, nil)
}

func (h *handler) Delete(c *gin.Context) {
	creatorIDStr := c.Param("id")
	creatorID, err := utils.ToUint(creatorIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID creator not valid", nil, err.Error(), nil)
		return
	}

	err = h.service.Delete(uint(creatorID))
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to delete creator", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Creator deleted successfully", nil, nil, nil)
}