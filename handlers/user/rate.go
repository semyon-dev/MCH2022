package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"net/http"
	"sort"
)

func RateNKO(c *gin.Context) {
	nkoID, err := primitive.ObjectIDFromHex(c.Param("nkoID"))
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
	err = db.AddRateToNKO(nkoID, rate)
	if err != nil {
		cErrors.ReplyOnlyErr(cErrors.InternalServerError)
		return
	}
	c.JSON(http.StatusOK, cErrors.ReplyFullOK)
}

func GetWholeRatingList(c *gin.Context) {

	users := db.GetUsers()

	usersRating := make([]model.UserRatingListItem, 0)
	for i, user := range users {
		users[i].Rating = db.CountRating(user.Rates)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Rating > users[j].Rating
	})

	for i, user := range users {
		var isClient bool

		userID := user.ID.Hex()
		clientID := c.GetString("userID")

		if userID == clientID {
			isClient = true
		}

		usersRating = append(usersRating, model.UserRatingListItem{
			Place:    i + 1,
			Name:     user.Name,
			UserID:   user.ID,
			Rating:   user.Rating,
			IsClient: isClient,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  usersRating,
	})

}
