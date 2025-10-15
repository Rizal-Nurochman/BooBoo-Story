package creators

import (
	"net/http"
	"strconv"

	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type creatorHandler struct {
	service Service
}

func NewCreatorHandler(service Service) *creatorHandler {
	return &creatorHandler{service}
}

// Handler untuk kreator mengambil profilnya sendiri
func (h *creatorHandler) GetMyProfile(c *gin.Context) {
	// Diambil dari middleware auth. Kita asumsikan middleware menyimpan ID user sebagai uint
	CreatorID := c.MustGet("user_id").(uint) 

	creator, err := h.service.GetCreatorProfile(CreatorID)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Profil kreator tidak ditemukan", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan profil kreator", creator, nil, nil)
}

// Handler untuk kreator memperbarui profilnya sendiri
func (h *creatorHandler) UpdateMyProfile(c *gin.Context) {
	var input UpdateCreatorInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input tidak valid", nil, err.Error(), nil)
		return
	}

	CreatorID := c.MustGet("user_id").(uint)

	updatedCreator, err := h.service.UpdateCreatorProfile(CreatorID, input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal memperbarui profil", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil memperbarui profil", updatedCreator, nil, nil)
}

// --- Handler Baru untuk Admin ---

// Handler untuk mendapatkan profil kreator berdasarkan ID kreator
func (h *creatorHandler) GetCreatorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID tidak valid", nil, err.Error(), nil)
		return
	}

	creator, err := h.service.GetCreatorProfileByID(uint(id))
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Kreator tidak ditemukan", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan profil kreator", creator, nil, nil)
}

// Handler untuk menghapus kreator berdasarkan ID kreator
func (h *creatorHandler) DeleteCreatorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID tidak valid", nil, err.Error(), nil)
		return
	}

	_, err = h.service.DeleteCreatorByID(uint(id))
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Gagal menghapus kreator", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil menghapus kreator", nil, nil, nil)
}