package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlabz/knife4go/constant"
	"github.com/jasonlabz/knife4go/utils"
)

const (
	// TODO 路径要改
	API_DOCS_RELATIVE_PATH = constant.RootPath + "/v3/api-docs"
)

// @param content string swag int 命令生成的swagger.json文件里的内容
func AddApiDocRouter(router gin.IRouter, path, swaggerJson string) {
	if path == "" {
		path = API_DOCS_RELATIVE_PATH
	}
	utils.GetJson(router, path, swaggerJson)
}
