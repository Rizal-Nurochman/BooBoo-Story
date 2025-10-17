package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/BooBooStory/config"
	validations "github.com/BooBooStory/config/validations"
	"github.com/BooBooStory/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	googleAPI "google.golang.org/api/oauth2/v2"
)

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Me(c *gin.Context)
	GoogleLoginHandler(c *gin.Context)
	GoogleCallbackHandler(c *gin.Context)
	RequestPasswordReset(c *gin.Context)
	VerifyAndResetPassword(c *gin.Context)
}

type handler struct {
	service     Service
	oauthConfig *oauth2.Config
}

func NewHandler(service Service, oauthConfig *oauth2.Config) Handler {
	return &handler{
		service:     service,
		oauthConfig: oauthConfig,
	}
}

func (h *handler) Register(c *gin.Context) {
	var req validations.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "invalid request", nil, err.Error(), nil)
		return
	}
	user, err := h.service.Register(req)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", err.Error(), nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusCreated, "success", "user registered successfully", gin.H{
		"id":    fmt.Sprintf("%d", user.ID),
		"name":  user.Name,
		"email": user.Email,
	}, nil, nil)
}

func (h *handler) Me(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		utils.JSON(c, http.StatusUnauthorized, "error", "User not authenticated", nil, errors.New("unauthorized"), nil)
		return
	}
	id, ok := userID.(uint)
	if !ok {
		utils.JSON(c, http.StatusInternalServerError, "error", "Invalid user ID type", nil, errors.New("invalid user_id type"), nil)
		return
	}
	user, err := h.service.Me(id)
	if err != nil {
		utils.JSON(c, http.StatusNotFound, "error", err.Error(), nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusOK, "success", "User fetched successfully", gin.H{
		"id":     fmt.Sprintf("%d", user.ID),
		"name":   user.Name,
		"email":  user.Email,
		"role":   user.Role,
		"avatar": user.Avatar,
		"points": user.Points,
		"streak": user.Streak,
		"level":  user.Level,
	}, nil, nil)
}

func (h *handler) Login(c *gin.Context) {
	var req validations.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "invalid request", nil, err.Error(), nil)
		return
	}
	user, token, err := h.service.Login(req)
	if err != nil {
		utils.JSON(c, http.StatusUnauthorized, "error", err.Error(), nil, err.Error(), nil)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   30 * 24 * 3600,
		HttpOnly: true,
		Secure:   config.Envs.APP_ENV == "production", // ubah ke true jika FE sudah pakai HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	utils.JSON(c, http.StatusOK, "success", "login successful", gin.H{
		"id":    fmt.Sprintf("%d", user.ID),
		"name":  user.Name,
		"email": user.Email,
	}, nil, nil)
}

func (h *handler) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   config.Envs.APP_ENV == "production", // ubah ke true jika FE sudah pakai HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	utils.JSON(c, http.StatusOK, "success", "logout successful", nil, nil, nil)
}

func (h *handler) GoogleLoginHandler(c *gin.Context) {
	state := "random-string-for-csrf-prevention"
	url := h.oauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *handler) GoogleCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		utils.JSON(c, http.StatusBadRequest, "error", "authorization code not provided", nil, nil, nil)
		return
	}
	token, err := h.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "failed to exchange code", nil, err.Error(), nil)
		return
	}
	client := h.oauthConfig.Client(context.Background(), token)
	service, err := googleAPI.New(client)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "failed to create Google client", nil, err.Error(), nil)
		return
	}
	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "failed to get user info", nil, err.Error(), nil)
		return
	}
	user, appToken, err := h.service.LoginOrRegisterWithGoogle(userInfo)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "failed to login or register with Google", nil, err.Error(), nil)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    appToken,
		Path:     "/",
		MaxAge:   30 * 24 * 3600,
		HttpOnly: true,
		Secure:   config.Envs.APP_ENV == "production",
		SameSite: http.SameSiteLaxMode,
	})

	utils.JSON(c, http.StatusOK, "success", "successfully logged in with Google", gin.H{
		"id":    fmt.Sprintf("%d", user.ID),
		"name":  user.Name,
		"email": user.Email,
	}, nil, nil)
}

func (h *handler) RequestPasswordReset(c *gin.Context) {
	var req validations.ForgotPasswordInput
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}
	err := h.service.RequestPasswordReset(req.Email)
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "error", "An internal error occurred", nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusOK, "success", "If an account with that email exists, a password reset link has been sent.", nil, nil, nil)
}

func (h *handler) VerifyAndResetPassword(c *gin.Context) {
	var req validations.ResetPasswordInput
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", "Invalid request body", nil, err.Error(), nil)
		return
	}
	err := h.service.VerifyAndResetPassword(req)
	if err != nil {
		utils.JSON(c, http.StatusBadRequest, "error", err.Error(), nil, err.Error(), nil)
		return
	}
	utils.JSON(c, http.StatusOK, "success", "Password has been successfully reset.", nil, nil, nil)
}
