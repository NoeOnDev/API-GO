package controllers

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/helpers"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func AddPet(context *gin.Context) {
	var input models.Pet
	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.SavePet()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllPets(context *gin.Context) {
    user, err := helpers.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    pets, err := models.GetAllPetsByUserID(user.ID)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": pets})
}