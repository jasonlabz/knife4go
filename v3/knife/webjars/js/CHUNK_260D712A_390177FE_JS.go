package js

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlabz/knife4go/constant"
	"github.com/jasonlabz/knife4go/utils"
)

const (
	CHUNK_260D712A_390177FE_JS_RELATIVE_PATH = constant.RootPath + "/webjars/js/chunk-260d712a.390177fe.js"
	// 文件内容的16进制表示
)

func AddRouterOfChunk260d712a390177feJs(router gin.IRouter) {

	utils.GetJs(router, CHUNK_260D712A_390177FE_JS_RELATIVE_PATH, CHUNK_260D712A_390177FE_JS_HEX_CONTENT)

}