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
	appJsonSocialMedia = "application/json"
)

func NewSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userId := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

	socialMedia := models.SocialMedia{}

	if contentType == appJsonSocialMedia {
		c.BindJSON(&socialMedia)
	} else {
		c.Bind(&socialMedia)
	}

	socialMedia.UserID = userId

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	socialmedias := views.SocialMediaCreate{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, socialmedias)
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

	socialmedias := []models.SocialMedia{}
	socialmediaviews := []views.SocialMedias{}

	db.Find(&socialmedias, "user_id = ?", userId)
	for i, _ := range socialmedias {
		db.Model(&socialmedias[i]).Association("User").Find(&socialmedias[i].User)
		socialmediaviews = append(socialmediaviews, views.SocialMedias{
			ID:             socialmedias[i].ID,
			Name:           socialmedias[i].Name,
			SocialMediaUrl: socialmedias[i].SocialMediaUrl,
			UserID:         socialmedias[i].UserID,
			CreatedAt:      socialmedias[i].CreatedAt,
			UpdatedAt:      socialmedias[i].UpdatedAt,
			User: views.SocialMediaUser{
				ID:       socialmedias[i].User.ID,
				Username: socialmedias[i].User.Username,
			},
		})
	}

	if len(socialmediaviews) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Social Media Yet",
			"message": "Your social media is not available",
		})
	}
	c.JSON(http.StatusOK, socialmediaviews)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	socialmediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialmedia := models.SocialMedia{}

	if contentType == appJsonSocialMedia {
		c.BindJSON(&socialmedia)
	} else {
		c.Bind(&socialmedia)
	}

	socialmedia.ID = uint(socialmediaId)

	err := db.Debug().Model(&socialmedia).Updates(models.SocialMedia{
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request yas",
			"message": socialmedia,
		})
		return
	}

	socialmedias := views.SocialMediaUpdate{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
		UserID:         socialmedia.UserID,
		UpdatedAt:      socialmedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, socialmedias)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialmediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	socialmedia := models.SocialMedia{}
	socialmedia.ID = uint(socialmediaId)

	err := db.Debug().Model(&socialmedia).Delete(&socialmedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
