package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mch2022/cErrors"
	"mch2022/session"
	"net/http"
	"strings"
)

func Middleware(c *gin.Context) {

	var notAuth = []string{"/auth", "/register"}

	for _, s := range notAuth {
		if strings.Contains(c.Request.URL.Path, s) {
			c.Next()
			return
		}
	}

	ID, userID, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, cErrors.ReplyOnlyErr(cErrors.InvalidToken))
		return
	}
	sessionID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println("error in middleware")
		c.AbortWithStatusJSON(http.StatusUnauthorized, cErrors.ReplyOnlyErr(cErrors.InvalidToken))
		return
	}
	c.Set("sessionID", sessionID.Hex())
	c.Set("userID", userID)
	c.Next()
}
