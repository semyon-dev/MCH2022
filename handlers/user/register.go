package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"mch2022/service"
	"mch2022/session"
	"net/http"
)

func Register(c *gin.Context) {

	var inputUser model.UserWithCredentials

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	if inputUser.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	newUser, err := service.CreateUser(inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": cErrors.ReplyErr(err),
		})
		return
	}

	newSession, err := session.Create(newUser.Id.Hex(), newUser.Id.Hex(), c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	err = db.Insert(db.SessionsCollection, newSession)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": cErrors.ReplyOK(),
		"data": gin.H{
			"token":        newSession.Token,
			"refreshToken": newSession.RefreshToken,
		},
	})
}
