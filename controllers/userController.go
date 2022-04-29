package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"final-project/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	user := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	user.Password = helpers.HassPass(user.Password)
	err := db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	userIdParam, _ := strconv.Atoi(c.Param("userId"))

	if contentType == appJson {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	user.ID = uint(userIdParam)

	err := db.Model(&user).Where("id = ?", userIdParam).Updates(models.User{
		Username: user.Username,
		Email:    user.Email,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Debug().Where("id = ?", userIdParam).Take(&user)

	UserUpdate := views.UserUpdate{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, UserUpdate)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	userIdParam, _ := strconv.Atoi(c.Param("userId"))

	err := db.Debug().Where("id = ?", userIdParam).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Uer not found",
		})
		return
	}

	db.Debug().Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email",
		})
		return
	}

	compare := helpers.ComparePass([]byte([]byte(User.Password)), []byte(password))
	if !compare {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid password",
		})
		return
	}

	token := helpers.GenerateJwtToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
