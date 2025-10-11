package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/BooBooStory/config"
	"github.com/BooBooStory/config/models"
	validations "github.com/BooBooStory/config/schemas"
	email "github.com/BooBooStory/utils"
	"github.com/BooBooStory/v1/users"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/oauth2/v2"
	"gorm.io/gorm"
)

type Service interface {
	Register(user validations.UserRegister) (*models.User, error)
	Login(user validations.UserLogin) (*models.User, string, error)
	LoginOrRegisterWithGoogle(googleUserInfo *oauth2.Userinfo) (*models.User, string, error)
	Logout() (string, error)
	RequestPasswordReset(email string) error
	VerifyAndResetPassword(input validations.ResetPasswordInput) error
}

type service struct {
	userRepo users.Repository
	emailSvc email.Service
}

func NewService(userRepo users.Repository, emailSvc email.Service) Service {
	return &service{userRepo, emailSvc}
}

func (s *service) generateTokenJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Envs.JwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func generateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *service) Register(user validations.UserRegister) (*models.User, error) {
	findUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if findUser != nil {
		return nil, errors.New("email already registered")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	inputUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passwordHash),
	}
	newUser, err := s.userRepo.CreateUser(inputUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *service) Login(user validations.UserLogin) (*models.User, string, error) {
	findUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("user not found, please register first")
		}
		return nil, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}
	token, err := s.generateTokenJWT(findUser.ID)
	if err != nil {
		return nil, "", err
	}
	return findUser, token, nil
}

func (s *service) LoginOrRegisterWithGoogle(googleUserInfo *oauth2.Userinfo) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(googleUserInfo.Email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := &models.User{
			Name:     googleUserInfo.Name,
			Email:    googleUserInfo.Email,
			Password: "",
		}
		createdUser, err := s.userRepo.CreateUser(newUser)
		if err != nil {
			return nil, "", err
		}
		user = createdUser
	} else if err != nil {
		return nil, "", err
	}
	token, err := s.generateTokenJWT(user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s *service) Logout() (string, error) {
	expiredToken := ""
	return expiredToken, nil
}

func (s *service) RequestPasswordReset(email string) error {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	token, err := generateSecureToken(32)
	if err != nil {
		return errors.New("gagal membuat token reset")
	}

	expiresAt := time.Now().Add(15 * time.Minute)
	user.ResetPasswordToken = token
	user.ResetPasswordTokenExpiresAt = expiresAt

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	go s.emailSvc.SendPasswordResetEmail(user.Email, token)

	return nil
}

func (s *service) VerifyAndResetPassword(input validations.ResetPasswordInput) error {
	if input.Password != input.PasswordConfirmation {
		return errors.New("password dan konfirmasi password tidak cocok")
	}

	user, err := s.userRepo.FindByResetToken(input.Token)
	if err != nil {
		return errors.New("token reset tidak valid atau sudah kedaluwarsa")
	}

	if time.Now().After(user.ResetPasswordTokenExpiresAt) {
		return errors.New("token reset tidak valid atau sudah kedaluwarsa")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	user.ResetPasswordToken = ""

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}