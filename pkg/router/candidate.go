package router

import (
	"VoteMe_BE_TX/pkg/consts"
	"VoteMe_BE_TX/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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
	name := c.Query("name")
	candidate, err := services.GetOneCandidateDB(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败",
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

func Vote(c *gin.Context) {
	names := c.PostForm("names")
	ticket := c.PostForm("ticket")
	ticketDB := services.GetCurrentTicketDB()
	if time.Since(ticketDB.Time) > consts.VALID_TIME {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ticket过期",
		})
		return
	}
	if ticketDB.TicketString != ticket {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ticket错误",
		})
		return
	}
	nameList := strings.Fields(names)
	toVoteNum := len(nameList)
	if toVoteNum > ticketDB.Limit {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "超出投票数量",
		})
		return
	} else {
		//更新ticket
		services.UpdateTicket(ticketDB.TicketString, ticketDB.Limit-toVoteNum)
		for _, name := range nameList {
			_, err := services.GetOneCandidateDB(name)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "存在错误的候选人姓名",
				})
				return
			}
			err = services.AddScoreToCandidate(name)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"msg": err,
				})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "投票完毕",
	})
	return
}
