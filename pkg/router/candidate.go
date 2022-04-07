package router

import (
	"VoteMe_BE_TX/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCandidates(c *gin.Context) {
	candidates := services.GetCandidatesDB()
	var results []map[string]interface{}
	for _, temp := range candidates {
		results = append(results, map[string]interface{}{
			"name":  temp.Name,
			"score": temp.Score,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": results,
	})
	return
}
func GetOneCandidate(c *gin.Context) {
	name := c.Param("name")
	candidate, err := services.GetOneCandidateDB(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "查询失败",
			"msg2": err,
		})
		return
	}
	//var results []map[string]interface{}
	//for _, temp := range candidates {
	//	results = append(results, map[string]interface{}{
	//		"name":  temp.Name,
	//		"score": temp.Score,
	//	})
	//}
	c.JSON(http.StatusOK, gin.H{
		"msg": candidate,
	})
	return
}
