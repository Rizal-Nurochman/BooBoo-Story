package progresses

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
	var input validations.CreateProgressInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input tidak valid", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("user_id").(uint)

	newProgress, err := h.service.Create(userID, input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal menyimpan progres", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusCreated, "success", "Progres berhasil disimpan", newProgress, nil, nil)
}

func (h *handler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	q := c.DefaultQuery("q", "")
	includes := c.QueryArray("include")

	page, _ := utils.ToUint(pageStr)
	limit, _ := utils.ToUint(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	userID := c.MustGet("user_id").(uint)

	progresses, total, err := h.service.GetAll(userID, includes, int(page), int(limit), q)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal mendapatkan data progres", nil, err.Error(), nil)
		return
	}

	meta := gin.H{"total": total, "page": page, "limit": limit}
	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan data progres", progresses, nil, meta)
}

func (h *handler) GetByID(c *gin.Context) {
	progressIDStr := c.Param("id")
	progressID, err := utils.ToUint(progressIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID progres tidak valid", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("user_id").(uint)
	includes := c.QueryArray("include")

	progress, err := h.service.GetByID(userID, uint(progressID), includes)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Progres tidak ditemukan atau bukan milik Anda", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan progres", progress, nil, nil)
}

func (h *handler) Update(c *gin.Context) {
	progressIDStr := c.Param("id")
	progressID, err := utils.ToUint(progressIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID progres tidak valid", nil, err.Error(), nil)
		return
	}

	var input validations.UpdateProgressInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input tidak valid", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("user_id").(uint)

	updatedProgress, err := h.service.Update(userID, uint(progressID), input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal memperbarui progres", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Progres berhasil diperbarui", updatedProgress, nil, nil)
}

func (h *handler) Delete(c *gin.Context) {
	progressIDStr := c.Param("id")
	progressID, err := utils.ToUint(progressIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID progres tidak valid", nil, err.Error(), nil)
		return
	}

	userID := c.MustGet("user_id").(uint)

	err = h.service.Delete(userID, uint(progressID))
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal menghapus progres atau bukan milik Anda", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Progres berhasil dihapus", nil, nil, nil)
}