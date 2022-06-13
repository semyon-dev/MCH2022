package nko

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"net/http"
	"strconv"
)

func FindNkos(c *gin.Context) {

	ratingStart, _ := strconv.ParseInt(c.Query("ratingStart"), 10, 64)
	ratingEnd, _ := strconv.ParseInt(c.Query("ratingEnd"), 10, 64)
	tags := c.QueryArray("tags")
	searchQuery := c.Query("search_query")
	nkos := db.GetNKOsByFilters(searchQuery, tags)
	if nkos == nil {
		nkos = []model.NKO{}
	}

	var result = []model.NKO{}

	for _, nko := range nkos {
		nko.Rating = db.CountRating(nko.Rates)
		if nko.Rating >= float64(ratingStart) && nko.Rating <= float64(ratingEnd) {
			result = append(result, nko)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  result,
	})
}

func GetNKObyID(c *gin.Context) {
	nkoID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		cErrors.ReplyOnlyErr(cErrors.NKONotFound)
		return
	}
	nko, isExist := db.GetNKO(nkoID)
	if !isExist {
		cErrors.ReplyOnlyErr(cErrors.NKONotFound)
		return
	}

	nko.Rating = db.CountRating(nko.Rates)

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  nko,
	})
}
