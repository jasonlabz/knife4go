# knife4go
simply Assembled knife4j + gin-swagger, it means an enhanced version of gin-swagger with nice UI.  

work well on go1.21+

# go get
`go get github.com/jasonlabz/knife4go`

## usage
Based on the usage of gin-swagger, add sth as follows:
- openapi 2.0
```go
import (
    knife4goFiles "github.com/jasonlabz/knife4go"
    knife4goGin "github.com/jasonlabz/knife4go/gin"
    _ "github.com/.../docs"
    ...
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	// swagger base router
	// knife4go: beautify swagger-ui
    router.GET("/swagger/*any", knife4goGin.WrapHandler(knife4goFiles.Handler))
    ...
}
```

- openapi 3.0
```go
import (
    knife4go "github.com/jasonlabz/knife4go"
    _ "github.com/.../docs"
    ...
)

func InitRouters() *gin.Engine {
	// swagger base router
	// knife4go: beautify swagger-ui
	_ = knife4go.InitSwaggerKnife(serverGroup)
    ...
}
```
## Disclaimer
Public welfare projects.  
The disclaimer asserts that the individual won't be held responsible for any.

## MIT LICENSE
[LICENSE](./LICENSE)

# Links
- gin-swagger: https://github.com/swaggo/gin-swagger
- knife4j: https://github.com/xiaoymin/knife4j
- swag: https://github.com/swaggo/swag/cmd/swag
