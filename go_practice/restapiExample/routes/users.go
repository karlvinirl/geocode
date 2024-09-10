package routes

import (
	"fmt"
	"net/http"

	"example.com/go_practice/restapi/models"
	"example.com/go_practice/restapi/utils"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data"})
		return
	}

	err = user.ValidateCredentials()
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User is valid", "token": token})

}

func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User Saved Successfully"})

}
