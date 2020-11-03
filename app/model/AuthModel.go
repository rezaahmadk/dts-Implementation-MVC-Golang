package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rezaahmadk/dts-Implementation-MVC-Golang/app/utils"
	"gorm.io/gorm"
)

type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthModel struct {
	DB *gorm.DB
}

func (model AuthModel) Login(auth Auth) (bool, error, string) {
	var account Account
	result := model.DB.Where(&Account{Name: auth.Name}).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account Not Found"), ""
		}

		return false, result.Error, ""
	}

	err := utils.HashComparator([]byte(account.Password), []byte(auth.Password))
	if err != nil {
		return false, errors.Errorf("Incorrect Password"), ""
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":            auth.Name,
		"account_account": account.AccountNumber,
	})

	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return false, err, ""
	}

	return true, nil, token
}
