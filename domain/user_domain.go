package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}


type CustomerInfo struct {
	Name string
	Email string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

type User struct {
	Base     `valid:"required"`
	Username string `json:"username" gorm:"type:varchar(255);not null;default:null" valid:"notnull"`
	Name     string `json:"name" gorm:"type:varchar(255);not null;default:null" valid:"notnull"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index;not null;default:null" valid:"notnull,email"`
	Password string `json:"-" gorm:"type:varchar(255); not null;default:null" valid:"notnull"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index" valid:"notnull"`
}

func NewUser(user *User) (*User, error) {
	
	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Prepare() error {


	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)
	token, err := CreateToken(user)

	if err != nil {
		return err
	}

	user.Token = token

	return nil

}

func CreateToken(paylod *User) (string, error) {
	claims := CustomClaimsExample{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"",
		CustomerInfo{Name: paylod.Name, Email: paylod.Email},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte("secureSecretText"))

	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyPassword(hash, password string) bool {

	err :=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}