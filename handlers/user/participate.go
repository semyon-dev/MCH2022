package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"net/http"
)

func ProjectParticipate(c *gin.Context) {

	var p string
	if p = c.Param("id"); p == "" {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	projectID, _ := primitive.ObjectIDFromHex(p)
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))
	db.AddUserToProject(projectID, userID)

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
	})
}

func DeleteParticipant(c *gin.Context) {

	var p string
	if p = c.Param("id"); p == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": cErrors.InvalidBodyInput,
		})
		return
	}

	var u string
	if u = c.Param("userID"); u == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": cErrors.InvalidBodyInput,
		})
		return
	}

	projectID, _ := primitive.ObjectIDFromHex(p)

	userID, _ := primitive.ObjectIDFromHex(u)

	db.DeleteUserFromProject(projectID, userID)

	c.JSON(http.StatusNoContent, gin.H{
		"error": cErrors.ReplyOK(),
	},
	)
}
