package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"deepthinking.do/safetygo/api/auth"
	"deepthinking.do/safetygo/api/models"
	"deepthinking.do/safetygo/api/responses"
	"deepthinking.do/safetygo/api/utils/formaterror"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {
	user := models.User{}
	if err := server.DB.Debug().Raw("select username, password from users where username = ? and password = unhex(md5(?));", username, password).Scan(&user).Error; err != nil {
		return "", errors.New("username or password incorrect")
	}
	// err = server.DB.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
	// if err != nil {
	// 	return "", err
	// }
	// err = deepthinking.VerifyPassword(user.Password, password)
	// if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
	// 	return "", err
	// }
	return auth.CreateToken(uint32(user.Id))
}
