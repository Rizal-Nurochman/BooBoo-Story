package bookmarks

import (
	"net/http"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type bookmarkHandler struct {
	service Service
}

type CreateBookmarkInput struct {
	StoryID uint `json:"story_id" binding:"required"`
}

func NewBookmarkHandler(service Service) *bookmarkHandler {
	return &bookmarkHandler{service}
}

func (h *bookmarkHandler) CreateBookmark(c *gin.Context) {
	var input CreateBookmarkInput
	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil{
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("user_id").(uint)
	bookmark, err := h.service.CreateBookmark(userID, input.StoryID)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to create bookmark", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Bookmark created successfully", bookmark, nil, nil)
}

func (h *bookmarkHandler) GetBookmarks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	q := c.DefaultQuery("q", "")
	includes := c.QueryArray("include")

	page, err := utils.ToUint(pageStr)
	if err != nil || page < 1 || page == 0 {
		page = 1
	}

	limit, err := utils.ToUint(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	bookmarks, total, err := h.service.GetAllBookmarks(userID, includes, int(page), int(limit), q)
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

func (h *bookmarkHandler) GetBookmarkByID(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	bookmarkID := c.Param("id")
	includes := c.QueryArray("include")

	bookmark, err := h.service.GetBookmarkByID(userID, bookmarkID, includes)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal mendapatkan bookmark", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan bookmark", bookmark, nil, nil)
}

func (h *bookmarkHandler) DeleteBookmark(c *gin.Context) {
	bookmarkID := c.Param("id")

	err := h.service.DeleteBookmark(bookmarkID)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal menghapus bookmark", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil menghapus bookmark", nil, nil, nil)
}