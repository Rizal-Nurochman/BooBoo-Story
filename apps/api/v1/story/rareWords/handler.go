package rarewords

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



// Create implements Handler.
func (h *handler) Create(c *gin.Context) {
	var input validations.StoryContentRareWordCreate
	err := c.ShouldBindBodyWithJSON((&input))

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	rareWord, err := h.service.Create(input)

	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to create rare word", nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusCreated,"success","Rare word created successfully", rareWord, nil, nil)
}

// Delete implements Handler.
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	
	err := h.service.Delete(idStr)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to delete rare word", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Rare word deleted successfully", nil, nil, nil)
}

// GetAll implements Handler.
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

// Update implements Handler.
func (h *handler) Update(c *gin.Context) {
	panic("unimplemented")
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}
