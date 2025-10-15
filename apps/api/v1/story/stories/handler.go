package stories

import (
	"net/http"

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
}

type handler struct {
	service Service
}

func (h *handler) Create(c *gin.Context) {
	var input validations.StoryCreate
	err := c.ShouldBindBodyWithJSON((&input))

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	story, err := h.service.Create(input)

	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to create story", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusCreated,"success","Story created successfully", story, nil, nil)
}

func (h *handler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	q := c.DefaultQuery("q", "")
	includes := c.QueryArray("include")
	
	page, err := utils.ToUint(pageStr)
	if err != nil || page == 0 {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid page parameter", nil, "page must be a positive integer", nil)
		return
	}
	
	limit, err := utils.ToUint(limitStr)

	if err != nil || limit == 0 {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid limit parameter", nil, "limit must be a positive integer", nil)
		return
	}
	stories, total, err := h.service.GetAll(includes, int(page), int(limit), q)
	
	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to fetch stories", nil, err.Error(), nil)
		return
	}

	meta := map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": total,
	}
	
	utils.JSON(c, http.StatusOK,"success","Stories fetched successfully", stories, nil, meta)
}

func (h *handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	includes := c.QueryArray("include")

	story, err := h.service.GetByID(idStr, includes)
	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to fetch story", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK,"success","Story fetched successfully", story, nil, nil)
}

func (h *handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	
	var input validations.StoryUpdate
	err := c.ShouldBindBodyWithJSON((&input))
	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	updatedStory, err := h.service.Update(idStr, input)
	if err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to update story", nil, err.Error(), nil)
		return
	}
	
	utils.JSON(c, http.StatusOK,"success","Story updated successfully", updatedStory, nil, nil)
}

func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	if err := h.service.Delete(idStr);err != nil{
		utils.JSON(c, http.StatusInternalServerError,"error","Failed to delete story", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK,"success","Story deleted successfully", nil, nil, nil)
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
