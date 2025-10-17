package quizs

import (
	"net/http"

	"github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service // Menggunakan interface dari service.go
}

// Handler interface mendefinisikan semua fungsi yang akan diekspos
type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Delete(c *gin.Context)
}

// NewHandler adalah constructor untuk handler
func NewHandler(service Service) *handler {
	return &handler{service}
}

// Create menangani pembuatan kuis baru
func (h *handler) Create(c *gin.Context) {
	var input validations.CreateQuizInput

	// Bind body JSON ke DTO
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input tidak valid", nil, err.Error(), nil)
		return
	}

	// Panggil service untuk membuat kuis
	newQuiz, err := h.service.Create(input)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal membuat kuis", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusCreated, "success", "Kuis berhasil dibuat", newQuiz, nil, nil)
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

	quizzes, total, err := h.service.GetAll(includes, int(page), int(limit), q)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal mendapatkan data kuis", nil, err.Error(), nil)
		return
	}

	meta := gin.H{"total": total, "page": page, "limit": limit}
	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan data kuis", quizzes, nil, meta)
}

func (h *handler) GetByID(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := utils.ToUint(quizIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID kuis tidak valid", nil, err.Error(), nil)
		return
	}

	quiz, err := h.service.GetByID(uint(quizID))
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", "Kuis tidak ditemukan", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Berhasil mendapatkan kuis", quiz, nil, nil)
}

// Delete menangani penghapusan kuis
func (h *handler) Delete(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := utils.ToUint(quizIDStr)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "ID kuis tidak valid", nil, err.Error(), nil)
		return
	}

	err = h.service.Delete(uint(quizID))
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Gagal menghapus kuis", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "Kuis berhasil dihapus", nil, nil, nil)
}