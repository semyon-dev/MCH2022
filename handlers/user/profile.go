package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/service"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	obj, _ := primitive.ObjectIDFromHex(c.GetString("userID"))
	user, ok := db.GetUserById(obj)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, cErrors.ReplyOnlyErr(cErrors.UserNotFound))
		return
	}

	user.Rating = db.CountRating(user.Rates)

	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"error": cErrors.ReplyOK(),
	})
}

func GetNKO(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.GetString("userID"))
	nko, ok := db.GetNKO(id)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, cErrors.ReplyOnlyErr(cErrors.UserNotFound))
		return
	}

	nko.Rating = db.CountRating(nko.Rates)
	c.JSON(http.StatusOK, gin.H{
		"data":  nko,
		"error": cErrors.ReplyOK(),
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": cErrors.ReplyErr(cErrors.InvalidUserID),
		})
		return
	}
	user, ok := db.GetUserById(objectId)
	if ok {

		user.Rating = db.CountRating(user.Rates)

		c.JSON(http.StatusOK, gin.H{
			"data":  user,
			"error": cErrors.ReplyOK(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": cErrors.ReplyErr(cErrors.UserNotFound),
	})
}

func GetNKOProjects(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))
	createdProjects, err := service.GetMyProjects(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": cErrors.ReplyErr(err),
			"data":  createdProjects,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  createdProjects,
	})
}

func ProjectParticipants(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))
	_, ok := db.GetUserById(userID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": cErrors.ReplyErr(cErrors.InvalidToken),
		})
		return
	}
	projectID, err := primitive.ObjectIDFromHex(c.Param("projectID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": cErrors.ReplyErr(cErrors.InvalidProjectID),
		})
		return
	}
	project, isExist := db.GetProjectByID(projectID)
	if !isExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": cErrors.ReplyErr(cErrors.ProjectNotFound),
		})
		return
	}
	var limitInt, pageInt int
	limit, isExist := c.GetQuery("limit")
	limitInt, err = strconv.Atoi(limit)
	if err != nil || !isExist {
		limitInt = 20
	}
	page, isExist := c.GetQuery("page")
	pageInt, err = strconv.Atoi(page)
	if err != nil || !isExist {
		pageInt = 0
	}
	users := db.GetUsersByIdsWithPagination(project.Participants, int64(limitInt), int64(limitInt*pageInt))

	for _, user := range users {
		user.Rating = db.CountRating(user.Rates)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  users,
	})
}
