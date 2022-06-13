package nko

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mch2022/cErrors"
	"mch2022/model"
	"mch2022/service"
	"net/http"
)

func CreateProject(c *gin.Context) {
	var pr model.Project
	var err error

	err = c.ShouldBindJSON(&pr)
	if err != nil {
	}

	pr.AuthorId, err = primitive.ObjectIDFromHex(c.GetString("userID"))
	if err != nil {
		log.Println(err)
		return
	}

	id, err := service.CreateProject(pr)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": cErrors.ReplyOK(),
		"data": gin.H{
			"id": id,
		},
	})
}
