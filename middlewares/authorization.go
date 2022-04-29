package middlewares

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJsonAuthorization = "application/json"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userId := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

		product := models.Photo{}

		err = db.Where("id = ?", productID).First(&product).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}

		if product.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this product",
			})
			return
		}

		c.Next()
	}
}

func CommentCreateAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		contentType := helpers.GetContentType(c)
		comment := models.Comment{}
		userID := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

		if contentType == appJsonAuthorization {
			c.ShouldBindJSON(&comment)
		} else {
			c.ShouldBind(&comment)
		}

		photoId := comment.PhotoID
		photo := models.Photo{}

		err := db.Debug().Where("id = ?", photoId).First(&photo).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this user",
			})
			return
		}

		c.Set("message", comment.Message)
		c.Set("photoId", photoId)
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userID := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		comment := models.Comment{}

		err = db.Debug().Where("id = ?", commentId).First(&comment).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Comment not found",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this comment",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userID := uint(c.MustGet("userToken").(jwt.MapClaims)["userId"].(float64))

		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		socialMedia := models.SocialMedia{}

		err = db.Debug().Where("id = ?", socialMediaId).First(&socialMedia).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Social Media not found",
			})
			return
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this social media",
			})
			return
		}

		c.Next()
	}
}
