package info

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Parm []string
var Secret string

func GetAllInfo(c *gin.Context) {
	value,err := allinfo()
	if err != nil{
		c.JSON(http.StatusAccepted, gin.H{"status": -1, "metrics": err})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": 0, "metrics": value})
	return
}

func GetAppointInfo(c *gin.Context){
	parm := c.Param("info")
	value,err := appointinfo(parm)
	if err != nil{
		c.JSON(http.StatusAccepted, gin.H{"status": -1, "metrics": err})
		return
	}
	c.JSON(http.StatusAccepted,gin.H{"status": 0 ,"metrics":value})
	return
}


