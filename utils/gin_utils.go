package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetHtml(router gin.IRouter, relativePath string, content string) {
	rs := []byte(content)
	router.GET(relativePath, func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.Header("content-type", "text/html;charset=UTF-8")
		ctx.Header("content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})

}

func GetJson(router gin.IRouter, relativePath string, json string) {
	router.GET(relativePath, func(ctx *gin.Context) {
		rs := []byte(json)

		ctx.Status(http.StatusOK)
		ctx.Header("content-type", "application/json;charset=UTF-8")
		ctx.Header("content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})
}

func GetJs(router gin.IRouter, relativePath string, hexContent string) {
	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}

	router.GET(relativePath, func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.Header("content-type", "application/javascript;charset=UTF-8")
		ctx.Header("content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})
}

func GetCss(router gin.IRouter, relativePath string, content string) {
	router.GET(relativePath, func(ctx *gin.Context) {
		rs := []byte(content)

		ctx.Status(http.StatusOK)
		ctx.Header("content-type", "text/css;charset=UTF-8")
		ctx.Header("content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})
}

func GetOther(router gin.IRouter, relativePath string, hexContent string) {

	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	router.GET(relativePath, func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.Header("Content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})

}
