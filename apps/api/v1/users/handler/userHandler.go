package handler

import (
	"net/http"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/utils"
	"github.com/BooBooStory/v1/users/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{service}
}

// Endpoint ketika user daftar akun
func (h *userHandler) RegisterHandler(c *gin.Context) {
	var inputDataByUser models.UserRegister
	err := c.ShouldBindJSON(&inputDataByUser)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input data ada yang tidak sesuai atau masih kosong", nil, err.Error(), nil)
		return
	}

	newUser, err := h.service.DaftarAkun(inputDataByUser)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal daftar akun", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "berhasil daftar akun", newUser, nil, nil)
}
