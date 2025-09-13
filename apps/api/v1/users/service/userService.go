package service

import (
	"errors"
	"os"
	"time"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/v1/users/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	DaftarAkun(user models.UserRegister) (*models.User, error)
	LoginAkun(user models.UserLogin) (*models.User, string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

//Daftar Akun
func (s *userService) DaftarAkun(user models.UserRegister) (*models.User, error) {
	inputUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		// Password harus di-hash sebelum disimpan
	}

	_, err := s.repository.FindByEmail(user.Email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return inputUser, err
	}

	inputUser.Password = string(passwordHash)
	newUser, err := s.repository.CreateUser(inputUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) GenerateTokenJWT(IDUser uint) (string, error)  {
	//Payload claims adalah data yang ingin kita simpan di dalam token
	claims:=jwt.MapClaims{} //Membuat wadah kosong untuk menyimpan informasi token
	claims["user_id"]=IDUser //Menyimpan ID user ke dalam token
	claims["exp"] = time.Now().Add(time.Hour*24).Unix() //Token berlaku selama 24 jam

	//Membuat token dengan metode HS256 dan claims di atas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	//Menandatangani token dengan secret key dari .env dan mengubahnya menjadi string
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *userService) LoginAkun(user models.UserLogin) (*models.User, string, error) {
	email := user.Email
	userDitemukan, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	password := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(userDitemukan.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	//Kalo passwordnya cocok, generate token JWT
	token, err := s.GenerateTokenJWT(uint(userDitemukan.ID))
	if err != nil {
		return nil, "", err
	}

	return userDitemukan, token, nil
}