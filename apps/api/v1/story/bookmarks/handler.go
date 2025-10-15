package bookmarks

import (
	"net/http"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type bookmarkHandler struct {
	service Service
}

func NewBookmarkHandler(service Service) *bookmarkHandler {
	return &bookmarkHandler{service}
}

func (h *bookmarkHandler) CreateBookmark(c *gin.Context) {
	var input struct {
		StoryID string `json:"story_id" binding:"required"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input tidak valid", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("currentUser").(string)

	newBookmark, err := h.service.CreateBookmark(userID, input.StoryID)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal membuat bookmark", nil, err.Error(), nil)
		return
	}
	
	utils.JSON(c, http.StatusOK, "success", "Berhasil membuat bookmark", newBookmark, nil, nil)
}

func (h *bookmarkHandler) GetBookmarks(c *gin.Context) {
	userID := c.MustGet("currentUser").(string)

	bookmarks, err := h.service.GetBookmarks(userID)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal mendapatkan bookmark", nil, err.Error(), nil)
		return
	}
	
	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan semua bookmark", bookmarks, nil, nil)
}

func (h *bookmarkHandler) GetBookmarkByID(c *gin.Context) {
	userID := c.MustGet("currentUser").(string)
	id := c.Param("id")

	bookmark, err := h.service.GetBookmarkByID(userID, id)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Bookmark tidak ditemukan", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan bookmark", bookmark, nil, nil)
}

func (h *bookmarkHandler) DeleteBookmark(c *gin.Context) {
	userID := c.MustGet("currentUser").(string)
	id := c.Param("id")

	_, err := h.service.DeleteBookmark(userID, id)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal menghapus bookmark", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil menghapus bookmark", nil, nil, nil)
}