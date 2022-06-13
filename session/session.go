package session

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"mch2022/config"
	"mch2022/model"
	"net/http"
	"strings"
	"time"
)

var TimeExpiredError = errors.New("TimeExpiredError")
var InvalidTokenError = errors.New("InvalidTokenError")

func Create(id, userID, userAgent, IP string) (session model.Session, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["userID"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AccessTokenLifetime)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	session.Token, err = at.SignedString([]byte(config.AccessSecret))
	if err != nil {
		return session, err
	}
	session.RefreshToken = NewRefreshToken()
	session.UserAgent = userAgent
	session.IP = IP
	session.UserID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return session, err
	}
	session.ExpiresIn = primitive.NewDateTimeFromTime(time.Now().Add(time.Minute * time.Duration(config.RefreshTokenLifetime)))
	return session, err
}

type MyCustomClaims struct {
	Id     string `json:"id"`
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (id, userID string, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessSecret), err
	})

	if err != nil {
		if !strings.Contains(err.Error(), "token is expired") {
			log.Println(err)
		}
	}

	if token == nil {
		return "", "", InvalidTokenError
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		id = claims.Id
		userID = claims.UserID
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// Token is either expired or not active yet
			if ve.Errors == 16 {
				// Token is either expired
				return id, userID, TimeExpiredError
			} else {
				fmt.Println("Couldn't handle this token:", err)
				return "", "", InvalidTokenError
			}
		}
	}

	return id, userID, nil
}

func NewRefreshToken() string {
	b := make([]byte, 32)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	if _, err := r.Read(b); err != nil {
		return ""
	}

	return fmt.Sprintf("%x", b)
}

// just parsing without checking validation
func parseBearer(c *gin.Context) (token string, isValid bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	return headerParts[1], true
}

// ParseBearer common checking
func ParseBearer(c *gin.Context) (id, userID string, isValid bool) {
	token, isvalid := parseBearer(c)
	if !isvalid {
		return
	}
	id, userID, err := ParseToken(token)
	if err != nil {
		return
	}
	return id, userID, true
}
