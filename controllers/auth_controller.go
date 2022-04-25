package controllers

// TODO: Extract service from controller

import (
	"fmt"
	"net/http"
	"ufp/smartlion-api/domain"
	"ufp/smartlion-api/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	credentialsError  = fmt.Errorf("invalid credentials")
	unauthorizedError = fmt.Errorf("unauthorized access")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(ctx *gin.Context) {
	var credentials domain.User
	var user domain.User

	if err := ctx.BindJSON(&credentials); err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
		return
	}

	services.Database.Find(&user, "email = ?", credentials.Email)
	if user.Email == "" || !CheckPasswordHash(credentials.Password, user.Password) {
		closeWithError(ctx, http.StatusUnauthorized, credentialsError)
		return
	}

	token := services.GenerateTokenJWT(credentials)

	if token == "" {
		closeWithError(ctx, http.StatusUnauthorized, unauthorizedError)
		return
	}

	closeWithData(ctx, http.StatusOK, gin.H{"token": token})
}

func RegisterHandler(ctx *gin.Context) {
	var creds domain.User

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hash, _ := HashPassword(creds.Password)

	creds.Password = hash
	result := services.Database.Save(&creds)
	if result.RowsAffected != 0 {
		closeWithData(ctx, http.StatusOK, gin.H{"user_id": creds.ID})
		return
	}

	closeWithError(ctx, http.StatusNotAcceptable, credentialsError)
}

func RefreshHandler(ctx *gin.Context) {

	user := domain.User{
		Email: ctx.GetHeader("email"),
	}

	token := services.GenerateTokenJWT(user)

	if token == "" {
		closeWithError(ctx, http.StatusUnauthorized, unauthorizedError)
		return
	}

	closeWithData(ctx, http.StatusOK, gin.H{"token": token})
}
