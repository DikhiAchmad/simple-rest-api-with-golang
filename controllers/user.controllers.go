package controllers

import (
	"net/http"
	"simple-crud-golang/models"
	"simple-crud-golang/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
    Email	   string    `json:"email"`
    Name	   string    `json:"name"`
 	Password   string    `json:"password"`
}
type UpdateUserInput struct {
    Email	   string    `json:"email"`
    Name	   string    `json:"name"`
 	Password   string    `json:"password"`
}


// GET /user
// Get all user
func FindUser(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var users []models.User
    db.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}

// create User
func CreateUser(c *gin.Context) {
    // Validate input
    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	password, _ := utils.HashPassword(input.Password)

    // Create user
    user := models.User{Name: input.Name, Email: input.Email, Password: password}

    db := c.MustGet("db").(*gorm.DB)
    db.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// find user
func FindUserById(c *gin.Context) { 
    // Get model if exist
    var user models.User

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// update user
func UpdateUser(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }

    // Validate input
    var input UpdateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.User

    updatedInput.Name = input.Name
    updatedInput.Email = input.Email
    updatedInput.Password = input.Password

    db.Model(&user).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// delete
func DeleteUser(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }

    db.Delete(&user)

    c.JSON(http.StatusOK, gin.H{"data": true})
}