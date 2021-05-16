package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ValidCORSDoman(c *gin.Context, allowDomains ...string) string {
	origin := c.Request.Header.Get("Origin")
	isdebug := c.GetBool("isDebug")
	if isdebug {
		return origin
	} else {
		return ValidCORSDomanOrigin(origin, allowDomains...)
	}
}

func ValidCORSDomanOrigin(origin string, allowDomains ...string) string {
	if origin == "null" {
		return origin
	}
	allowDomains = append(allowDomains, "smm.cn", "smmadmin.cn", "metal.com", "mmiprices.com", "anhuida.com", "anpiaoda.com", "drageasy.com")
	for _, domain := range allowDomains {
		pattern := fmt.Sprintf(`(?P<host>.*?.%s)(:\d+)?\/?`, domain)
		r := regexp.MustCompile(pattern)
		result := r.FindStringSubmatch(origin)
		if len(result) > 0 {
			return origin
		}
	}

	return "https://www.smm.cn"
}

func DoResponse(code int, msg string, res interface{}, c *gin.Context) *string {
	c.Header("Access-Control-Allow-Origin", ValidCORSDoman(c))
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, SMM-TOKEN, X-Auth-Token, SMM-ADMIN-TOKEN, SMM-SOURCE, SMM-DEVICE, x-requested-with")

	jQuery := c.Query("callback")
	response := CommonResponse{
		code, msg, res}

	s, err := json.Marshal(response)
	if err != nil {
		log.Println("----error----Marshal----", err)
	}
	r := string(s)

	var resStr string
	if jQuery != "" {
		resStr = jQuery + "(" + string(r) + ")"
	} else {
		resStr = string(r)

		c.Header("Content-Type", "application/json; charset=UTF-8")
	}

	c.String(200, resStr)
	return &r
}
