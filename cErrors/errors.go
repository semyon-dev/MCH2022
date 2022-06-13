// Package cErrors for custom app errors
package cErrors

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ProjectNotFound = errors.New("project not found")
var AlreadyRegistered = errors.New("user has already registered for this project")
var UserIsBanned = errors.New("user is banned on this project")
var InternalServerError = errors.New("internal server error")
var NotAdminOfProject = errors.New("you are not admin of project")
var EmailAlreadyUsed = errors.New("email already used")
var UserNotFoundByEmail = errors.New("user not found by email in database")
var UserNotFound = errors.New("user not found")
var NKONotFound = errors.New("nko not found")
var InvalidPassword = errors.New("invalid password")
var InvalidProjectID = errors.New("invalid ProjectID")
var InvalidUserID = errors.New("invalid userID")
var InvalidBodyInput = errors.New("invalid body input")
var NotRegisteredToProject = errors.New("you are not registered to this project")
var InvalidToken = errors.New("invalid token")
var TokenExpired = errors.New("token is expired")
var SessionExpired = errors.New("session expired")
var InvalidAuthorID = errors.New("invalid or empty authorID")
var InvalidTime = errors.New("can't parse time")
var InvalidNKOID = errors.New("InvalidNKOID")

var ErrToCode = map[error]int{

	nil: 0,

	// registration
	EmailAlreadyUsed: 2,

	// authorization (email & pass)
	UserNotFoundByEmail: 10,
	InvalidPassword:     13,
	InvalidToken:        14,
	TokenExpired:        15,
	SessionExpired:      16,

	// registering to project (or other)
	ProjectNotFound:        40,
	AlreadyRegistered:      41,
	UserIsBanned:           42,
	InvalidProjectID:       43,
	NotRegisteredToProject: 44,

	// administrating of project
	NotAdminOfProject: 70,

	// other,
	InvalidBodyInput: 80,
	InvalidUserID:    81,
	UserNotFound:     82,
	InvalidTime:      83,
	NKONotFound:      84,
	InvalidNKOID:     84,
}

func ReplyOnlyErr(err error) gin.H {
	return gin.H{"error": ReplyErr(err)}
}

func ReplyErr(err error) gin.H {
	code, ok := ErrToCode[err]
	if !ok {
		code = 100
	}
	return gin.H{"message": err.Error(), "code": code}
}

func ReplyOK() gin.H {
	return gin.H{"message": "ok", "code": 0}
}

var ReplyFullOK = gin.H{"error": ReplyOK()}
