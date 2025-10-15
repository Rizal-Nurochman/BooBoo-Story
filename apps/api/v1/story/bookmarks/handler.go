package bookmarks

import (
	"net/http"
	"strconv"

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

	userID := c.MustGet("user_id").(string)

	newBookmark, err := h.service.CreateBookmark(userID, input.StoryID)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal membuat bookmark", nil, err.Error(), nil)
		return
	}
	
	utils.JSON(c, http.StatusOK, "success", "Berhasil membuat bookmark", newBookmark, nil, nil)
}

func (h *bookmarkHandler) GetBookmarks(c *gin.Context) {
	userID := c.MustGet("currentUser").(string)

	// Ambil query param 'page' dan 'limit' dari URL
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	
	bookmarks, total, err := h.service.GetBookmarks(userID, page, limit)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal mendapatkan bookmark", nil, err.Error(), nil)
		return
	}

	// Buat metadata untuk respons paginasi
	meta := gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan semua bookmark", bookmarks, nil, meta)
}

func (h *bookmarkHandler) DeleteBookmark(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	id := c.Param("id")

	_, err := h.service.DeleteBookmark(userID, id)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal menghapus bookmark", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil menghapus bookmark", nil, nil, nil)
}