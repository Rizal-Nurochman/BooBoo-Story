package usersservice

import (
	"errors"
	"os"
	"time"

	"github.com/BooBooStory/config/models"
	"github.com/BooBooStory/v1/users/usersrepository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/oauth2/v2"
	"gorm.io/gorm"
)

type UserService interface {
	DaftarAkun(user models.UserRegister) (*models.User, error)
	LoginAkun(user models.UserLogin) (*models.User, string, error)
	LoginOrRegisterWithGoogle(googleUserInfo *oauth2.Userinfo) (*models.User, string, error)
}

type userService struct {
	repository usersrepository.UserRepository
}

func NewUserService(repository usersrepository.UserRepository) *userService {
	return &userService{repository}
}

//Daftar Akun
func (s *userService) DaftarAkun(user models.UserRegister) (*models.User, error) {
	inputUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		// Password harus di-hash sebelum disimpan
	}

	userDitemukan, err := s.repository.FindByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
    // Handle error database yang tidak terduga (bukan user tidak ditemukan)
    return nil, err
	}
	if userDitemukan != nil {
			// Jika user ditemukan tanpa error
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

func (s *userService) LoginOrRegisterWithGoogle(googleUserInfo *oauth2.Userinfo) (*models.User, string, error) {
	// 1. Cek apakah user dengan email ini sudah ada di database
	user, err := s.repository.FindByEmail(googleUserInfo.Email)

	// Jika tidak ditemukan, buat user baru (Registrasi Otomatis)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := &models.User{
			Name:  googleUserInfo.Name,
			Email: googleUserInfo.Email,
			// Password bisa dikosongkan atau diisi random string,
			// karena loginnya tidak akan pakai password.
			Password: "",
		}

		createdUser, err := s.repository.CreateUser(newUser)
		if err != nil {
			return nil, "", err
		}
		user = createdUser
	} else if err != nil {
		// Handle error database lain
		return nil, "", err
	}

	// 2. Jika user sudah ada atau berhasil dibuat, generate JWT token internal kita
	token, err := s.GenerateTokenJWT(uint(user.ID)) // Perhatikan konversi tipe ID
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}