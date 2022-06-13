package nko

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mch2022/cErrors"
	"mch2022/db"
	"mch2022/model"
	"net/http"
	"strconv"
	"time"
)

func FindProjects(c *gin.Context) {
	ageStart, _ := strconv.ParseInt(c.Query("ageStart"), 10, 64)
	ageEnd, _ := strconv.ParseInt(c.Query("ageEnd"), 10, 64)
	participation := c.Query("participation")
	direction := c.Query("direction")
	tags := c.QueryArray("tags")
	roles := c.QueryArray("roles")
	skills := c.QueryArray("skills")
	name := c.Query("name")
	location := c.Query("location")
	timeParam := c.Query("time")
	searchQuery := c.Query("search_query")
	var authorID primitive.ObjectID
	if c.Query("authorID") != "" {
		authorID, _ = primitive.ObjectIDFromHex(c.Query("authorID"))
	}
	projects := db.GetProjectsByFilters(int(ageStart), int(ageEnd), participation, location, direction, name, searchQuery, tags, roles, skills, authorID)
	if projects == nil {
		projects = []model.Project{}
	}

	var itogProjects = []model.Project{}
	if timeParam != "" {
		parse, err := time.Parse(time.RFC3339, timeParam)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": cErrors.ReplyErr(cErrors.InvalidTime),
			})
			return
		}
		for _, v := range projects {
			if v.TimeStart.Time().Year() == parse.Year() && v.TimeStart.Time().Day() == parse.Day() && v.TimeStart.Time().Month() == parse.Month() {
				itogProjects = append(itogProjects, v)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  projects,
	})
}

func GetProjectByID(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("projectID"))
	if err != nil {
		return
	}
	pr, _ := db.GetProjectByID(projectID)

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  pr,
	})
}

func GetFiltersRefs(c *gin.Context) {

	pr := db.GetProjects()

	m := map[string]map[string]bool{
		"roles":  {},
		"skills": {},
		"tags":   {},
	}

	for _, project := range pr {

		for _, role := range project.Roles {
			if _, ok := m["roles"][role]; !ok {
				m["roles"][role] = true
			}
		}

		for _, skill := range project.Skills {
			if _, ok := m["skills"][skill]; !ok {
				m["skills"][skill] = true
			}
		}

		for _, tag := range project.Tags {
			if _, ok := m["tags"][tag]; !ok {
				m["tags"][tag] = true
			}
		}

	}

	result := make(map[string][]string, 0)

	for s, m2 := range m {
		result[s] = make([]string, 0)
		for s2, _ := range m2 {
			result[s] = append(result[s], s2)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  result,
	})

}

func SearchProjects(c *gin.Context) {
	name := c.Query("search_query")
	pr := db.SearchProjects(name)
	if pr == nil {
		pr = []model.Project{}
	}
	c.JSON(http.StatusOK, gin.H{
		"error": cErrors.ReplyOK(),
		"data":  pr,
	})
}
