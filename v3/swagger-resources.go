package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlabz/knife4go/constant"
	"github.com/jasonlabz/knife4go/utils"
)

const (
	// TODO 路径要改
	SWAGGER_RESOURCES_CONTENT     = `{"configUrl": "` + constant.RootPath + `/v3/api-docs/swagger-config","oauth2RedirectUrl": "` + constant.RootPath + `/swagger-ui/oauth2-redirect.html","url": "` + constant.RootPath + `/v3/api-docs","validatorUrl": ""}`
	SWAGGER_RESOURCES_CONFIG_PATH = constant.RootPath + "/v3/api-docs/swagger-config"
)

func AddSwaggerResourcesRouter(router gin.IRouter, path string) {
	if path == "" {
		path = SWAGGER_RESOURCES_CONFIG_PATH
	}
	utils.GetJson(router, SWAGGER_RESOURCES_CONFIG_PATH, SWAGGER_RESOURCES_CONTENT)
}
