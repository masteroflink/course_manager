package controllers

import (
	"encoding/json"
	"io"
	"log"
	"main/api/auth"
	"main/api/formaterror"
	"main/api/models"
	"main/api/responses"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) SignIn(email string, password string) (string, error) {
	user := models.User{}

	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return auth.CreateToken(user.ID)
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
