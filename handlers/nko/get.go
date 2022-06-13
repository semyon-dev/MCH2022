package nko

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"net/http"
)

func FindNkos(c *gin.Context) {
	nkos := db.GetNKOs()
	if nkos == nil {
		nkos = []model.NKO{}
	}

	for _, nko := range nkos {
		nko.Rating = db.CountRating(nko.Rates)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  nkos,
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
