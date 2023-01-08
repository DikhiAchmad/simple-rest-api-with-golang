package routes

import (
	"simple-crud-golang/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })
    r.GET("/user", controllers.FindUser)
    r.POST("/user", controllers.CreateUser)
    r.GET("/user/:id", controllers.FindUserById)
    r.PATCH("/user/:id", controllers.UpdateUser)
    r.DELETE("user/:id", controllers.DeleteUser)
    
    return r
}