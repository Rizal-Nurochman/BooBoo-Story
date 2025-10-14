package contents

import (
	"net/http"
	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

// ✅ Create
func (h *handler) Create(c *gin.Context) {
	var input validations.StoryContentCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	rareWord, err := h.service.Create(input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to create rare word", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusCreated, "success", "Rare word created successfully", rareWord, nil, nil)
}

// ✅ GetAll (by story_content_id)
func (h *handler) GetAll(c *gin.Context) {
	storyContentIDStr := c.Query("story_content_id")
	if storyContentIDStr == "" {
		utils.JSON(c, http.StatusBadRequest, "error", "story_content_id query parameter is required", nil, "story_content_id is required", nil)
		return
	}

	rareWords, err := h.service.GetAll(storyContentIDStr)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to fetch rare words", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Rare words fetched successfully", rareWords, nil, nil)
}

// ✅ Update
func (h *handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	var input validations.StoryContentUpdate

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	updatedRareWord, err := h.service.Update(idStr, input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to update rare word", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Rare word updated successfully", updatedRareWord, nil, nil)
}

// ✅ Delete
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		utils.JSON(c, http.StatusBadRequest, "error", "id parameter is required", nil, "id is required", nil)
		return
	}

	if err := h.service.Delete(idStr); err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to delete rare word", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Rare word deleted successfully", nil, nil, nil)
}