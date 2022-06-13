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

	ratingStart, _ := strconv.ParseFloat(c.Query("ratingStart"), 10)
	ratingEnd, _ := strconv.ParseFloat(c.Query("ratingEnd"), 10)
	tags := c.QueryArray("tags")
	searchQuery := c.Query("search_query")
	nkos := db.GetNKOsByFilters(searchQuery, tags)
	if nkos == nil {
		nkos = []model.NKO{}
	}

	var result = []model.NKO{}

	for i, _ := range nkos {
		nkos[i].Rating = db.CountRating(nkos[i].Rates)
		if nkos[i].Rating >= ratingStart && nkos[i].Rating <= ratingEnd {
			result = append(result, nkos[i])
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
