package binding

import (
	"github.com/fzpanxi/protoc-gen-go-gin/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, out interface{}, err error) {
	if err != nil {
		ctx.JSON(errors.Convert(err))
	} else {
		ctx.JSON(http.StatusOK, out)
	}
}
