 ```go
 package main

 import (
     "github.com/gin-gonic/gin"
     "gorm.io/driver/sqlite"
     "gorm.io/gorm"
 )

 type Recipe struct {
     ID          uint   `json:"id" gorm:"primaryKey"`
     Title       string `json:"title"`
     Description string `json:"description"`
 }

 var db *gorm.DB
 var err error

 func main() {
     db, err = gorm.Open(sqlite.Open("recipes.db"), &gorm.Config{})
     if err != nil {
         panic("failed to connect database")
     }

     db.AutoMigrate(&Recipe{})

     r := gin.Default()

     r.GET("/recipes", GetRecipes)
     r.POST("/recipes", CreateRecipe)
     r.PUT("/recipes/:id", UpdateRecipe)
     r.DELETE("/recipes/:id", DeleteRecipe)

     r.Run()
 }

 func GetRecipes(c *gin.Context) {
     var recipes []Recipe
     db.Find(&recipes)
     c.JSON(200, recipes)
 }

 func CreateRecipe(c *gin.Context) {
     var recipe Recipe
     if err := c.ShouldBindJSON(&recipe); err == nil {
         db.Create(&recipe)
         c.JSON(200, recipe)
     } else {
         c.JSON(400, err.Error())
     }
 }

 func UpdateRecipe(c *gin.Context) {
     var recipe Recipe
     if err := db.Where("id = ?", c.Param("id")).First(&recipe).Error; err == nil {
         if err := c.ShouldBindJSON(&recipe); err == nil {
             db.Save(&recipe)
             c.JSON(200, recipe)
         } else {
             c.JSON(400, err.Error())
         }
     } else {
         c.JSON(404, err.Error())
     }
 }

 func DeleteRecipe(c *gin.Context) {
     var recipe Recipe
     if err := db.Where("id = ?", c.Param("id")).First(&recipe).Error; err == nil {
         db.Delete(&recipe)
         c.JSON(200, gin.H{"message": "Recipe deleted"})
     } else {
         c.JSON(404, err.Error())
     }
 }
 ```
