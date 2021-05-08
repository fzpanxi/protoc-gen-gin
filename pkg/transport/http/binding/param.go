package binding

import (
	"github.com/gin-gonic/gin"
)

func RestParams(ctx *gin.Context) map[string]string {
	values := make(map[string]string)
	for _, v := range ctx.Params {
		values[v.Key] = v.Value
	}
	return values
}
func QueryParams(ctx *gin.Context) map[string]string {
	values := make(map[string]string)
	for k, v := range ctx.Request.URL.Query() {
		if len(v) > 0 {
			values[k] = v[0]
		}
	}
	return values
}
