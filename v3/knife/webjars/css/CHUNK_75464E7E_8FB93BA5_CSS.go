package css

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlabz/knife4go/constant"
	"github.com/jasonlabz/knife4go/utils"
)

const (
	CHUNK_75464E7E_8FB93BA5_CSS_RELATIVE_PATH = constant.RootPath + "/webjars/css/chunk-75464e7e.8fb93ba5.css"
	// 文件内容的16进制表示
	CHUNK_75464E7E_8FB93BA5_CSS_HEX_CONTENT = `.api-title[data-v-b092cbdc]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}`
)

func AddRouterOfChunk75464e7e8fb93ba5Css(router gin.IRouter) {

	utils.GetCss(router, CHUNK_75464E7E_8FB93BA5_CSS_RELATIVE_PATH, CHUNK_75464E7E_8FB93BA5_CSS_HEX_CONTENT)

}
