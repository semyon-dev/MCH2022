package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/session"
	"net/http"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil || jsonInput.Password == "" {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	user, exist := db.GetUserWithCredentials("email", jsonInput.Email)
	if !exist {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.UserNotFoundByEmail))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonInput.Password)) != nil {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidPassword))
		return
	}

	newSession, err := session.Create(user.Id.Hex(), user.Id.Hex(), c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	err = db.Insert(db.SessionsCollection, newSession)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data": gin.H{
			"token":        newSession.Token,
			"refreshToken": newSession.RefreshToken,
		},
	})
}

func AuthNKO(c *gin.Context) {

	jsonInput := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil || jsonInput.Password == "" {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidBodyInput))
		return
	}

	user, exist := db.GetNKOByField("email", jsonInput.Email)
	if !exist {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.UserNotFoundByEmail))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonInput.Password)) != nil {
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InvalidPassword))
		return
	}

	newSession, err := session.Create(user.Id.Hex(), user.Id.Hex(), c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	err = db.Insert(db.SessionsCollection, newSession)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, cErrors.ReplyOnlyErr(cErrors.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data": gin.H{
			"token":        newSession.Token,
			"refreshToken": newSession.RefreshToken,
		},
	})
}
