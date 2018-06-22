package pkg

import (
	"github.com/gin-gonic/gin"
	"time"
	"crypto/md5"
	"encoding/hex"
	"io"
	"encoding/base64"
	"crypto/rand"
)

var UserMap map[string]string

func init() {
	UserMap = make(map[string]string)
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(token) > 7 {
		return token[7:]
	}
	return ""
}

func GetTimeNow() string {
	//格式化必须是这个时间点，Go语言诞生时间？
	return time.Now().Format("2006-01-02 15:04:05.00")
}


func GenGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getmd5string(base64.URLEncoding.EncodeToString(b))
}

func getmd5string(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
