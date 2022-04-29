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
	appJsonComment = "application/json"
)

func NewComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userId := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

	comment := models.Comment{}

	if contentType == appJsonComment {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserID = userId
	comment.PhotoID = c.MustGet("photoId").(uint)
	comment.Message = c.MustGet("message").(string)

	err := db.Debug().Create(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	commentCreate := views.CommentCreate{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	c.JSON(http.StatusCreated, commentCreate)
}

func GetComments(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

	comment := []models.Comment{}
	commentGet := []views.GetComments{}
	user := models.User{}

	db.Where("id = ?", userId).First(&user)

	db.Find(&comment, "user_id = ?", userId)
	for i, _ := range comment {
		db.Model(&comment[i]).Association("Photo").Find(&comment[i].Photo)
		commentGet = append(commentGet, views.GetComments{
			ID:        comment[i].ID,
			Message:   comment[i].Message,
			PhotoID:   comment[i].PhotoID,
			UserID:    comment[i].UserID,
			UpdatedAt: comment[i].UpdatedAt,
			CreatedAt: comment[i].CreatedAt,
			User: views.UserComment{
				ID:       comment[i].UserID,
				Email:    user.Email,
				Username: user.Username,
			},
			Photo: views.PhotoComment{
				ID:       comment[i].Photo.ID,
				Title:    comment[i].Photo.Title,
				Caption:  comment[i].Photo.Caption,
				PhotoUrl: comment[i].Photo.PhotoUrl,
				UserID:   comment[i].Photo.UserID,
			},
		})
	}

	if len(commentGet) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Comment Yet",
			"message": "Your comment is not available",
		})
		return
	}

	c.JSON(http.StatusOK, commentGet)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	commentID, _ := strconv.Atoi(c.Param("commentId"))

	comment := models.Comment{}

	if contentType == appJsonComment {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.ID = uint(commentID)

	err := db.Debug().Model(&comment).Updates(models.Comment{Message: comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	db.Where("id = ?", comment.ID).First(&comment)
	photo := models.Photo{}
	db.Where("id = ?", comment.PhotoID).First(&photo)

	commentUpdate := views.CommentUpdate{
		ID:        comment.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	c.JSON(http.StatusCreated, commentUpdate)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	commentID, _ := strconv.Atoi(c.Param("commentId"))

	comment := models.Comment{}

	if contentType == appJsonComment {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.ID = uint(commentID)

	err := db.Debug().Model(&comment).Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
