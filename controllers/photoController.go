package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"final-project/views"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJsonPhoto = "application/json"
)

func NewPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	if contentType == appJsonPhoto {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	userID := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))
	photo.UserID = userID
	err := db.Debug().Create(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	photoCreate := views.PhotoCreateView{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	c.JSON(http.StatusCreated, photoCreate)
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	photos := []models.Photo{}
	userID := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))
	photos_views := []views.PhotoView{}

	db.Find(&photos)
	for i, _ := range photos {
		db.Model(&photos[i]).Association("User").Find(&photos[i].User)
		if photos[i].UserID == userID {
			photos_views = append(photos_views, views.PhotoView{
				ID:        photos[i].ID,
				Title:     photos[i].Title,
				Caption:   photos[i].Caption,
				PhotoUrl:  photos[i].PhotoUrl,
				UserID:    photos[i].UserID,
				CreatedAt: photos[i].CreatedAt,
				UpdatedAt: photos[i].UpdatedAt,
				User: views.UserView{
					Email:    photos[i].User.Email,
					Username: photos[i].User.Username,
				},
			})
		}
	}

	if len(photos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not found",
			"message": "No photos found",
		})
		return
	}

	c.JSON(http.StatusOK, photos_views)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJsonPhoto {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.ID = uint(photoID)

	err := db.Debug().Model(&photo).Updates(models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	db.Where("id = ?", photoID).Find(&photo)

	photoCreate := views.PhotoUpdateView{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	c.JSON(http.StatusCreated, photoCreate)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	photo := models.Photo{}

	db.Where("id = ?", photoID).Find(&photo)

	err := db.Debug().Model(&photo).Delete(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
