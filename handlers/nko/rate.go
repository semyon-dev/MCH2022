package nko

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"net/http"
)

func RateUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		cErrors.ReplyOnlyErr(cErrors.InvalidProjectID)
		return
	}
	var rate model.Rate
	rate.Id = primitive.NewObjectID()
	rate.AuthorId, _ = primitive.ObjectIDFromHex(c.GetString("userID"))
	err = c.ShouldBindJSON(&rate)
	if err != nil {
		cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput)
		return
	}
	err = db.AddRateToUser(userID, rate)
	if err != nil {
		cErrors.ReplyOnlyErr(cErrors.InternalServerError)
		return
	}
	c.JSON(http.StatusOK, cErrors.ReplyFullOK)
}
