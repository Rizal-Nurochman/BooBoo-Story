package usershandler

import (
	"context"
	"net/http"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/utils"
	"github.com/BooBooStory/v1/users/usersservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	googleAPI "google.golang.org/api/oauth2/v2"
)

type userHandler struct {
	service usersservice.UserService
}

func NewUserHandler(service usersservice.UserService) *userHandler {
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

// Endpoint ketika user login
func (h *userHandler) LoginHandler(c *gin.Context) {
	var inputDataByUser models.UserLogin
	err := c.ShouldBindJSON(&inputDataByUser)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Input data ada yang tidak sesuai atau masih kosong", nil, err.Error(), nil)
		return
	}

	userDitemukan, token, err := h.service.LoginAkun(inputDataByUser)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Gagal login", nil, err.Error(), nil)
		return
	}

	utils.JSON(c, http.StatusOK, "success", "berhasil login", gin.H{
		"user":  userDitemukan,
		"token": token,
	}, nil, nil)
}

// Endpoint baru untuk memulai login dengan Google
func (h *userHandler) GoogleLoginHandler(c *gin.Context, oauthConfig *oauth2.Config) {
	// Generate URL untuk otorisasi Google.
	// State digunakan untuk mencegah serangan CSRF.
	state := "random-string-for-csrf-prevention"
	url := oauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Endpoint baru untuk menangani callback dari Google
func (h *userHandler) GoogleCallbackHandler(c *gin.Context, oauthConfig *oauth2.Config) {
	// 1. Dapatkan 'code' dari query parameter
	code := c.Query("code")
	if code == "" {
		utils.JSON(c, http.StatusBadRequest, "error", "Authorization code not provided", nil, nil, nil)
		return
	}

	// 2. Tukarkan 'code' dengan token dari Google
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to exchange code", nil, err.Error(), nil)
		return
	}

	// 3. Gunakan token untuk mendapatkan info user dari Google API
	client := oauthConfig.Client(context.Background(), token)
	service, err := googleAPI.New(client)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to create Google service client", nil, err.Error(), nil)
		return
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "Failed to get user info", nil, err.Error(), nil)
		return
	}

	// 4. Panggil service untuk login atau registrasi
	user, appToken, err := h.service.LoginOrRegisterWithGoogle(userInfo)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Failed to login or register with Google", nil, err.Error(), nil)
		return
	}

	// 5. Kirim response sukses dengan JWT token aplikasi Anda
	utils.JSON(c, http.StatusOK, "success", "Successfully logged in with Google", gin.H{
		"user":  user,
		"token": appToken,
	}, nil, nil)
}