package initialize

import (
	"access-log-app/pkg/model"
	"access-log-app/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/xgtcode/log-demo"
	"net/http"
	"strconv"
)

func InitRouter() *gin.Engine{
	router := gin.Default()
	router.POST("/hello", helloFunc)
	router.POST("/count", CountFunc)
	return router
}

func helloFunc(c *gin.Context){
	log.Info("收到hello请求")
	var req = &model.Req{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg" : err.Error(),
		})
		return
	}
	err = service.Visit(req.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg" : err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

func CountFunc(c *gin.Context){
	log.Info("收到count请求")
	var req = &model.Req{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg" : err.Error(),
		})
		return
	}
	count, err := service.CountVisitLog(req.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg" : err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
			"count": strconv.Itoa(count),
		})
	}
}
